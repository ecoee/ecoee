package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetObjectIDFromHex(id string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, primitive.ErrInvalidHex
	}

	return objectID, nil
}
