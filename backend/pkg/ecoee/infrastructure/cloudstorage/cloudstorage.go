package cloudstorage

import (
	"cloud.google.com/go/storage"
	"context"
	"ecoee/pkg/config"
	"fmt"
	"github.com/pkg/errors"
	"log/slog"
)

type Repository struct {
	bucket     *storage.BucketHandle
	bucketName string
}

func NewRepository(ctx context.Context, config config.Config) (*Repository, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create storage client %v", errors.WithStack(err)))
		return nil, err
	}

	bucket := client.Bucket(config.GCPConfig.CloudStorageConfig.BucketName)
	return &Repository{bucket: bucket, bucketName: config.GCPConfig.CloudStorageConfig.BucketName}, nil
}

func (r *Repository) Upload(ctx context.Context, objectName string, data []byte) (string, error) {
	wc := r.bucket.Object(objectName).NewWriter(ctx)
	wc.ContentType = "application/json"
	if _, err := wc.Write(data); err != nil {
		slog.Error(fmt.Sprintf("failed to write object %v", errors.WithStack(err)))
		return "", err
	}

	if err := wc.Close(); err != nil {
		slog.Error(fmt.Sprintf("failed to close writer %v", errors.WithStack(err)))
		return "", err
	}

	imageURL := r.bucketName + "/" + objectName
	return imageURL, nil
}
