package db

import (
	"context"
	"ecoee/pkg/ecoee/domain/model"
	"ecoee/pkg/ecoee/infrastructure/db/dto"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

const (
	_userCollection = "user"
)

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(ctx context.Context, user model.User) (model.User, error) {
	collection := r.db.Collection(_userCollection)
	userDTO := dto.User{
		ID:             user.ID,
		Name:           user.Name,
		OrganizationID: user.OrganizationID,
		TotalUserPoint: user.TotalUserPoint,
	}
	filter := bson.M{"organization_id": user.OrganizationID, "user_id": user.ID}
	if err := collection.FindOneAndUpdate(ctx, filter, bson.M{
		"$set": userDTO,
	}).Err(); err != nil {
		slog.Error(fmt.Sprintf("failed to save user: %v", err))
		return model.User{}, errors.Wrapf(err, "failed to save user %s", user.ID)
	}

	return user, nil
}

func (r *UserRepository) Create(ctx context.Context, user model.User) (model.User, error) {
	collection := r.db.Collection(_userCollection)
	userDTO := dto.User{
		ID:               user.ID,
		Name:             user.Name,
		OrganizationID:   user.OrganizationID,
		OrganizationName: user.OrganizationName,
		TotalUserPoint:   0,
	}
	_, err := collection.InsertOne(ctx, userDTO)
	if err != nil {
		return model.User{}, errors.Wrapf(err, "failed to create user %s", user.ID)
	}

	return user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, orgID, userID string) (model.User, error) {
	collection := r.db.Collection(_userCollection)
	filter := bson.M{"organization_id": orgID, "user_id": userID}
	user := &dto.User{}
	if err := collection.FindOne(ctx, filter).Decode(user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.User{}, model.ErrUserNotFound
		}
		return model.User{}, errors.WithStack(err)
	}

	return toDomainUser(user), nil
}

func toDomainUser(user *dto.User) model.User {
	return model.User{
		ID:               user.ID,
		Name:             user.Name,
		OrganizationID:   user.OrganizationID,
		OrganizationName: user.OrganizationName,
		TotalUserPoint:   user.TotalUserPoint,
	}
}
