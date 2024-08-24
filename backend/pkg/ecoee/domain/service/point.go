package service

import (
	"context"
	"ecoee/pkg/ecoee/domain/model"
	"fmt"
	"log/slog"
	"sort"
)

type OrganizationPointRankerQueryResponse struct {
	OrganizationID string
	Rankers        []Ranker
}

type Ranker struct {
	User             model.User
	AccumulatedPoint int
}

type PointService interface {
	ListOrganizationPointRankers(ctx context.Context, orgID string) (OrganizationPointRankerQueryResponse, error)
	ListUserPointDesc(ctx context.Context, userID string) ([]model.UserPoint, error)
}

type pointService struct {
	pointRepository model.PointRepository
	userRepository  model.UserRepository
}

func (s *pointService) ListOrganizationPointRankers(ctx context.Context, orgID string) (OrganizationPointRankerQueryResponse, error) {
	points, err := s.pointRepository.ListOrgPoints(ctx, orgID)
	if err != nil {
		return OrganizationPointRankerQueryResponse{}, err
	}

	// This is a map to store the accumulated points of each user.
	// The key is the user ID and the value is the accumulated point.
	accumulatedPointsByUserID := make(map[string]int)
	for _, point := range points {
		slog.Info(fmt.Sprintf("point id: %s, user id: %s, amount: %d", point.ID, point.UserID, point.Amount))
		accumulatedPointsByUserID[point.UserID] += point.Amount
	}

	// order by amount desc
	rankers := make([]Ranker, 0)
	for userID, accumulatedPoint := range accumulatedPointsByUserID {
		slog.Info(fmt.Sprintf("orgId: %s, userId: %s, accumulatedPoint: %d", orgID, userID, accumulatedPoint))
		user, err := s.userRepository.GetByID(ctx, orgID, userID)
		if err != nil {
			return OrganizationPointRankerQueryResponse{}, err
		}
		rankers = append(rankers, Ranker{
			User:             user,
			AccumulatedPoint: accumulatedPoint,
		})
	}
	sort.Slice(rankers, func(i, j int) bool {
		return rankers[i].AccumulatedPoint > rankers[j].AccumulatedPoint
	})

	return OrganizationPointRankerQueryResponse{
		OrganizationID: orgID,
		Rankers:        rankers,
	}, nil
}

func (s *pointService) ListUserPointDesc(ctx context.Context, userID string) ([]model.UserPoint, error) {
	points, err := s.pointRepository.ListUserPoints(ctx, userID)
	if err != nil {
		return nil, err
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].CreatedAt.After(points[j].CreatedAt)
	})

	slog.Info(fmt.Sprintf("user %s has %d points", userID, len(points)))
	return points, nil
}

func NewPointService(pointRepository model.PointRepository, userRepository model.UserRepository) *pointService {
	return &pointService{pointRepository: pointRepository, userRepository: userRepository}
}
