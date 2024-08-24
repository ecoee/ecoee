package model

import (
	"context"
)

type ImageUploader interface {
	Upload(ctx context.Context, objectName, contentType string, data []byte) (string, error)
}
