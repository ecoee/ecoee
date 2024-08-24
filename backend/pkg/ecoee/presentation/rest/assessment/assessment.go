package assessment

import (
	"ecoee/pkg/ecoee/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type RecycleAssessmentRequest struct {
	Format string `json:"format" binding:"required"`
}

type RecycleAssessmentResponse struct {
	IsSuccess bool   `json:"is_success"`
	Feedback  string `json:"feedback"`
}

type Registry struct {
	assessor domain.Assessor
}

func NewRegistry(
	assessor domain.Assessor,
) *Registry {
	return &Registry{
		assessor: assessor,
	}
}

func (r *Registry) Register(router *gin.Engine) {
	router.POST("/assess", r.assessRecycle)
}

func (r *Registry) assessRecycle(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}

	// convert file to byte
	src, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer src.Close()

	// Read the file
	data := make([]byte, file.Size)
	_, err = src.Read(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	rar := domain.RecycleAssessmentRequest{
		Format: "jpeg",
		Data:   data,
	}
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
