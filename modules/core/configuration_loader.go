package core

import (
	"io/ioutil"
	"os"
)

// ConfigurationLoader ...
type ConfigurationLoader struct {
}

// LoadFromFile ...
func (c *ConfigurationLoader) LoadFromFile(filename string) []byte {

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return []byte{}
	}

	file, _ := ioutil.ReadFile(filename)
	return []byte(file)
}
