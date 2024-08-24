package dto

type User struct {
	ID               string `bson:"user_id"`
	Name             string `bson:"user_name"`
	OrganizationID   string `bson:"organization_id"`
	OrganizationName string `bson:"organization_name"`
	TotalUserPoint   int    `bson:"total_user_point"`
}
