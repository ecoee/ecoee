package dto

type Organization struct {
	ID                   string          `bson:"organization_id"`
	Name                 string          `bson:"organization_name"`
	TotalPoint           int             `bson:"total_point"`
	MinimumDonationPoint int             `bson:"minimum_donation_point"`
	DonationRanking      DonationRanking `bson:"donation_ranking"`
}

type OrganizationUser struct {
	User             User `bson:"user"`
	AccumulatedPoint int  `bson:"accumulated_point"`
}

type DonationRanking struct {
	OrganizationUsers []OrganizationUser `bson:"rankers"`
}
