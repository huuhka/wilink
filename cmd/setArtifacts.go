package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setArtifactsCmd represents the setArtifacts command
var setArtifactsCmd = &cobra.Command{
	Use:   "setArtifacts",
	Short: "Set the Artifact IDs for a Classic Release",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setArtifacts called")
	},
}

func init() {
	releaseCmd.AddCommand(setArtifactsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setArtifactsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setArtifactsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}