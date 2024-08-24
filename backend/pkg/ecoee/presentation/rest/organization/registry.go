package organization

import (
	"ecoee/pkg/ecoee/domain"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

type CreateOrganizationRequest struct {
	Name                 string `json:"name" binding:"required"`
	MinimumDonationPoint int    `json:"minimum_donation_point"`
}

type Organization struct {
	ID                   string             `json:"organization_id"`
	Name                 string             `json:"organization_name"`
	TotalDonationPoint   int                `json:"total_donation_point"`
	MinimumDonationPoint int                `json:"minimum_donation_point"`
	DonationRanking      []OrganizationUser `json:"donation_ranking"`
}

type OrganizationUser struct {
	UserID           string `json:"user_id"`
	UserName         string `json:"user_name"`
	AccumulatedPoint int    `json:"accumulated_point"`
}

type Registry struct {
	organizationRepository domain.OrganizationRepository
}

func NewRegistry(organizationRepository domain.OrganizationRepository) *Registry {
	return &Registry{
		organizationRepository: organizationRepository,
	}
}

func (r *Registry) Register(e *gin.Engine) {
	e.POST("/api/orgs", r.createOrganization)
	e.GET("/api/orgs/:orgId/donation", r.getOrganization)
}

func (r *Registry) createOrganization(ctx *gin.Context) {
	req := &CreateOrganizationRequest{}
	if err := ctx.ShouldBindBodyWithJSON(req); err != nil {
		slog.Error(fmt.Sprintf("failed to bind request: %v", err))
		ctx.Status(http.StatusBadRequest)
		return
	}

	newOrg := domain.Organization{
		ID:                   uuid.NewString(),
		Name:                 req.Name,
		TotalDonationPoint:   0,
		MinimumDonationPoint: req.MinimumDonationPoint,
	}
	org, err := r.organizationRepository.Save(ctx, newOrg)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to save organization: %v", err))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	dto := fromDomainOrganization(org)
	ctx.JSON(http.StatusCreated, dto)
}

func (r *Registry) getOrganization(ctx *gin.Context) {
	orgID := ctx.Param("orgId")
	org, err := r.organizationRepository.GetByID(ctx, orgID)
	if err != nil {
		if errors.Is(err, domain.ErrOrganizationNotFound) {
			ctx.Status(http.StatusNotFound)
			return
		}

		slog.Error(fmt.Sprintf("failed to get organization: %v", err))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	dto := fromDomainOrganization(org)
	ctx.JSON(http.StatusOK, dto)
}

func fromDomainOrganization(org domain.Organization) Organization {
	return Organization{
		ID:                   org.ID,
		Name:                 org.Name,
		TotalDonationPoint:   org.TotalDonationPoint,
		MinimumDonationPoint: org.MinimumDonationPoint,
		//DonationRanking:      fromDomainDonationRanking(org.DonationRanking),
	}
}

//func fromDomainDonationRanking(ranking domain.DonationRanking) []OrganizationUser {
//	orgUsers := make([]OrganizationUser, 0, len(ranking.OrganizationUsers))
//	for _, user := range ranking.OrganizationUsers {
//		orgUsers = append(orgUsers, OrganizationUser{
//			UserID:           user.User.ID,
//			UserName:         user.User.Name,
//			AccumulatedPoint: user.AccumulatedPoint,
//		})
//	}
//
//	return orgUsers
//}
