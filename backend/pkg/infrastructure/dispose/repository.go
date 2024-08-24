package dispose

import (
	"context"
	"ecoee/pkg/domain/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const _collectionName = "ecoee"

type Repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(ctx context.Context, dispose model.Dispose) (model.Dispose, error) {
	collection := r.db.Collection(_collectionName)
	filter := bson.D{{Key: "name", Value: dispose.Name}}
	d := DTO{}
	if err := collection.FindOne(ctx, filter).Decode(&d); err != nil {
		if err != mongo.ErrNoDocuments {
			return model.Dispose{}, err
		}

		dto := DTO{ID: primitive.NewObjectID(), Name: dispose.Name, Count: dispose.Count}
		res, err := collection.InsertOne(ctx, dto)
		if err != nil {
			return model.Dispose{}, err
		}

		id := res.InsertedID.(primitive.ObjectID).Hex()
		return model.Dispose{ID: id, Name: dispose.Name, Count: dispose.Count}, nil
	}

	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"count": dispose.Count}})
	if err != nil {
		return model.Dispose{}, err
	}

	return model.Dispose{ID: d.ID.Hex(), Name: d.Name, Count: dispose.Count}, nil
}
