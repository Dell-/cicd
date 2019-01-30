package config

type Project struct {
	Version string `yaml:"version"`
	Setup   struct {
		Git struct {
			SshKey string `yaml:"ssh-key"`
		} `yaml:"git"`
	} `yaml:"setup"`
	Services map[string]struct {
		Image       string            `yaml:"image"`
		Environment map[string]string `yaml:"environment"`
		Command     string            `yaml:"command"`
		Links       []string          `yaml:"links"`
	} `yaml:"services"`
	Pipelines map[string][]string `yaml:"pipelines"`
}
