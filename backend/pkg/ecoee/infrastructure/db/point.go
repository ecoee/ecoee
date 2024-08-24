package db

import (
	"context"
	"ecoee/pkg/ecoee/domain"
	"ecoee/pkg/ecoee/infrastructure/db/dto"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

const (
	_userPointCollection = "user_point"
	_orgPointCollection  = "org_point"
)

type PointRepository struct {
	db *mongo.Database
}

func (r *PointRepository) ListUserPoints(ctx context.Context, orgID, userID string) ([]domain.UserPoint, error) {
	cursor, err := r.db.Collection(_userCollection).Find(ctx, bson.M{"org_id": orgID, "user_id": userID})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find user points")
	}

	var points []dto.UserPoint
	if err := cursor.All(ctx, &points); err != nil {
		return nil, errors.Wrapf(err, "failed to decode user points")
	}

	var userPoints []domain.UserPoint
	for _, point := range points {
		userPoints = append(userPoints, toDomainUserPoint(point))
	}

	return userPoints, nil
}

func (r *PointRepository) ListOrgDonationRankers(ctx context.Context, orgID string) ([]domain.OrganizationUser, error) {
	_, err := r.db.Collection(_organizationCollection).Find(ctx, bson.M{"organization_id": orgID})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find org points")
	}

	return nil, nil
}

func (r *PointRepository) SaveUserPoint(ctx context.Context, point domain.UserPoint) error {
	_, err := r.db.Collection(_userPointCollection).InsertOne(ctx, dto.UserPoint{
		Point: dto.Point{
			ID:        point.ID,
			Amount:    point.Amount,
			CreatedAt: point.CreatedAt,
		},
		UserID: point.UserID,
	})

	if err != nil {
		slog.Error(fmt.Sprintf("failed to save user point: %v", err))
		return errors.Wrapf(err, "failed to save user point")
	}

	return nil
}

func (r *PointRepository) SaveOrgPoint(ctx context.Context, point domain.OrgPoint) error {
	_, err := r.db.Collection(_orgPointCollection).InsertOne(ctx, dto.OrgPoint{
		Point: dto.Point{
			ID:        point.ID,
			Amount:    point.Amount,
			CreatedAt: point.CreatedAt,
		},
		UserID: point.UserID,
		OrgID:  point.OrgID,
	})

	if err != nil {
		slog.Error(fmt.Sprintf("failed to save org point: %v", err))
		return errors.Wrapf(err, "failed to save org point")
	}

	return nil
}

func NewPointRepository(db *mongo.Database) *PointRepository {
	return &PointRepository{db: db}
}

func toDomainUserPoint(point dto.UserPoint) domain.UserPoint {
	return domain.UserPoint{
		Point: domain.Point{
			ID:        point.ID,
			Amount:    point.Amount,
			CreatedAt: point.CreatedAt,
		},
		UserID: point.UserID,
	}
}

func toDomainOrgPoint(point dto.OrgPoint) domain.OrgPoint {
	return domain.OrgPoint{
		Point: domain.Point{
			ID:        point.ID,
			Amount:    point.Amount,
			CreatedAt: point.CreatedAt,
		},
		OrgID: point.OrgID,
	}
}
