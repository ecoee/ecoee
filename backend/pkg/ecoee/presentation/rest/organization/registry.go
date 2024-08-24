package organization

import (
	"ecoee/pkg/ecoee/domain/model"
	"ecoee/pkg/ecoee/domain/service"
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
	organizationRepository model.OrganizationRepository
	pointService           service.PointService
}

func NewRegistry(organizationRepository model.OrganizationRepository, pointService service.PointService) *Registry {
	return &Registry{
		organizationRepository: organizationRepository,
		pointService:           pointService,
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

	newOrg := model.Organization{
		ID:                   uuid.NewString(),
		Name:                 req.Name,
		TotalDonationPoint:   0,
		MinimumDonationPoint: req.MinimumDonationPoint,
	}
	_, err := r.organizationRepository.Create(ctx, newOrg)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to save organization: %v", err))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, fromDomainOrganization(newOrg))
}

func (r *Registry) getOrganization(ctx *gin.Context) {
	orgID := ctx.Param("orgId")
	org, err := r.organizationRepository.GetByID(ctx, orgID)
	if err != nil {
		if errors.Is(err, model.ErrOrganizationNotFound) {
			ctx.Status(http.StatusNotFound)
			return
		}

		slog.Error(fmt.Sprintf("failed to get organization: %v", err))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	resp, err := r.pointService.ListOrganizationPointRankers(ctx, orgID)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get organization point rankers: %v", err))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	dto := fromDomainOrganizationWithRanking(org, resp)
	ctx.JSON(http.StatusOK, dto)
}

func fromDomainOrganization(org model.Organization) Organization {
	return Organization{
		ID:                   org.ID,
		Name:                 org.Name,
		TotalDonationPoint:   org.TotalDonationPoint,
		MinimumDonationPoint: org.MinimumDonationPoint,
	}
}

func fromDomainOrganizationWithRanking(org model.Organization, resp service.OrganizationPointRankerQueryResponse) Organization {
	return Organization{
		ID:                   org.ID,
		Name:                 org.Name,
		TotalDonationPoint:   org.TotalDonationPoint,
		MinimumDonationPoint: org.MinimumDonationPoint,
		DonationRanking:      fromDomainDonationRanking(resp),
	}
}

func fromDomainDonationRanking(resp service.OrganizationPointRankerQueryResponse) []OrganizationUser {
	orgUsers := make([]OrganizationUser, 0, len(resp.Rankers))
	for _, user := range resp.Rankers {
		orgUsers = append(orgUsers, OrganizationUser{
			UserID:           user.User.ID,
			UserName:         user.User.Name,
			AccumulatedPoint: user.AccumulatedPoint,
		})
	}

	return orgUsers
}
