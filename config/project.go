package config

import (
	"cicd-lite/context"
	"cicd-lite/settings"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Project struct {
	Version string `yaml:"version"`
	Setup   struct {
		Git struct {
			SshKey string            `yaml:"ssh-key"`
			Remote map[string]string `yaml:"remote"`
		} `yaml:"git"`
	} `yaml:"setup"`
	Service map[string]struct {
		Image       string            `yaml:"image"`
		Environment map[string]string `yaml:"environment"`
		Command     string            `yaml:"command"`
	} `yaml:"service"`
	Pipeline struct {
		Clone map[string]struct {
			Remote string `yaml:"remote"`
			Branch string `yaml:"branch"`
		} `yaml:"clone"`
		Build []string `yaml:"build"`
	} `yaml:"pipeline"`
}

func init() {
	context.Initialize()
}

func GetProject() (Project, error) {
	data, err := ioutil.ReadFile(settings.APP.Path + "/etc/project.yml")

	if err != nil {
		panic(err)
	}

	project := Project{}

	err = yaml.Unmarshal([]byte(data), &project)

	return project, err
}

func GetSource(project *Project) ([]byte, error) {
	return yaml.Marshal(project)
}
