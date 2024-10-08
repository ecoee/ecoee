package model

import "context"

type Organization struct {
	ID                   string
	Name                 string
	TotalDonationPoint   int
	MinimumDonationPoint int
}

func (o *Organization) AddPoint(point OrgPoint) {
	o.TotalDonationPoint += point.Amount
}

type OrganizationRepository interface {
	GetByID(ctx context.Context, orgID string) (Organization, error)
	Create(ctx context.Context, organization Organization) (Organization, error)
	Update(ctx context.Context, organization Organization) (Organization, error)
}
