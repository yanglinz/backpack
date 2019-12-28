package tools

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/berglas/pkg/berglas"
	"github.com/yanglinz/backpack/internal"
)

func bucketExists(bucketName string) bool {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	bucket := client.Bucket(bucketName)
	bucketAttrs, err := bucket.Attrs(ctx)

	if bucketAttrs != nil {
		return true
	}
	return false
}

// BootstrapSecrets for berglas
func BootstrapSecrets(backpack internal.Context) {
	ctx := context.Background()
	bucketNames := []string{
		"backpack-berglas-" + backpack.Name,
		"backpack-dev-berglas-" + backpack.Name,
	}
	for _, bucket := range bucketNames {
		exists := bucketExists(bucket)
		if !exists {
			err := berglas.Bootstrap(ctx, &berglas.StorageBootstrapRequest{
				ProjectID: backpack.Google.ProjectID,
				Bucket:    bucket,
			})
			if err != nil {
				panic(err)
			}
		}
	}
}
