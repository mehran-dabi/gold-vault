package ports

import (
	"context"
	"mime/multipart"
)

type (
	FileStorage interface {
		UploadFile(ctx context.Context, bucketName, objectName string, file multipart.File) error
	}
)
