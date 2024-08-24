package dto

type Campaign struct {
	ID             string `bson:"id"`
	OrganizationID string `bson:"organization_id"`
	Title          string `bson:"title"`
	Body           string `bson:"body"`
	ImageURL       string `bson:"image_url"`
	TotalVoted     int    `bson:"total_voted"`
}

type CampaignVotedUser struct {
	UserID     string `bson:"user_id"`
	CampaignID string `bson:"campaign_id"`
}
