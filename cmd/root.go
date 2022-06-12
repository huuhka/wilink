/*
Copyright © 2022 Pasi Huuhka pasi@huuhka.net

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wilink",
	Short: "A tool to help linking Azure DevOps yaml pipeline runs to work items",
	Long: `A tool to help linking Azure DevOps yaml pipeline runs to a work item.

This tool has a pre-requirement of a Classic Release Pipeline with stages for each environment you need.
The Classic Release also needs to have the release integration turned on.

When pre-requirements are met, this tool can be used to create an instance of that release definition,
then modify it's artifacts to match the ones your yaml release is using, and trigger each "manually" triggered
stage when the specific environment has been deployed. 
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("bearer", "", "oAuth token for Azure DevOps. Either this or the -pat argument is needed. Use $(System.AccessToken) in pipelines.")
	rootCmd.PersistentFlags().String("pat", "", "personal access token. Either this or the -bearer argument needs to be given.")
}