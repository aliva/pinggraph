package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func loadNodes(path string) []node {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("yamlFile.Get err   #%v ", err))
	}

	c := make([]node, 0)
	err = yaml.Unmarshal(f, &c)
	if err != nil {
		panic(fmt.Sprintf("Unmarshal: %v", err))
	}

	for i := range c {
		if c[i].Port == 0 {
			c[i].Port = 22
		}
		if c[i].User == "" {
			c[i].User = "root"
		}
		if c[i].Name == "" {
			c[i].Name = c[i].Host
		}
	}

	return c
}
