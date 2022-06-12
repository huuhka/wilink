/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// startStageCmd represents the startStage command
var startStageCmd = &cobra.Command{
	Use:   "startStage",
	Short: "Start a stage in a Classic Release",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("startStage called")
	},
}

func init() {
	releaseCmd.AddCommand(startStageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startStageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startStageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}