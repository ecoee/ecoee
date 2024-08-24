package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Registry struct{}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) Register(router *gin.Engine) {
	router.GET("/health", r.health)
}

func (r *Registry) health(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
