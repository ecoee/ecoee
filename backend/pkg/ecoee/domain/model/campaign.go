package model

import "context"

type Campaign struct {
	ID             string
	OrganizationID string
	Title          string
	Body           string
	ImageURL       string
	TotalVoted     int
}

type CampaignUserVoted struct {
	User     User
	Campaign Campaign
}

type CampaignRepository interface {
	Create(ctx context.Context, campaign Campaign) (Campaign, error)
	List(ctx context.Context, orgID string) ([]Campaign, error)
	Vote(ctx context.Context, campaignID, userID string) error
	HasVoted(ctx context.Context, userID string) (bool, error)
}
