package internal

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// Context for the overarching project
type Context struct {
	Root string
	Name string
}

func parseRootPath(cmd *cobra.Command) string {
	// Get the root project path based on flag or cwd
	root := ""
	if cmd.Flag("root") != nil {
		root = cmd.Flag("root").Value.String()
	} else {
		cwd, err := os.Getwd()
		root = cwd
		if err != nil {
			return ""
		}
	}

	// Convert to absolute path
	root, err := filepath.Abs(root)
	if err != nil {
		return ""
	}

	return root
}

// ParseContext returns Context for a given project
func ParseContext(cmd *cobra.Command) Context {
	rootPath := parseRootPath(cmd)
	if rootPath == "" {
		panic("Failed to parse project root path")
	}

	context := Context{
		Root: rootPath,
		Name: "",
	}
	return context
}
