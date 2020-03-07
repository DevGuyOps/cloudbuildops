package cloudbuildops

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type TriggerConfig struct {
	Git      TriggerConfigGit       `yaml:"git"`
	Triggers []TriggerConfigTrigger `yaml:"trigger"`
}

type TriggerConfigGit struct {
	Provider string `yaml:"provider"`
	Project  string `yaml:"project"`
	Repo     string `yaml:"repo"`
}

type TriggerConfigTrigger struct {
	Name           string            `yaml:"name"`
	Disabled       bool              `yaml:"disabled"`
	Projectid      string            `yaml:"projectid"`
	Branchname     string            `yaml:"branchname"`
	Tagname        string            `yaml:"tagname"`
	ConfigFilename string            `yaml:"configfilename"`
	Substitutions  map[string]string `yaml:"substitutions"`
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

func WriteTriggerConfig(filename string, triggerConfig *TriggerConfig) error {
	conf, err := yaml.Marshal(&triggerConfig)
	if err != nil {
		return err
	}

	// Ensure dir exists
	os.MkdirAll(filepath.Dir(filename), os.ModePerm)

	err = ioutil.WriteFile(filename, conf, 0644)
	if err != nil {
		return err
	}

	return nil
}
