package internal

import (
	"github.com/spf13/cobra"
)

// Context for the overarching project
type Context struct {
	Root string
	Name string
}

// ParseContext returns Context for a given project
func ParseContext(cmd *cobra.Command) Context {
	context := Context{
		Root: "",
		Name: "",
	}
	return context
}
