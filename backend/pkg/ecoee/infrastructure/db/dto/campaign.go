package dto

type Campaign struct {
	ID             string `bson:"id"`
	Name           string `bson:"name"`
	OrganizationID string `bson:"organization_id"`
	ImageURL       string `bson:"image_url"`
	TotalVoted     int    `bson:"total_voted"`
}

type CampaignVotedUser struct {
	UserID     string `bson:"user_id"`
	CampaignID string `bson:"campaign_id"`
}
