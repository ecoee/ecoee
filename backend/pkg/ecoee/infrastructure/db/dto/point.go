package dto

import "time"

type Point struct {
	ID        string    `bson:"id"`
	Amount    int       `bson:"amount"`
	CreatedAt time.Time `bson:"created_at"`
}

type UserPoint struct {
	Point  `bson:",inline"`
	UserID string `bson:"user_id"`
}

type OrgPoint struct {
	Point  `bson:",inline"`
	UserID string `bson:"user_id"`
	OrgID  string `bson:"organization_id"`
}
