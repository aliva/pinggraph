package main

import (
	"fmt"
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
		if config[i].User == "" {
			config[i].User = "root"
		}
		if config[i].Name == "" {
			config[i].Name = config[i].Host
		}
	}
	fmt.Println(config)

	return config
}
