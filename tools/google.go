package tools

import (
	"context"
	"strings"

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
	bucketName := "backpack-berglas-" + backpack.Name
	exists := bucketExists(bucketName)
	if !exists {
		err := berglas.Bootstrap(ctx, &berglas.StorageBootstrapRequest{
			ProjectID: backpack.Google.ProjectID,
			Bucket:    bucketName,
		})
		if err != nil {
			panic(err)
		}

		CreateSecret(backpack, CreateSecretRequest{
			Name:  "BERGLAS_APP_JSON",
			Value: "{}",
		})
		CreateSecret(backpack, CreateSecretRequest{
			Name:  "BERGLAS_APP_DEV_JSON",
			Value: "{}",
		})
	}
}

// ListSecrets outputs a list of secrets
func ListSecrets(backpack internal.Context) {
	bucketName := "backpack-berglas-" + backpack.Name
	shell := internal.GetCommand("berglas list " + bucketName)
	err := shell.Run()
	if err != nil {
		panic(err)
	}
}

// CreateSecretRequest params
type CreateSecretRequest struct {
	Name  string
	Value string
}

// CreateSecret creates or updates a secret
func CreateSecret(backpack internal.Context, req CreateSecretRequest) {
	bucketName := "backpack-berglas-" + backpack.Name
	bucketPath := bucketName + "/" + req.Name
	encryptionKey := "projects/" + backpack.Google.ProjectID + "/locations/global/keyRings/berglas/cryptoKeys/berglas-key"
	parts := []string{
		"berglas create", bucketPath, req.Value,
		"--key", encryptionKey,
	}
	command := strings.Join(parts, " ")
	shell := internal.GetCommand(command)
	err := shell.Run()
	if err != nil {
		panic(err)
	}
}

// UpdateSecretRequest params
type UpdateSecretRequest struct {
	Name  string
	Value string
}

// UpdateSecret creates or updates a secret
func UpdateSecret(backpack internal.Context, req UpdateSecretRequest) {
	bucketName := "backpack-berglas-" + backpack.Name
	bucketPath := bucketName + "/" + req.Name
	encryptionKey := "projects/" + backpack.Google.ProjectID + "/locations/global/keyRings/berglas/cryptoKeys/berglas-key"
	parts := []string{
		"berglas update", bucketPath, req.Value,
		"--key", encryptionKey,
	}
	command := strings.Join(parts, " ")
	shell := internal.GetCommand(command)
	err := shell.Run()
	if err != nil {
		panic(err)
	}
}
