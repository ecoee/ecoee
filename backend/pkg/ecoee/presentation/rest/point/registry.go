package point

import (
	"ecoee/pkg/ecoee/domain/model"
	"ecoee/pkg/ecoee/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"log/slog"
	"net/http"
)

type SavePointRequest struct {
	UserPointAmount         int `json:"user_point_amount"`
	OrganizationPointAmount int `json:"organization_point_amount"`
}

type DeductPointRequest struct {
	Amount int `json:"amount"`
}

type Registry struct {
	userRepository         model.UserRepository
	organizationRepository model.OrganizationRepository
	pointRepository        model.PointRepository
}

func NewRegistry(userRepository model.UserRepository,
	organizationRepository model.OrganizationRepository,
	pointRepository model.PointRepository,
) *Registry {
	return &Registry{
		userRepository:         userRepository,
		organizationRepository: organizationRepository,
		pointRepository:        pointRepository,
	}
}

func (r *Registry) Register(e *gin.Engine) {
	e.POST("/api/orgs/:orgId/users/:userId/points/add", r.addPoint)
	e.POST("/api/orgs/:orgId/users/:userId/points/deduct", r.deductUserPoint)
}

func (r *Registry) addPoint(ctx *gin.Context) {
	orgId := ctx.Param("orgId")
	userId := ctx.Param("userId")

	savePointReq := &SavePointRequest{}
	if err := ctx.ShouldBindBodyWithJSON(savePointReq); err != nil {
		slog.Error(fmt.Sprintf("failed to bind request: %v", err))
		ctx.Status(http.StatusBadRequest)
		return
	}
	user, err := r.userRepository.GetByID(ctx, orgId, userId)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get user: %v", err))
		if errors.Is(err, model.ErrUserNotFound) {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.Status(http.StatusInternalServerError)
		return
	}

	organization, err := r.organizationRepository.GetByID(ctx, orgId)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get organization: %v", err))
		if errors.Is(err, model.ErrOrganizationNotFound) {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.Status(http.StatusInternalServerError)
		return
	}

	// validated

	slog.Info("validated")

	if savePointReq.UserPointAmount > 0 {
		userPoint := model.UserPoint{
			Point: model.Point{
				ID:        uuid.NewString(),
				Amount:    savePointReq.UserPointAmount,
				CreatedAt: util.Now(),
			},
			UserID: user.ID,
		}

		if err := r.pointRepository.SaveUserPoint(ctx, userPoint); err != nil {
			slog.Error(fmt.Sprintf("failed to save user point: %v", err))
			ctx.Status(http.StatusInternalServerError)
			return
		}

		user.AddPoint(userPoint)
		_, err = r.userRepository.Save(ctx, user)
		if err != nil {
			slog.Error(fmt.Sprintf("failed to save user: %v", err))
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	if savePointReq.OrganizationPointAmount > 0 {
		orgPoint := model.OrgPoint{
			Point: model.Point{
				ID:        uuid.NewString(),
				Amount:    savePointReq.OrganizationPointAmount,
				CreatedAt: util.Now(),
			},
			UserID: userId,
			OrgID:  orgId,
		}
		organization.AddPoint(orgPoint)

		if err := r.pointRepository.SaveOrgPoint(ctx, orgPoint); err != nil {
			slog.Error(fmt.Sprintf("failed to save org point: %v", err))
			ctx.Status(http.StatusInternalServerError)
			return
		}

		_, err = r.organizationRepository.Update(ctx, organization)
		if err != nil {
			slog.Error(fmt.Sprintf("failed to save organization: %v", err))
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	ctx.Status(http.StatusOK)
}

func (r *Registry) deductUserPoint(ctx *gin.Context) {
	deductReq := &DeductPointRequest{}
	if err := ctx.ShouldBindBodyWithJSON(deductReq); err != nil {
		slog.Error(fmt.Sprintf("failed to bind request: %v", err))
		ctx.Status(http.StatusBadRequest)
		return
	}

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

	// 차감은 user 만 있음 org는 없음
	userPoint := model.UserPoint{
		Point: model.Point{
			ID:        uuid.NewString(),
			Amount:    deductReq.Amount,
			CreatedAt: util.Now(),
		},
		UserID: userId,
	}
	user.DeductPoint(userPoint)
	_, err = r.userRepository.Save(ctx, user)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to save user: %v", err))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	if err := r.pointRepository.SaveUserPoint(ctx, userPoint); err != nil {
		slog.Error(fmt.Sprintf("failed to save user point: %v", err))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
