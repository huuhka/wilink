package adowrappers

import (
	"context"
	"fmt"
	"github.com/microsoft/azure-devops-go-api/azuredevops/release"
)

// Create Release
// Get Release
// Set Artifacts (release, artifacts array)
// Start Stage by Name for Specific Release

type ReleaseDefinition struct {
	Id      int
	Name    string
	Project string
}

type Release struct {
	Id      int
	Name    string
	Project string
	Stages  []Stage
}

type Stage struct {
	Id     int
	Name   string
	Status release.EnvironmentStatus
}

// GetReleaseDefinition gets a classic release definition by a given Id
func GetReleaseDefinition(client release.Client, project string, id int) (ReleaseDefinition, error) {
	rel, err := getReleaseDefinition(project, id, client)

	if err != nil {
		return ReleaseDefinition{}, fmt.Errorf("could not get release by ID %d in project %s, Error: %w", id, project, err)
	}

	return ReleaseDefinition{
		Id:      *rel.Id,
		Name:    *rel.Name,
		Project: project,
	}, err
}

// CreateRelease creates an instance of a classic release
func CreateRelease(client release.Client, rel ReleaseDefinition) (Release, error) {

	relStartData := release.ReleaseStartMetadata{
		DefinitionId: &rel.Id,
	}

	createArgs := release.CreateReleaseArgs{
		ReleaseStartMetadata: &relStartData,
		Project:              &rel.Project,
	}

	ctx := context.Background()
	r, err := client.CreateRelease(ctx, createArgs)
	if err != nil {
		return Release{}, fmt.Errorf("could not create release instance for Id %d, Error: %w", rel.Id, err)
	}
	return Release{
		Id:      *r.Id,
		Name:    *r.Name,
		Project: rel.Project,
		Stages:  parseStages(r),
	}, nil
}

// StartStage starts a classic release stage by a given StageId (EnvironmentId in Azure DevOps API)
func (r *Release) StartStage(client release.Client, name string) error {
	var stageToStart *Stage
	for i := range r.Stages {
		if r.Stages[i].Name == name {
			stageToStart = &r.Stages[i]
			break
		}
	}
	if stageToStart == nil {
		return fmt.Errorf("could not find stage %s in release %s", name, r.Name)
	}

	envUpdate := release.ReleaseEnvironmentUpdateMetadata{
		Status: &release.EnvironmentStatusValues.InProgress,
	}

	ctx := context.Background()
	env, err := client.UpdateReleaseEnvironment(ctx, release.UpdateReleaseEnvironmentArgs{
		EnvironmentUpdateData: &envUpdate,
		Project:               &r.Project,
		ReleaseId:             &r.Id,
		EnvironmentId:         &stageToStart.Id,
	})
	if err != nil {
		return fmt.Errorf("could not start stage %q for release %d, Error: %w", name, r.Id, err)
	}
	*stageToStart = *parseStage(env)

	return nil
}

// GetRelease gets a classic release by a given Id
func GetRelease(client release.Client, releaseId int, project string) (Release, error) {
	args := release.GetReleaseArgs{
		Project:   &project,
		ReleaseId: &releaseId,
	}

	ctx := context.Background()

	rel, err := client.GetRelease(ctx, args)
	if err != nil {
		return Release{}, fmt.Errorf("could not get release %d, Error: %w", releaseId, err)
	}

	return Release{
		Id:      *rel.Id,
		Name:    *rel.Name,
		Project: *rel.ProjectReference.Name,
		Stages:  parseStages(rel),
	}, err
}

// getReleaseDefinition gets a classic release definition by a given Id as represented by the Azure DevOps API
func getReleaseDefinition(project string, relDefId int, client release.Client) (*release.ReleaseDefinition, error) {
	args := release.GetReleaseDefinitionArgs{
		Project:      &project,
		DefinitionId: &relDefId,
	}

	ctx := context.Background()
	rel, err := client.GetReleaseDefinition(ctx, args)

	return rel, err
}

func parseStages(rel *release.Release) []Stage {
	if rel.Environments == nil {
		return []Stage{}
	}

	stages := make([]Stage, len(*rel.Environments))
	for i := range *rel.Environments {
		env := (*rel.Environments)[i]
		stages[i] = Stage{
			Id:     *env.Id,
			Name:   *env.Name,
			Status: *env.Status,
		}
	}
	return stages
}

func parseStage(env *release.ReleaseEnvironment) *Stage {
	return &Stage{
		Id:     *env.Id,
		Name:   *env.Name,
		Status: *env.Status,
	}
}