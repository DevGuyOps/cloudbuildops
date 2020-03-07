package main

import (
	"log"
	"os"

	"github.com/GuySWatson/cloudbuildops"
	"github.com/urfave/cli"
)

func main() {
	cb := cloudbuildops.CB{}
	cb.Init()

	app := &cli.App{
		Name:        "Cloud Build Ops",
		Description: "Cloud Build Ops is a tool that lets you manage Cloud Build pipeline configuration from yaml files which makes managing Cloud Build much easier and faster.",
		Commands: []cli.Command{
			{
				Name:  "get",
				Usage: "Write all existing cloud build pipelines to file",
				Action: func(c *cli.Context) error {
					err := cb.Get(c.String("p"), c.String("o"))
					if err != nil {
						return err
					}

					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "projectid,p",
						Value:    "",
						Usage:    "Project ID of the GCP project",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "output,o",
						Value:    "",
						Usage:    "Output directory to publish config files",
						Required: true,
					},
				},
			},
			{
				Name:  "push",
				Usage: "Create/Update cloud build pipelines from the proveided config files",
				Action: func(c *cli.Context) error {
					triggers := []cloudbuildops.TriggerConfig{}
					for _, filename := range c.Args() {
						triggers = append(triggers, cloudbuildops.ReadTriggerConfig(filename))
					}

					err := cb.Push(triggers)
					if err != nil {
						return err
					}

					return nil
				},
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:     "config,c",
						Usage:    "Path to config files (Supports wildcards)",
						Required: true,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
