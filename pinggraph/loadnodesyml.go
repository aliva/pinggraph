package pinggraph

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// LoadNodes loads yaml file containing a list of nodes
func LoadNodes(path string) []Node {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("yamlFile.Get err   #%v ", err))
	}

	c := make([]Node, 0)
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
