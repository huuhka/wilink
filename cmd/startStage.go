package cmd

import (
	"context"
	"fmt"
	"github.com/huuhka/wilink/adowrappers"
	"github.com/microsoft/azure-devops-go-api/azuredevops/release"
	"log"

	"github.com/spf13/cobra"
)

// startStageCmd represents the startStage command
var startStageCmd = &cobra.Command{
	Use:   "startStage",
	Short: "Start a stage in a Classic Release",
	Run:   startStage,
}

func init() {
	releaseCmd.AddCommand(startStageCmd)

	releaseCmd.Flags().IntP("releaseId", "r", 0, "Id for the Classic Release Instance")
	releaseCmd.MarkFlagRequired("releaseId")

	startStageCmd.Flags().String("stage", "", "Name of the stage to start")
	startStageCmd.MarkFlagRequired("stage")
}

func startStage(cmd *cobra.Command, args []string) {
	fmt.Println("startStage called")

	bearer, err := cmd.Flags().GetString("bearer")
	if err != nil {
		log.Fatal(err)
	}

	project, err := cmd.Flags().GetString("project")
	if err != nil {
		log.Fatalf("project parameter not set")
	}
	releaseId, err := cmd.Flags().GetInt("releaseId")
	if err != nil {
		log.Fatalf("releaseId parameter not set")
	}
	organizationUrl, err := cmd.Flags().GetString("organizationUrl")
	if err != nil {
		log.Fatalf("organizationUrl parameter not set")
	}
	stage, err := cmd.Flags().GetString("stage")
	if err != nil {
		log.Fatalf("stage parameter not set")
	}

	c, err := release.NewClient(context.Background(), adowrappers.NewOauthConnection(organizationUrl, bearer))
	if err != nil {
		log.Fatal(err)
	}

	//rel, err := adowrappers.GetRelease(c, project, releaseId)
	//if err != nil {
	//	log.Fatal(err)
	//}
	_, _, _, _ = stage, project, releaseId, c
	//fmt.Printf("%+v", rel)
}