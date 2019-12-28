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
		setupFiles, _ := cmd.Flags().GetBool("files")
		setupResources, _ := cmd.Flags().GetBool("resources")
		context := internal.ParseContext(cmd)

		if setupFiles {
			tools.CreateComposeConfig(context)
		}
		if setupResources {
			tools.BootstrapSecrets(context)
		}
	},
}

func init() {
	setupCmd.Flags().Bool("files", true, "setup project files")
	setupCmd.Flags().Bool("resources", false, "setup remote resources")
	rootCmd.AddCommand(setupCmd)
}
