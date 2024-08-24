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
	_organizationCollection = "organization"
)

type OrganizationRepository struct {
	db *mongo.Database
}

func NewOrganizationRepository(db *mongo.Database) *OrganizationRepository {
	return &OrganizationRepository{db: db}
}

func (r *OrganizationRepository) GetByID(ctx context.Context, orgId string) (model.Organization, error) {
	collection := r.db.Collection(_organizationCollection)
	filter := bson.M{"organization_id": orgId}
	org := &dto.Organization{}
	if err := collection.FindOne(ctx, filter).Decode(org); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Organization{}, model.ErrOrganizationNotFound
		}
		return model.Organization{}, errors.Wrapf(err, "failed to get organization by orgId %s", orgId)
	}

	return toDomainOrganization(org), nil
}

func (r *OrganizationRepository) Create(ctx context.Context, organization model.Organization) (model.Organization, error) {
	collection := r.db.Collection(_organizationCollection)
	orgDTO := dto.Organization{
		ID:                   organization.ID,
		Name:                 organization.Name,
		TotalPoint:           organization.TotalDonationPoint,
		MinimumDonationPoint: organization.MinimumDonationPoint,
	}
	_, err := collection.InsertOne(ctx, orgDTO)
	if err != nil {
		return model.Organization{}, errors.Wrapf(err, "failed to save organization %s", organization.ID)
	}

	return organization, nil
}

func (r *OrganizationRepository) Update(ctx context.Context, organization model.Organization) (model.Organization, error) {
	collection := r.db.Collection(_organizationCollection)
	orgDTO := dto.Organization{
		ID:                   organization.ID,
		Name:                 organization.Name,
		TotalPoint:           organization.TotalDonationPoint,
		MinimumDonationPoint: organization.MinimumDonationPoint,
	}
	filter := bson.M{"organization_id": organization.ID}
	if err := collection.FindOneAndUpdate(ctx, filter, bson.M{
		"$set": orgDTO,
	}).Err(); err != nil {
		return model.Organization{}, errors.Wrapf(err, "failed to update organization %s", organization.ID)
	}

	return organization, nil
}

func toDomainOrganization(org *dto.Organization) model.Organization {
	return model.Organization{
		ID:                   org.ID,
		Name:                 org.Name,
		TotalDonationPoint:   org.TotalPoint,
		MinimumDonationPoint: org.MinimumDonationPoint,
	}
}
