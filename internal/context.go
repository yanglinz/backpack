package internal

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
)

// Project defines a domain-reachable app
type Project struct {
	Path string
	Name string
}

// Services represents external dependencies
type Services struct {
	Postgres bool
	Redis    bool
}

type contextYaml struct {
	Name     string
	Projects []Project
	Services Services
}

// Context for the overarching repository
type Context struct {
	Root     string
	Name     string
	Projects []Project
	Services Services
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

	var parsedContext contextYaml
	source, err := ioutil.ReadFile(filepath.Join(rootPath, "pack.yml"))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &parsedContext)
	if err != nil {
		panic(err)
	}

	context := Context{
		Root:     rootPath,
		Name:     parsedContext.Name,
		Projects: parsedContext.Projects,
		Services: parsedContext.Services,
	}
	return context
}
