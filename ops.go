package cloudbuildops

import (
	"context"
	"fmt"

	"google.golang.org/api/cloudbuild/v1"
	"google.golang.org/api/sourcerepo/v1"
)

type CB struct {
	cloudBuild *cloudbuild.Service
	sourceRepo *sourcerepo.Service
}

func (cb *CB) Init() error {
	ctx := context.Background()

	// Connect to Cloud Build
	cloudbuildService, err := cloudbuild.NewService(ctx)
	if err != nil {
		return err
	}
	cb.cloudBuild = cloudbuildService

	// Connect to Source Repositories
	sourcerepoService, err := sourcerepo.NewService(ctx)
	if err != nil {
		return err
	}
	cb.sourceRepo = sourcerepoService

	return nil
}

func (cb *CB) GetRepo(name, projectID string) (*sourcerepo.Repo, error) {
	repo, err := cb.sourceRepo.Projects.Repos.Get(fmt.Sprintf("projects/%s/repos/%s", projectID, name)).Do()
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (cb *CB) CreateTrigger(trigger *cloudbuild.BuildTrigger) error {
	_, err := cb.cloudBuild.Projects.Triggers.Create(trigger.TriggerTemplate.ProjectId, trigger).Do()
	if err != nil {
		return err
	}

	return nil
}

func (cb *CB) GetTrigger(repoName, triggerName, projectID string) (*cloudbuild.BuildTrigger, error) {
	triggers, err := cb.cloudBuild.Projects.Triggers.List(projectID).Do()
	if err != nil {
		return nil, err
	}

	for _, trigger := range triggers.Triggers {
		if repoName == trigger.TriggerTemplate.RepoName && triggerName == trigger.Name {
			return trigger, nil
		}
	}

	return nil, nil
}

func (cb *CB) UpdateTrigger(filename, projectID, triggerID string, trigger *cloudbuild.BuildTrigger) error {
	_, err := cb.cloudBuild.Projects.Triggers.Patch(projectID, triggerID, trigger).Do()
	if err != nil {
		return err
	}

	return nil
}
