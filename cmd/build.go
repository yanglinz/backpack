package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/internal"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "🛠  Build the docker images",
	Long:  "🛠  Build the docker images",
	Run: func(cmd *cobra.Command, args []string) {
		backpack := internal.ParseContext(cmd)

		command := "docker-compose build"
		shell := internal.GetCommand(command)
		shell.Dir = backpack.Root
		err := shell.Run()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
