package db

import (
	"context"
	"ecoee/pkg/ecoee/domain"
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

func (r *OrganizationRepository) GetByID(ctx context.Context, orgId string) (domain.Organization, error) {
	collection := r.db.Collection(_organizationCollection)
	filter := bson.M{"organization_id": orgId}
	org := &dto.Organization{}
	if err := collection.FindOne(ctx, filter).Decode(org); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Organization{}, domain.ErrOrganizationNotFound
		}
		return domain.Organization{}, errors.Wrapf(err, "failed to get organization by orgId %s", orgId)
	}

	return toDomainOrganization(org), nil
}

func (r *OrganizationRepository) Save(ctx context.Context, organization domain.Organization) (domain.Organization, error) {
	collection := r.db.Collection(_organizationCollection)
	orgDTO := dto.Organization{
		ID:                   organization.ID,
		Name:                 organization.Name,
		TotalPoint:           organization.TotalDonationPoint,
		MinimumDonationPoint: organization.MinimumDonationPoint,
	}
	_, err := collection.InsertOne(ctx, orgDTO)
	if err != nil {
		return domain.Organization{}, errors.Wrapf(err, "failed to save organization %s", organization.ID)
	}

	return organization, nil
}

func (r *OrganizationRepository) Update(ctx context.Context, organization domain.Organization) (domain.Organization, error) {
	collection := r.db.Collection(_organizationCollection)
	orgDTO := dto.Organization{
		ID:                   organization.ID,
		Name:                 organization.Name,
		TotalPoint:           organization.TotalDonationPoint,
		MinimumDonationPoint: organization.MinimumDonationPoint,
	}
	filter := bson.M{"id": organization.ID}
	if err := collection.FindOneAndUpdate(ctx, filter, bson.M{
		"$set": orgDTO,
	}).Err(); err != nil {
		return domain.Organization{}, errors.Wrapf(err, "failed to update organization %s", organization.ID)
	}

	return organization, nil
}

//func toDonationRankingDTO(ranking domain.DonationRanking) dto.DonationRanking {
//	return dto.DonationRanking{
//		OrganizationUsers: toOrganizationUserDTOs(ranking.OrganizationUsers),
//	}
//}

func toOrganizationUserDTOs(users []domain.OrganizationUser) []dto.OrganizationUser {
	orgUsers := make([]dto.OrganizationUser, 0, len(users))
	for _, user := range users {
		orgUsers = append(orgUsers, dto.OrganizationUser{
			User:             toUserDTO(user.User),
			AccumulatedPoint: user.AccumulatedPoint,
		})
	}

	return orgUsers
}

func toUserDTO(user domain.User) dto.User {
	return dto.User{
		ID:   user.ID,
		Name: user.Name,
	}
}

func toDomainOrganization(org *dto.Organization) domain.Organization {
	return domain.Organization{
		ID:                   org.ID,
		Name:                 org.Name,
		TotalDonationPoint:   org.TotalPoint,
		MinimumDonationPoint: org.MinimumDonationPoint,
		//DonationRanking:      toDomainDonationRanking(org.DonationRanking),
	}
}

//func toDomainDonationRanking(ranking dto.DonationRanking) domain.DonationRanking {
//	rankers := make([]domain.OrganizationUser, 0, len(ranking.OrganizationUsers))
//	for _, ranker := range ranking.OrganizationUsers {
//		rankers = append(rankers, domain.OrganizationUser{
//			User:             toDomainUser(&ranker.User),
//			AccumulatedPoint: ranker.AccumulatedPoint,
//		})
//	}
//
//	return domain.DonationRanking{OrganizationUsers: rankers}
//}
