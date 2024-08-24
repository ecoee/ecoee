package dto

type Organization struct {
	ID                   string `bson:"organization_id"`
	Name                 string `bson:"organization_name"`
	TotalPoint           int    `bson:"total_point"`
	MinimumDonationPoint int    `bson:"minimum_donation_point"`
}
