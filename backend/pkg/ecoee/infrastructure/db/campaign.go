package db

import (
	"context"
	"ecoee/pkg/ecoee/domain/model"
	"ecoee/pkg/ecoee/infrastructure/db/dto"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	_campaignCollection          = "campaign"
	_campaignVotedUserCollection = "campaign_voted_user"
)

type CampaignRepository struct {
	db *mongo.Database
}

func NewCampaignRepository(db *mongo.Database) *CampaignRepository {
	return &CampaignRepository{db: db}
}

func (r *CampaignRepository) Create(ctx context.Context, campaign model.Campaign) (model.Campaign, error) {
	collection := r.db.Collection(_campaignCollection)
	campaignDTO := &dto.Campaign{
		ID:             campaign.ID,
		Title:          campaign.Title,
		Body:           campaign.Body,
		OrganizationID: campaign.OrganizationID,
		ImageURL:       campaign.ImageURL,
		TotalVoted:     0,
	}
	_, err := collection.InsertOne(ctx, campaignDTO)
	if err != nil {
		return model.Campaign{}, errors.Wrapf(err, "failed to insert campaign: %v", campaign)
	}

	return campaign, nil
}

func (r *CampaignRepository) List(ctx context.Context, orgID string) ([]model.Campaign, error) {
	collection := r.db.Collection(_campaignCollection)
	cursor, err := collection.Find(ctx, bson.M{
		"organization_id": orgID,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list campaigns for organization: %s", orgID)
	}
	defer cursor.Close(ctx)

	var campaigns []model.Campaign
	for cursor.Next(ctx) {
		campaign := &dto.Campaign{}
		if err := cursor.Decode(campaign); err != nil {
			return nil, errors.Wrap(err, "failed to decode campaign")
		}
		campaigns = append(campaigns, model.Campaign{
			ID:             campaign.ID,
			Title:          campaign.Title,
			Body:           campaign.Body,
			OrganizationID: campaign.OrganizationID,
			ImageURL:       campaign.ImageURL,
			TotalVoted:     campaign.TotalVoted,
		})
	}

	return campaigns, nil
}

func (r *CampaignRepository) Vote(ctx context.Context, campaignID, userID string) error {
	collection := r.db.Collection(_campaignVotedUserCollection)
	campaignVotedUser := &dto.CampaignVotedUser{
		UserID: userID,
	}
	_, err := collection.InsertOne(ctx, campaignVotedUser)
	if err != nil {
		return errors.Wrapf(err, "failed to vote campaign: %s for user: %s", campaignID, userID)
	}

	return nil
}

func (r *CampaignRepository) HasVoted(ctx context.Context, campaignID, userID string) (bool, error) {
	collection := r.db.Collection(_campaignVotedUserCollection)
	filter := dto.CampaignVotedUser{
		UserID: userID,
	}
	res := &dto.CampaignVotedUser{}
	if err := collection.FindOne(ctx, filter).Decode(res); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}

		return false, errors.Wrapf(err, "failed to check if user: %s has voted campaign: %s", userID, campaignID)
	}

	return true, nil
}
