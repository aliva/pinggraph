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
	c := make([]Host, 0)
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
