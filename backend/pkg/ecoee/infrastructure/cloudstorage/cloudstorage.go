package cloudstorage

import (
	"cloud.google.com/go/storage"
	"context"
	"ecoee/pkg/config"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
	"log/slog"
)

const (
	_credentialPath = "./storage/storage-service-account.json"
)

type Repository struct {
	bucket     *storage.BucketHandle
	bucketName string
}

func NewRepository(ctx context.Context, config config.Config) (*Repository, error) {
	slog.Info(fmt.Sprintf("storage=%v", config.GCPConfig.Storage))
	credential, err := json.Marshal(config.GCPConfig.Storage)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to marshal storage config %v", errors.WithStack(err)))
		return nil, err
	}
	opt := option.WithCredentialsJSON(credential)
	client, err := storage.NewClient(ctx, opt)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create storage client %v", errors.WithStack(err)))
		return nil, err
	}

	bucket := client.Bucket(config.GCPConfig.CloudStorageConfig.BucketName)
	return &Repository{bucket: bucket, bucketName: config.GCPConfig.CloudStorageConfig.BucketName}, nil
}

func (r *Repository) Upload(ctx context.Context, objectName, contentType string, data []byte) (string, error) {
	wc := r.bucket.Object(objectName).NewWriter(ctx)
	wc.ContentType = contentType
	if _, err := wc.Write(data); err != nil {
		slog.Error(fmt.Sprintf("failed to write object %v", errors.WithStack(err)))
		return "", err
	}

	if err := wc.Close(); err != nil {
		slog.Error(fmt.Sprintf("failed to close writer %v", errors.WithStack(err)))
		return "", err
	}

	imageURL := "https://storage.googleapis.com/" + r.bucketName + "/" + objectName
	return imageURL, nil
}
