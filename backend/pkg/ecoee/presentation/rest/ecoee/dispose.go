package ecoee

import (
	"ecoee/pkg/domain/model"
	"ecoee/pkg/ecoee/infrastructure/dispose"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Dispose struct {
	ID    string `json:"id"`
	Name  string `json:"name" binding:"required"`
	Count int    `json:"count" binding:"required"`
}

type Response struct {
	Message string `json:"message"`
}

type Registry struct {
	disposeRepository *dispose.Repository
}

func NewRegistry(disposeRepository *dispose.Repository) *Registry {
	return &Registry{
		disposeRepository: disposeRepository,
	}
}

func (r *Registry) Register(router *gin.Engine) {
	router.GET("/", r.base)
	router.POST("/dispose", r.dispose)
}

func (r *Registry) base(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{Message: "ecoee is working fine 👍"})
}

func (r *Registry) dispose(ctx *gin.Context) {
	req := &Dispose{}

	if err := ctx.ShouldBindBodyWithJSON(req); err != nil {
		slog.Error(fmt.Sprintf("failed to bind request: %v", err))
		ctx.Status(http.StatusBadRequest)
		return
	}

	newDispose := model.Dispose{
		ID:    req.ID,
		Name:  req.Name,
		Count: req.Count,
	}
	result, err := r.disposeRepository.Save(ctx, newDispose)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to save req: %v", err))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	res := Dispose{ID: result.ID, Name: result.Name, Count: result.Count}
	_ = ctx.ShouldBindBodyWithJSON(res)
}
