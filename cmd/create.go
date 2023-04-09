package cmd

import (
	"context"
	"fmt"
	"github.com/huuhka/wilink/adowrappers"
	"github.com/microsoft/azure-devops-go-api/azuredevops/release"
	"github.com/spf13/cobra"
	"log"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new instance of a Classic Release Definition",
	Run:   createFunc,
}

func init() {
	releaseCmd.AddCommand(createCmd)
	createCmd.Flags().IntP("definitionId", "d", 0, "Id for the Classic Release Definition")
	createCmd.MarkFlagRequired("definitionId")
}

func createFunc(cmd *cobra.Command, args []string) {
	fmt.Println("create called")

	bearer, err := cmd.Flags().GetString("bearer")
	if err != nil {
		log.Fatal(err)
	}

	project, err := cmd.Flags().GetString("project")
	if err != nil {
		log.Fatalf("project parameter not set")
	}
	definitionId, err := cmd.Flags().GetInt("definitionId")
	if err != nil {
		log.Fatalf("definitionId parameter not set")
	}
	organizationUrl, err := cmd.Flags().GetString("organizationUrl")
	if err != nil {
		log.Fatalf("organizationUrl parameter not set")
	}

	c, err := release.NewClient(context.Background(), adowrappers.NewOauthConnection(organizationUrl, bearer))
	if err != nil {
		log.Fatal(err)
	}

	relDef, err := adowrappers.GetReleaseDefinition(c, project, definitionId)
	if err != nil {
		log.Fatal(err)
	}

	rel, err := adowrappers.CreateRelease(c, relDef)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", rel.Id)
}