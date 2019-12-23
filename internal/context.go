package internal

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
)

type projectYaml struct {
	Name string
}

// Context for the overarching project
type Context struct {
	Root string
	Name string
}

func parseRootPath(cmd *cobra.Command) (string, error) {
	// Get the root project path based on flag or cwd
	root := ""
	if cmd.Flag("root") != nil {
		root = cmd.Flag("root").Value.String()
	} else {
		cwd, err := os.Getwd()
		root = cwd
		if err != nil {
			return "", err
		}
	}

	// Convert to absolute path
	root, err := filepath.Abs(root)
	if err != nil {
		return "", err
	}

	return root, nil
}

// ParseContext returns Context for a given project
func ParseContext(cmd *cobra.Command) Context {
	rootPath, err := parseRootPath(cmd)
	if err != nil {
		panic(err)
	}

	var project projectYaml
	source, err := ioutil.ReadFile(filepath.Join(rootPath, "pack.yml"))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &project)
	if err != nil {
		panic(err)
	}

	context := Context{
		Root: rootPath,
		Name: project.Name,
	}
	return context
}
