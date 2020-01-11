package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/development"
	"github.com/yanglinz/backpack/internal"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "🐳 Run development server",
	Long:  "🐳 Run development server",
	Run: func(cmd *cobra.Command, args []string) {
		backpack := internal.ParseContext(cmd)
		development.RunTaskfile(backpack)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
