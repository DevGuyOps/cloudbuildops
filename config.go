package cloudbuildops

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type TriggerConfig struct {
	Git struct {
		Provider string `yaml:"provider"`
		Project  string `yaml:"project"`
		Repo     string `yaml:"repo"`
	} `yaml:"git"`
	Triggers []struct {
		Name           string            `yaml:"name"`
		Disabled       bool              `yaml:"disabled"`
		Projectid      string            `yaml:"projectid"`
		Branchname     string            `yaml:"branchname"`
		Tagname        string            `yaml:"tagname"`
		ConfigFilename string            `yaml:"configfilename"`
		Substitutions  map[string]string `yaml:"substitutions"`
	} `yaml:"trigger"`
}

func ReadTriggerConfig(filename string) TriggerConfig {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var triggerConfig TriggerConfig
	err = yaml.Unmarshal(yamlFile, &triggerConfig)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return triggerConfig
}
