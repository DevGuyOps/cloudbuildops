package main

import (
	"fmt"
	"log"

	"github.com/GuySWatson/cloudbuildops"
	"google.golang.org/api/cloudbuild/v1"
)

func main() {
	cb := cloudbuildops.CB{}
	cb.Init()

	triggerConfig := cloudbuildops.ReadTriggerConfig("example.yml")
	repoName := fmt.Sprintf("%s_%s_%s",
		triggerConfig.Git.Provider,
		triggerConfig.Git.Project,
		triggerConfig.Git.Repo,
	)

	for _, trigger := range triggerConfig.Triggers {
		// Check repo exists
		_, err := cb.GetRepo(repoName, trigger.Projectid)
		if err != nil {
			log.Panicf("Error getting repo (Probably doesn't exist): %s", err)
		}

		// Check trigger already exists
		currentTrigger, err := cb.GetTrigger(repoName, trigger.Name, trigger.Projectid)
		if err != nil {
			log.Panicf("Error getting repo (Probably doesn't exist): %s", err)
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
				log.Panicf("Error creating trigger %s: %s", trigger.Name, err)
			}

			log.Printf("Trigger created: %s", trigger.Name)
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
				log.Panicf("Error updating trigger %s: %s", trigger.Name, err)
			}

			log.Printf("Trigger updated: %s", trigger.Name)
		}
	}
}
