package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// LoadConfig docs
func LoadConfig(path string) []Host {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	config := make([]Host, 0)
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	for i := range config {
		if config[i].Port == 0 {
			config[i].Port = 22
		}
	}

	return config
}
