package domain

import "context"

type User struct {
	ID               string
	Name             string
	OrganizationID   string
	OrganizationName string
	TotalUserPoint   int
}

func (u *User) AddPoint(point UserPoint) {
	u.TotalUserPoint += point.Amount
}

func (u *User) DeductPoint(point UserPoint) {
	u.TotalUserPoint -= point.Amount
}

type UserRepository interface {
	Create(ctx context.Context, user User) (User, error)
	GetByID(ctx context.Context, orgID, userID string) (User, error)
	Save(ctx context.Context, user User) (User, error)
}
