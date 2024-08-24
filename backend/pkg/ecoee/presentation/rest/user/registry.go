package user

import (
	"ecoee/pkg/ecoee/domain/model"
	"ecoee/pkg/ecoee/domain/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"log/slog"
	"net/http"
	"time"
)

type User struct {
	ID               string  `json:"user_id"`
	Name             string  `json:"user_name" binding:"required"`
	OrganizationID   string  `json:"organization_id"`
	OrganizationName string  `json:"organization_name"`
	TotalPoint       int     `json:"total_point"`
	PointHistory     []Point `json:"donation_history"`
}

type Point struct {
	ID        string    `json:"point_id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Registry struct {
	userRepository         model.UserRepository
	organizationRepository model.OrganizationRepository
	pointService           service.PointService
}

func NewRegistry(userRepository model.UserRepository,
	organizationRepository model.OrganizationRepository,
	pointService service.PointService,
) *Registry {
	return &Registry{
		userRepository:         userRepository,
		organizationRepository: organizationRepository,
		pointService:           pointService,
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
		if errors.Is(err, model.ErrOrganizationNotFound) {
			ctx.Status(http.StatusNotFound)
			return
		}

		ctx.Status(http.StatusInternalServerError)
		return
	}

	user := model.User{
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

	ctx.JSON(http.StatusCreated, fromDomainUser(createdUser, org, nil))
}

func (r *Registry) getUserProfile(ctx *gin.Context) {
	orgId := ctx.Param("orgId")
	userId := ctx.Param("userId")

	user, err := r.userRepository.GetByID(ctx, orgId, userId)
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			ctx.Status(http.StatusNotFound)
			return
		}
		slog.Error(fmt.Sprintf("failed to get user: %v", err))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	org, err := r.organizationRepository.GetByID(ctx, orgId)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get organization: %v", err))
		if errors.Is(err, model.ErrOrganizationNotFound) {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.Status(http.StatusInternalServerError)
		return
	}

	userPoints, err := r.pointService.ListUserPointDesc(ctx, userId)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get user point history: %v", err))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	if len(userPoints) > 3 {
		userPoints = userPoints[:3]
	}
	ctx.JSON(http.StatusOK, fromDomainUser(user, org, userPoints))
}

func fromDomainUser(u model.User, org model.Organization, ups []model.UserPoint) User {
	return User{
		ID:               u.ID,
		Name:             u.Name,
		OrganizationID:   org.ID,
		OrganizationName: org.Name,
		TotalPoint:       u.TotalUserPoint,
		PointHistory:     fromDomainUserPoints(ups),
	}
}

func fromDomainUserPoints(ups []model.UserPoint) []Point {
	points := make([]Point, 0, len(ups))
	for _, up := range ups {
		points = append(points, Point{
			ID:        up.ID,
			UserID:    up.UserID,
			Title:     up.Title,
			Amount:    up.Amount,
			CreatedAt: up.CreatedAt,
		})
	}
	return points
}
