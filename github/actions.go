package github

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/yanglinz/backpack/internal"
)

// CreateWorkflows generate github actions configs
func CreateWorkflows(backpack internal.Context) {
	workflowDir := filepath.Join(backpack.Root, ".github/workflows")
	os.MkdirAll(workflowDir, 0777)

	sourcePath := filepath.Join(backpack.Root, ".backpack/github/action-workflow.yml")
	targetPath := filepath.Join(backpack.Root, ".github/workflows/main.yml")
	content, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(targetPath, content, 0644)
	if err != nil {
		panic(err)
	}
}
