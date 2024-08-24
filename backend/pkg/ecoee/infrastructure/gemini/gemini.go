package gemini

import (
	"cloud.google.com/go/vertexai/genai"
	"context"
	"ecoee/pkg/config"
	"ecoee/pkg/ecoee/domain/model"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
	"log/slog"
)

type GeminiResponse struct {
	Result int `json:"result"`
}

const (
	_modelName = "gemini-1.5-pro-001"
	//_modelName      = "gemini-1.0-pro-vision-001"
	_credentialPath = "./vertexai-service-account.json"
)

type Repository struct {
	gemini *genai.GenerativeModel
}

func NewRepository(ctx context.Context, config config.Config) (*Repository, error) {
	opt := option.WithCredentialsFile(_credentialPath)
	client, err := genai.NewClient(ctx, config.GCPConfig.ProjectID, config.GCPConfig.Location, opt)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create GenAI client %v", errors.WithStack(err)))
		return nil, err
	}

	gemini := client.GenerativeModel(_modelName)

	return &Repository{gemini: gemini}, nil
}

func (g *Repository) Assess(ctx context.Context, request model.RecycleAssessmentRequest) (model.RecycleAssessmentResponse, error) {
	slog.Info(fmt.Sprintf("assess request: %v", request))
	img := genai.ImageData(request.Format, request.Data)
	prompt := genai.Text("" +
		" If this waste image has no plastic label and any content inside," +
		" then give json formatted result" +
		" based on the following criteria:\n" +
		" 1. Positive(Plastic label removed and any content inside)" +
		" 2: Negative(Plastic label not removed or some content inside)" +
		" 3: I don't know(If it's not plastic waste or unidentifiable)\n" +
		" NOTE: Please provide the result in the following format without any description.\n" +
		"   {\"result\": 1 }\"\n")
	resp, err := g.gemini.GenerateContent(ctx, img, prompt)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to generate content %v", errors.WithStack(err)))
		return model.RecycleAssessmentResponse{}, err
	}

	var answer genai.Text
	for _, r := range resp.Candidates {
		for _, p := range r.Content.Parts {
			answer = p.(genai.Text)
			break
		}
	}

	slog.Info(fmt.Sprintf("answer: %s", answer))

	response := &GeminiResponse{}
	if err := json.Unmarshal([]byte(answer), response); err != nil {
		slog.Error(fmt.Sprintf("failed to unmarshal response %v", errors.WithStack(err)))
		return model.RecycleAssessmentResponse{}, err
	}

	return model.RecycleAssessmentResponse{
		Result: response.Result,
	}, nil
}
