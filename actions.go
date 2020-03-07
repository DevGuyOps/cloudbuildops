package cloudbuildops

import (
	"fmt"
	"strings"

	"google.golang.org/api/cloudbuild/v1"
)

func (cb *CB) Push(triggerConfigList []TriggerConfig) error {
	for _, triggerConfig := range triggerConfigList {
		repoName := fmt.Sprintf("%s_%s_%s",
			triggerConfig.Git.Provider,
			triggerConfig.Git.Project,
			triggerConfig.Git.Repo,
		)

		for _, trigger := range triggerConfig.Triggers {
			// Check repo exists
			_, err := cb.GetRepo(repoName, trigger.Projectid)
			if err != nil {
				return fmt.Errorf("Error getting repo (Probably doesn't exist): %s", err)
			}

			// Check trigger already exists
			currentTrigger, err := cb.GetTrigger(repoName, trigger.Name, trigger.Projectid)
			if err != nil {
				return fmt.Errorf("Error getting repo (Probably doesn't exist): %s", err)
			}

			if currentTrigger == nil {
				// Add trigger
				err := cb.CreateTrigger(&cloudbuild.BuildTrigger{
					Name:          trigger.Name,
					Disabled:      trigger.Disabled,
					Filename:      trigger.ConfigFilename,
					Substitutions: trigger.Substitutions,
					TriggerTemplate: &cloudbuild.RepoSource{
						BranchName: trigger.Branchname,
						TagName:    trigger.Tagname,
						ProjectId:  trigger.Projectid,
						RepoName:   repoName,
					},
				})
				if err != nil {
					return fmt.Errorf("Error creating trigger %s: %s", trigger.Name, err)
				}

				fmt.Printf("Trigger created: %s\n", trigger.Name)
			} else {
				// Update trigger
				err := cb.UpdateTrigger(trigger.ConfigFilename, trigger.Projectid,
					currentTrigger.Id, &cloudbuild.BuildTrigger{
						Name:          currentTrigger.Name,
						Disabled:      trigger.Disabled,
						Filename:      trigger.ConfigFilename,
						Substitutions: trigger.Substitutions,
						TriggerTemplate: &cloudbuild.RepoSource{
							BranchName: trigger.Branchname,
							TagName:    trigger.Tagname,
							ProjectId:  trigger.Projectid,
							RepoName:   repoName,
						},
					})
				if err != nil {
					return fmt.Errorf("Error updating trigger %s: %s", trigger.Name, err)
				}

				fmt.Printf("Trigger updated: %s\n", trigger.Name)
			}
		}
	}

	return nil
}

func (cb *CB) Get(projectID, outputDir string) error {
	triggerList, err := cb.GetTriggerList(projectID)
	if err != nil {
		return err
	}

	// Group triggers by repo
	repoTriggersMap := map[string][]cloudbuild.BuildTrigger{}
	for _, trigger := range triggerList {
		repoTriggersMap[trigger.TriggerTemplate.RepoName] = append(repoTriggersMap[trigger.TriggerTemplate.RepoName], *trigger)
	}

	for repoName, repoTriggers := range repoTriggersMap {
		repoNameSplit := strings.Split(repoName, "_")

		triggerConfigs := []TriggerConfigTrigger{}
		for _, repoTrigger := range repoTriggers {
			triggerConfigs = append(triggerConfigs, TriggerConfigTrigger{
				Name:           repoTrigger.Name,
				Disabled:       repoTrigger.Disabled,
				Projectid:      repoTrigger.TriggerTemplate.ProjectId,
				Branchname:     repoTrigger.TriggerTemplate.BranchName,
				Tagname:        repoTrigger.TriggerTemplate.TagName,
				ConfigFilename: repoTrigger.Filename,
				Substitutions:  repoTrigger.Substitutions,
			})
		}

		// Prepare triggers
		triggerConfig := &TriggerConfig{
			Git: TriggerConfigGit{
				Provider: repoNameSplit[0],
				Project:  repoNameSplit[1],
				Repo:     repoNameSplit[2],
			},
			Triggers: triggerConfigs,
		}

		err = WriteTriggerConfig(fmt.Sprintf("%s/%s.yml", outputDir, repoName), triggerConfig)
		if err != nil {
			return fmt.Errorf("Error writing trigger config %s: %s", repoName, err)
		}

		fmt.Printf("Wrote %s to file\n", repoName)
	}

	return nil
}
