package domain

import (
	"context"
)

type ImageUploader interface {
	Upload(ctx context.Context, objectName string, data []byte) (string, error)
}
