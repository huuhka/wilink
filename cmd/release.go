package cmd

import (
	"github.com/spf13/cobra"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "This command is a hub for Classic Release Pipeline actions",
	Long:  `This command is a hub for Classic Release Pipeline actions`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("release called")
	//},
}

func init() {
	rootCmd.AddCommand(releaseCmd)
	releaseCmd.PersistentFlags().StringP("project", "p", "", "Name of the project")
	releaseCmd.PersistentFlags().StringP("organizationUrl", "o", "", "Url of the organization")
	releaseCmd.MarkPersistentFlagRequired("project")
	releaseCmd.MarkPersistentFlagRequired("organizationUrl")
}