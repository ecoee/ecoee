package dispose

import "go.mongodb.org/mongo-driver/bson/primitive"

type DTO struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"name"`
	Count int                `bson:"count"`
}
