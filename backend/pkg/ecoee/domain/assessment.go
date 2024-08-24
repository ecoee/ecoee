package domain

import "context"

type RecycleAssessmentRequest struct {
	Format string
	Data   []byte
}

type RecycleAssessmentResponse struct {
	IsSuccess bool
	Feedback  string
}

type Assessor interface {
	Assess(ctx context.Context, query RecycleAssessmentRequest) (RecycleAssessmentResponse, error)
}
