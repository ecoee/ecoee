package db

import (
	"context"
	"ecoee/pkg/ecoee/domain"
	"ecoee/pkg/ecoee/infrastructure/db/dto"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	_campaignCollection     = "campaign"
	_campaignUserCollection = "campaign_user"
)

type CampaignRepository struct {
	db *mongo.Database
}

func NewCampaignRepository(db *mongo.Database) *CampaignRepository {
	return &CampaignRepository{db: db}
}

func (r *CampaignRepository) Create(ctx context.Context, campaign domain.Campaign) (domain.Campaign, error) {
	collection := r.db.Collection(_campaignCollection)
	campaignDTO := &dto.Campaign{
		ID:             campaign.ID,
		Name:           campaign.Name,
		OrganizationID: campaign.OrganizationID,
		ImageURL:       campaign.ImageURL,
		TotalVoted:     0,
	}
	_, err := collection.InsertOne(ctx, campaignDTO)
	if err != nil {
		return domain.Campaign{}, errors.Wrapf(err, "failed to insert campaign: %v", campaign)
	}

	return campaign, nil
}

func (r *CampaignRepository) List(ctx context.Context, orgID string) ([]domain.Campaign, error) {
	collection := r.db.Collection(_campaignCollection)
	filter := dto.Campaign{OrganizationID: orgID}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list campaigns for organization: %s", orgID)
	}
	defer cursor.Close(ctx)

	var campaigns []domain.Campaign
	for cursor.Next(ctx) {
		campaign := &dto.Campaign{}
		if err := cursor.Decode(campaign); err != nil {
			return nil, errors.Wrap(err, "failed to decode campaign")
		}
		campaigns = append(campaigns, domain.Campaign{
			ID:             campaign.ID,
			Name:           campaign.Name,
			OrganizationID: campaign.OrganizationID,
			ImageURL:       campaign.ImageURL,
			TotalVoted:     campaign.TotalVoted,
		})
	}

	return campaigns, nil
}

func (r *CampaignRepository) Vote(ctx context.Context, campaignID, userID string) error {
	collection := r.db.Collection(_campaignUserCollection)
	campaignVotedUser := &dto.CampaignVotedUser{
		CampaignID: campaignID,
		UserID:     userID,
	}
	_, err := collection.InsertOne(ctx, campaignVotedUser)
	if err != nil {
		return errors.Wrapf(err, "failed to vote campaign: %s for user: %s", campaignID, userID)
	}

	return nil
}

func (r *CampaignRepository) HasVoted(ctx context.Context, campaignID, userID string) (bool, error) {
	collection := r.db.Collection(_campaignUserCollection)
	filter := dto.CampaignVotedUser{
		CampaignID: campaignID,
		UserID:     userID,
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