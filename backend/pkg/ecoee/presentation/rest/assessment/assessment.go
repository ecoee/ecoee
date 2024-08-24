package assessment

import (
	"ecoee/pkg/ecoee/domain"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecycleAssessmentRequest struct {
	ImageURL string `json:"image_url" binding:"required"`
}

type RecycleAssessmentResponse struct {
	IsSuccess bool   `json:"is_success"`
	Feedback  string `json:"feedback"`
}

type Registry struct {
	assessor domain.Assessor
}

func NewRegistry(assessor domain.Assessor) *Registry {
	return &Registry{assessor: assessor}
}

func (r *Registry) Register(router *gin.Engine) {
	router.POST("/assessment", r.assessRecycle)
}

func (r *Registry) assessRecycle(ctx *gin.Context) {
	req := &RecycleAssessmentRequest{}
	if err := ctx.ShouldBindBodyWithJSON(req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	rar := domain.RecycleAssessmentRequest{ImageURL: req.ImageURL}
	resp, err := r.assessor.Assess(ctx, rar)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to assess: %v", err))
		ctx.Status(http.StatusInternalServerError)
	}

	ctx.JSON(http.StatusOK, RecycleAssessmentResponse{
		IsSuccess: resp.IsSuccess,
		Feedback:  resp.Feedback,
	})
}
