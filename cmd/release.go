/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
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
	releaseCmd.PersistentFlags().IntP("definitionId", "d", 0, "Id for the Classic Release Definition")
	releaseCmd.PersistentFlags().StringP("project", "p", "", "Name of the project")
	releaseCmd.MarkPersistentFlagRequired("definitionId")
	releaseCmd.MarkPersistentFlagRequired("project")
}