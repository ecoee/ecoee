package campaign

import (
	"ecoee/pkg/ecoee/domain/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"log/slog"
	"net/http"
)

type CreateCampaignRequest struct {
	Title string `form:"title" binding:"required"`
	Body  string `form:"body" binding:"required"`
}

type Campaign struct {
	ID             string `json:"campaign_id"`
	OrganizationID string `json:"organization_id"`
	Title          string `json:"title" binding:"required"`
	Body           string `json:"body" binding:"required"`
	ImageURL       string `json:"image_url"`
	TotalVoted     int    `json:"total_voted"`
}

type CampaignUserVoted struct {
	HasVoted bool `json:"has_voted"`
}

type Registry struct {
	campaignRepository model.CampaignRepository
	userRepository     model.UserRepository
	imageUploader      model.ImageUploader
}

func NewRegistry(campaignRepository model.CampaignRepository,
	userRepository model.UserRepository,
	imageUploader model.ImageUploader,
) *Registry {
	return &Registry{campaignRepository: campaignRepository,
		userRepository: userRepository,
		imageUploader:  imageUploader,
	}
}

func (r *Registry) Registry(e *gin.Engine) {
	e.POST("/api/orgs/:orgId/campaigns", r.createCampaign)
	e.GET("/api/orgs/:orgId/campaigns", r.listCampaigns)
	e.POST("/api/orgs/:orgId/campaigns/:campaignId/users/:userId/vote", r.vote)
	e.GET("/api/orgs/:orgId/campaigns/:campaignId/users/:userId/vote", r.getVote)
}

func (r *Registry) createCampaign(ctx *gin.Context) {
	orgID := ctx.Param("orgId")
	campaign := &CreateCampaignRequest{}
	if err := ctx.ShouldBind(campaign); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get image: %v", err))
		ctx.Status(http.StatusBadRequest)
		return
	}

	mimeType := file.Header.Get("Content-Type")
	if mimeType != "image/jpeg" && mimeType != "image/png" {
		slog.Error("invalid image type")
		ctx.Status(http.StatusBadRequest)
		return
	}
	imageFormat := mimeType[6:]

	data := make([]byte, file.Size)
	src, err := file.Open()
	if err != nil {
		slog.Error(fmt.Sprintf("failed to open image: %v", err))
		ctx.Status(http.StatusInternalServerError)
		return
	}
	_, err = src.Read(data)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to read image: %v", err))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	imageName := fmt.Sprintf("%s.%s", uuid.NewString(), imageFormat)
	contentType := mimeType
	imageURL, err := r.imageUploader.Upload(ctx, imageName, contentType, data)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to upload image: %v", err))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	domainCampaign := model.Campaign{
		ID:             uuid.NewString(),
		OrganizationID: orgID,
		Title:          campaign.Title,
		Body:           campaign.Body,
		ImageURL:       imageURL,
		TotalVoted:     0,
	}
	_, err = r.campaignRepository.Create(ctx, domainCampaign)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, fromDomainCampaign(domainCampaign))
}

func (r *Registry) listCampaigns(ctx *gin.Context) {
	orgID := ctx.Param("orgId")
	campaigns, err := r.campaignRepository.List(ctx, orgID)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to list campaigns: %v", err))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, fromDomainCampaigns(campaigns))
}

func (r *Registry) vote(ctx *gin.Context) {
	orgID := ctx.Param("orgId")
	campaignID := ctx.Param("campaignId")
	userID := ctx.Param("userId")

	_, err := r.userRepository.GetByID(ctx, orgID, userID)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get user: %v", err))
		if errors.Is(err, model.ErrUserNotFound) {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.Status(http.StatusInternalServerError)
		return
	}

	hasVoted, err := r.campaignRepository.HasVoted(ctx, campaignID, userID)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to check if user has voted: %v", err))
		ctx.Status(http.StatusInternalServerError)
		return
	}
	if hasVoted {
		ctx.Status(http.StatusConflict)
		return
	}

	err = r.campaignRepository.Vote(ctx, campaignID, userID)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to vote campaign: %v", err))
		if errors.Is(err, model.ErrCampaignNotFound) {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (r *Registry) getVote(ctx *gin.Context) {
	orgID := ctx.Param("orgId")
	campaignID := ctx.Param("campaignId")
	userID := ctx.Param("userId")

	_, err := r.userRepository.GetByID(ctx, orgID, userID)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get user: %v", err))
		if errors.Is(err, model.ErrUserNotFound) {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.Status(http.StatusInternalServerError)
		return
	}

	hasVoted, err := r.campaignRepository.HasVoted(ctx, campaignID, userID)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to check if user has voted: %v", err))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, CampaignUserVoted{HasVoted: hasVoted})
}

func fromDomainCampaigns(campaigns []model.Campaign) []Campaign {
	result := []Campaign{}
	for _, campaign := range campaigns {
		result = append(result, fromDomainCampaign(campaign))
	}
	return result
}

func fromDomainCampaign(campaign model.Campaign) Campaign {
	return Campaign{
		ID:             campaign.ID,
		OrganizationID: campaign.OrganizationID,
		Title:          campaign.Title,
		Body:           campaign.Body,
		ImageURL:       campaign.ImageURL,
		TotalVoted:     campaign.TotalVoted,
	}
}
