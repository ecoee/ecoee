package user

import (
	"ecoee/pkg/ecoee/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"log/slog"
	"net/http"
	"time"
)

type User struct {
	ID               string  `json:"id"`
	Name             string  `json:"name" binding:"required"`
	OrganizationID   string  `json:"organization_id"`
	OrganizationName string  `json:"organization_name"`
	PointHistory     []Point `json:"point_history"`
}

type Point struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type SavePointRequest struct {
	SelfAmount         int `json:"self_amount"`
	OrganizationAmount int `json:"organization_amount"`
}

type DeductPointRequest struct {
	Amount int `json:"amount"`
}

type Registry struct {
	userRepository         domain.UserRepository
	organizationRepository domain.OrganizationRepository
}

func NewRegistry(userRepository domain.UserRepository,
	organizationRepository domain.OrganizationRepository,
) *Registry {
	return &Registry{
		userRepository:         userRepository,
		organizationRepository: organizationRepository,
	}
}

func (r *Registry) Register(e *gin.Engine) {
	e.POST("/api/orgs/:orgId/users", r.createUser)
	e.GET("/api/orgs/:orgId/users/:userId/profile", r.getUserProfile)
}

func (r *Registry) createUser(ctx *gin.Context) {
	u := &User{}
	if err := ctx.ShouldBindBodyWithJSON(u); err != nil {
		slog.Error("failed to bind request: %v", err)
		ctx.Status(http.StatusBadRequest)
		return
	}
	orgId := ctx.Param("orgId")
	org, err := r.organizationRepository.GetByID(ctx, orgId)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get organization: %v", err))
		if errors.Is(err, domain.ErrOrganizationNotFound) {
			ctx.Status(http.StatusNotFound)
			return
		}

		ctx.Status(http.StatusInternalServerError)
		return
	}

	user := domain.User{
		ID:               uuid.New().String(),
		Name:             u.Name,
		OrganizationID:   org.ID,
		OrganizationName: org.Name,
	}
	createdUser, err := r.userRepository.Create(ctx, user)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create user: %v", err))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, fromDomainUser(createdUser))
}

func (r *Registry) getUserProfile(ctx *gin.Context) {
	orgId := ctx.Param("orgId")
	userId := ctx.Param("userId")

	user, err := r.userRepository.GetByID(ctx, orgId, userId)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			ctx.Status(http.StatusNotFound)
			return
		}
		slog.Error(fmt.Sprintf("failed to get user: %v", err))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, fromDomainUser(user))
}

func fromDomainUser(u domain.User) User {
	return User{
		ID:               u.ID,
		Name:             u.Name,
		OrganizationID:   u.OrganizationID,
		OrganizationName: u.OrganizationName,
		PointHistory:     nil,
	}
}

//func fromDomainPointHistory(h []domain.Point) []Point {
//	history := []Point{}
//	for _, p := range h {
//		history = append(history, Point{
//			ID:        p.ID,
//			UserID:    p.UserID,
//			Amount:    p.Amount,
//			CreatedAt: p.CreatedAt,
//		})
//	}
//
//	return history
//}
