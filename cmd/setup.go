package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/internal"
	"github.com/yanglinz/backpack/tools"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "🏁 Setup project",
	Long:  "🏁 Setup project",
	Run: func(cmd *cobra.Command, args []string) {
		context := internal.ParseContext(cmd)
		tools.CreateComposeConfig(context)
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
