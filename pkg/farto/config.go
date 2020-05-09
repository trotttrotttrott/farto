package farto

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type config struct {
	S3Region string `yaml:"s3Region"`
	S3Bucket string `yaml:"s3Bucket"`
	S3Prefix string `yaml:"s3Prefix"`
}

func getConfig() (c config, err error) {
	y, err := ioutil.ReadFile("farto.yaml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(y, &c)
	return
}
