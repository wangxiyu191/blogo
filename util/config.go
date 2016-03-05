package util

import (
	//"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type LinkM struct {
	Url         string `yaml:"url"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type ConfigM struct {
	Domain      string  `yaml:"domain"`
	ArticlePath string  `yaml:"article_path"`
	GitRepo     string  `yaml:"git_repo"`
	Branch      string  `yaml:"branch"`
	Delete      bool    `yaml:"delete"`
	Links       []LinkM `yaml:"links"`
}

var (
	Config ConfigM
)

func (config *ConfigM) Load() {
	workDir, err := os.Getwd()
	data, err := ioutil.ReadFile(workDir + "/config.yaml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(data, config)
}
