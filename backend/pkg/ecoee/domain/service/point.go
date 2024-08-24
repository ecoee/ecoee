package service

import (
	"context"
	"ecoee/pkg/ecoee/domain/model"
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
		accumulatedPointsByUserID[point.UserID] += point.Amount
	}

	// order by amount desc
	rankers := make([]Ranker, 0)
	for userID, accumulatedPoint := range accumulatedPointsByUserID {
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

func newPointService(pointRepository model.PointRepository) *pointService {
	return &pointService{pointRepository: pointRepository}
}
