package domain

import (
	"context"
	"time"
)

type UserPoint struct {
	Point
	UserID string
}

type OrgPoint struct {
	Point
	UserID string
	OrgID  string
}

type Point struct {
	ID        string
	Amount    int
	CreatedAt time.Time
}

type OrganizationUser struct {
	User             User
	AccumulatedPoint int
}

type PointRepository interface {
	ListUserPoints(ctx context.Context, orgID, userID string) ([]UserPoint, error)
	ListOrgDonationRankers(ctx context.Context, orgID string) ([]OrganizationUser, error)
	SaveUserPoint(ctx context.Context, point UserPoint) error
	SaveOrgPoint(ctx context.Context, point OrgPoint) error
}
