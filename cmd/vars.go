package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/internal"
	"github.com/yanglinz/backpack/tools"
)

var varsListCmd = &cobra.Command{
	Use:   "list",
	Short: "📖 List variables",
	Long:  "📖 List variables",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		backpack := internal.ParseContext(cmd)
		tools.ListSecrets(backpack)
	},
}

var varsNewCmd = &cobra.Command{
	Use:   "new",
	Short: "💾 Put a new variable by name",
	Long:  "💾 Put a new variable by name",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		backpack := internal.ParseContext(cmd)
		tools.CreateSecret(backpack, tools.CreateSecretRequest{
			Name:  args[0],
			Value: args[1],
		})
	},
}

var varsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "💾 Update a new variable by name",
	Long:  "💾 Update a new variable by name",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		backpack := internal.ParseContext(cmd)
		tools.UpdateSecret(backpack, tools.UpdateSecretRequest{
			Name:  args[0],
			Value: args[1],
		})
	},
}

var varsCmd = &cobra.Command{
	Use:   "vars",
	Short: "🔑 Configure environmental variables and secrets",
	Long:  "🔑 Configure environmental variables and secrets",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	varsCmd.AddCommand(varsListCmd)

	varsNewCmd.Flags().String("env", internal.Development, "environment to put vars to")
	varsCmd.AddCommand(varsNewCmd)

	varsUpdateCmd.Flags().String("env", internal.Development, "environment to put vars to")
	varsCmd.AddCommand(varsUpdateCmd)

	rootCmd.AddCommand(varsCmd)
}
