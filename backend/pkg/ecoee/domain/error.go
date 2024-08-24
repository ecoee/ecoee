package domain

import "github.com/pkg/errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrCampaignNotFound     = errors.New("campaign not found")
	ErrOrganizationNotFound = errors.New("organization not found")
)
