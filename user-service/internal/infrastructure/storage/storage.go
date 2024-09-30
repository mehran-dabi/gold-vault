package storage

import (
	"context"
	"fmt"
	"mime/multipart"

	"goldvault/user-service/internal/core/application/ports"

	"github.com/minio/minio-go/v7"
)

type FileStorage struct {
	client *minio.Client
}

func NewFileStorage(client *minio.Client) ports.FileStorage {
	return &FileStorage{client: client}
}

func (fs *FileStorage) UploadFile(ctx context.Context, bucketName, objectName string, file multipart.File) error {
	// Ensure bucket exists or create one
	exists, err := fs.client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if !exists {
		err = fs.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("could not create bucket: %w", err)
		}
	}

	// Upload the file using a reader
	_, err = fs.client.PutObject(ctx, bucketName, objectName, file, -1, minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("could not upload file: %w", err)
	}

	return nil
}
