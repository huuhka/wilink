package old

import (
	"context"
	"github.com/huuhka/wilink/adotool"
	"github.com/microsoft/azure-devops-go-api/azuredevops/release"
	"log"
	"os"
)

func main2() {
	organizationUrl := "https://dev.azure.com/my" // todo: replace value with your organization url
	//personalAccessToken := "my" // todo: replace value with your PAT

	authToken := ""

	//connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)
	connection := adotool.NewOauthConnection(organizationUrl, authToken)

	ctx := context.Background()

	crClient, err := release.NewClient(ctx, connection)
	if err != nil {
		log.Println(err)
	}

	project := "Zure0100"
	relId := 1

	relDef, err := adotool.GetReleaseDefinition(crClient, project, relId)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	rel, err := adotool.CreateRelease(crClient, relDef)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = rel.StartStage(crClient, rel.Stages[0].Name)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Printf("%+v", rel.Stages[0])
}