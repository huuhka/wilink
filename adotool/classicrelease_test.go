package adotool_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/huuhka/wilink/adotool"
	"github.com/huuhka/wilink/helpers"
	"github.com/microsoft/azure-devops-go-api/azuredevops/release"
	"strings"
	"testing"
)

type MockClient struct {
	release.Client
	returnErr error
}

func (c MockClient) GetReleaseDefinition(ctx context.Context, args release.GetReleaseDefinitionArgs) (*release.ReleaseDefinition, error) {
	_ = ctx

	if c.returnErr != nil {
		return &release.ReleaseDefinition{}, c.returnErr
	}

	n := fmt.Sprintf("Definition %d", 1)

	return &release.ReleaseDefinition{
		Id:   args.DefinitionId,
		Name: &n,
	}, nil
}

func TestGetReleaseDefinition(t *testing.T) {
	tcs := []struct {
		name      string
		inputProj string
		inputID   int
		returnErr error
		expected  adotool.ReleaseDefinition
	}{
		{
			"should get classic release definition",
			"success-proj",
			1,
			nil,
			adotool.ReleaseDefinition{
				Id:      1,
				Project: "success-proj",
				Name:    "Definition 1",
			},
		},
		{
			"should fail on not found",
			"failure-proj",
			1,
			errors.New("release definition not found"),
			adotool.ReleaseDefinition{},
		},
	}

	for _, tc := range tcs {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// Create client
			client := &MockClient{
				returnErr: tc.returnErr,
			}

			got, err := adotool.GetReleaseDefinition(client, tc.inputProj, tc.inputID)
			if err != nil && !strings.Contains(err.Error(), tc.returnErr.Error()) {
				t.Errorf("expected error %q, got %q", tc.returnErr.Error(), err.Error())
			}

			// Compare request with expected
			if tc.returnErr == nil && !cmp.Equal(got, tc.expected) {
				t.Errorf("Test: %s, Got: %v, Expected: %v", tc.name, got, tc.expected)
			}
		})
	}
}

func (c MockClient) CreateRelease(ctx context.Context, args release.CreateReleaseArgs) (*release.Release, error) {
	_ = ctx

	if c.returnErr != nil {
		return &release.Release{}, c.returnErr
	}

	name := fmt.Sprintf("Release %d", *args.ReleaseStartMetadata.DefinitionId)
	stages := []release.ReleaseEnvironment{
		{
			Name:   helpers.Pointer("Stage1"),
			Id:     args.ReleaseStartMetadata.DefinitionId,
			Status: &release.EnvironmentStatusValues.NotStarted,
		},
	}
	return &release.Release{
		Id:           args.ReleaseStartMetadata.DefinitionId,
		Name:         &name,
		Environments: &stages,
	}, nil
}

func TestCreateRelease(t *testing.T) {
	tcs := []struct {
		name      string
		input     adotool.ReleaseDefinition
		returnErr error
		expected  adotool.Release
	}{
		{
			"should create new instance of a classic release with stages populated",
			adotool.ReleaseDefinition{
				Project: "success-proj",
				Name:    "Release 1",
				Id:      1,
			},
			nil,
			adotool.Release{
				Id:      1,
				Project: "success-proj",
				Name:    "Release 1",
				Stages: []adotool.Stage{
					{
						Id:     1,
						Name:   "Stage1",
						Status: release.EnvironmentStatusValues.NotStarted,
					},
				},
			},
		},
		{
			"should fail when given release definition is not found",
			adotool.ReleaseDefinition{
				Project: "failure-proj",
				Name:    "Release 1",
				Id:      1,
			},
			errors.New("release definition not found"),
			adotool.Release{},
		},
		// TODO: Make this test work and apply for other calls too.
		//{
		//	"should not panic when azure devops package produces nil pointers",
		//	adotool.ReleaseDefinition{
		//		Project: "failure-proj",
		//		Id:      1,
		//	},
		//	errors.New("received a nil pointer"),
		//	adotool.Release{},
		//},
	}

	for _, tc := range tcs {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// Create client
			client := &MockClient{
				returnErr: tc.returnErr,
			}

			got, err := adotool.CreateRelease(client, tc.input)
			if err != nil && !strings.Contains(err.Error(), tc.returnErr.Error()) {
				t.Fatalf("expected error %q, got %q", tc.returnErr.Error(), err.Error())
			}

			// Compare request with expected
			if tc.returnErr == nil && !cmp.Equal(got, tc.expected) {
				t.Errorf("Test: %s, Got: %v, Expected: %v", tc.name, got, tc.expected)
			}
		})
	}
}

func (c MockClient) UpdateReleaseEnvironment(ctx context.Context, args release.UpdateReleaseEnvironmentArgs) (*release.ReleaseEnvironment, error) {
	_ = ctx

	if c.returnErr != nil {
		return &release.ReleaseEnvironment{}, c.returnErr
	}

	return &release.ReleaseEnvironment{
		Id:     args.EnvironmentId,
		Status: &release.EnvironmentStatusValues.InProgress,
		Name:   helpers.Pointer("Stage1"),
	}, nil
}

func TestRelease_StartStage(t *testing.T) {
	tcs := []struct {
		name              string
		input             adotool.Release
		inputStageToStart string
		returnErr         error
		expected          adotool.Stage
	}{
		{
			"should set the value of the stage to in progress",
			adotool.Release{
				Id:      1,
				Project: "success-proj",
				Name:    "Release 1",
				Stages: []adotool.Stage{
					{
						Id:     1,
						Name:   "Stage1",
						Status: release.EnvironmentStatusValues.NotStarted,
					},
				},
			},
			"Stage1",
			nil,
			adotool.Stage{
				Id:     1,
				Name:   "Stage1",
				Status: release.EnvironmentStatusValues.InProgress,
			},
		},
		{
			"should fail when given stage is not found in release",
			adotool.Release{
				Id:      1,
				Project: "failure-proj",
				Name:    "Release 1",
				Stages: []adotool.Stage{
					{
						Id:     1,
						Name:   "Stage1",
						Status: release.EnvironmentStatusValues.NotStarted,
					},
				},
			},
			"Stage2",
			fmt.Errorf("could not find stage %s in release %s", "Stage2", "Release 1"),
			adotool.Stage{},
		},
	}

	for _, tc := range tcs {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// Create client
			client := &MockClient{
				returnErr: tc.returnErr,
			}

			err := tc.input.StartStage(client, tc.inputStageToStart)
			if err != nil && !strings.Contains(err.Error(), tc.returnErr.Error()) {
				t.Fatalf("expected error %q, got %q", tc.returnErr.Error(), err.Error())
			}

			// Compare request with expected
			if tc.returnErr == nil && !cmp.Equal(tc.input.Stages[0], tc.expected) {
				t.Errorf("Test: %s, Got: %v, Expected: %v", tc.name, tc.input.Stages[0], tc.expected)
			}
		})
	}
}