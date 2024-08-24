package gemini

import (
	"context"
	"ecoee/pkg/config"
	"ecoee/pkg/ecoee/domain"
	"encoding/json"
	"fmt"
	"log/slog"

	"cloud.google.com/go/vertexai/genai"
)

const (
	_modelName = "gemini-1.5-flash-001"
)

type Repository struct {
	gemini *genai.GenerativeModel
}

func NewRepository(ctx context.Context, config config.Config) (*Repository, error) {
	client, err := genai.NewClient(ctx, config.GCPConfig.ProjectID, config.GCPConfig.Location)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create GenAI client %v", err))
		return nil, err
	}

	gemini := client.GenerativeModel(_modelName)

	return &Repository{gemini: gemini}, nil
}

func (g *Repository) Assess(ctx context.Context, request domain.RecycleAssessmentRequest) (domain.RecycleAssessmentResponse, error) {
	img := genai.ImageData(request.Format, request.Data)

	// 1. 플라스틱 병을 인식 시킬 것
	// 2. 한국의 재활용 기준에 적합한지
	// 3. 불가능한 경우는 사용자가 취해야할 행동을 영어로 안내할 것
	// 4. 가능하다면 "OK" 문자열만 응답으로 반환할 것
	prompt := genai.Text("" +
		"I want to recycle this plastic bottle." +
		"1. Please identify this plastic bottle." +
		"2. Is it suitable for recycling in Korea?" +
		"3. If not, please tell the user what to do in English." +
		"4. If possible, return only the string 'OK' as a response.")
	resp, err := g.gemini.GenerateContent(ctx, img, prompt)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to generate content %v", err))
		return domain.RecycleAssessmentResponse{}, err
	}

	rb, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		slog.Error(fmt.Sprintf("failed to marshal response %v", err))
		return domain.RecycleAssessmentResponse{}, err
	}

	slog.Info(fmt.Sprintf("response: %s", rb))

	return domain.RecycleAssessmentResponse{
		IsSuccess: true,
		Feedback:  string(rb),
	}, nil
}
