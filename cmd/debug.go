package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/internal"
)

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "🔧 Output debug info",
	Long:  "🔧 Output debug info",
	Run: func(cmd *cobra.Command, args []string) {
		context := internal.ParseContext(cmd)
		fmt.Println(context)
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
}
