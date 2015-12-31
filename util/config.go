package util

import (
	//"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type ConfigM struct {
	WorkDir     string
	Domain      string `yaml:"domain"`
	ArticlePath string `yaml:"article_path"`
	GitRepo     string `yaml:"git_repo"`
	Branch      string `yaml:"branch"`
	Delete      bool   `yaml:"delete"`
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
