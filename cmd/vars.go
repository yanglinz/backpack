package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yanglinz/backpack/google"
	"github.com/yanglinz/backpack/internal"
	"github.com/yanglinz/backpack/symbols"
)

var varsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Sync variables from cloud to local file",
	Long:  "Sync variables from cloud to local file",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		backpack := internal.ParseContext(cmd)
		secret := google.GetSecrets(backpack, env)

		var envJSON map[string]string
		json.Unmarshal([]byte(secret), &envJSON)
		formattedJSON, err := json.MarshalIndent(envJSON, "", "  ")
		if err != nil {
			panic(err)
		}

		fmt.Println(string(formattedJSON))
	},
}

var varsPutCmd = &cobra.Command{
	Use:   "put",
	Short: "Sync variables from local file to cloud",
	Long:  "Sync variables from local file to cloud",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		backpack := internal.ParseContext(cmd)

		envFile := filepath.Join(backpack.Root, "etc/development.json")
		if env == symbols.EnvProduction {
			envFile = filepath.Join(backpack.Root, "etc/production.json")
		}
		envData, err := ioutil.ReadFile(envFile)
		if err != nil {
			panic(err)
		}

		var envJSON map[string]string
		json.Unmarshal(envData, &envJSON)
		formattedJSON, err := json.Marshal(envJSON)
		if err != nil {
			panic(err)
		}

		google.UpdateSecrets(backpack, google.UpdateSecretRequest{
			Env:   env,
			Value: string(formattedJSON),
		})
	},
}

var varsCmd = &cobra.Command{
	Use:   "vars",
	Short: "Configure environmental variables and secrets",
	Long:  "Configure environmental variables and secrets",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	varsGetCmd.Flags().String("env", symbols.EnvDevelopment, "environment")
	varsCmd.AddCommand(varsGetCmd)
	varsPutCmd.Flags().String("env", symbols.EnvDevelopment, "environment")
	varsCmd.AddCommand(varsPutCmd)

	rootCmd.AddCommand(varsCmd)
}
