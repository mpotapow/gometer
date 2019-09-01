package db

import (
	"encoding/json"
	"gometer/modules/core"
)

// Config ...
type Config struct {
	core.ConfigurationLoader
	DatabaseName string `json:"databaseName"`
}

// GetConfig ...
func GetConfig(configPath string) *Config {

	c := Config{}
	_ = json.Unmarshal(c.ConfigurationLoader.LoadFromFile(configPath+"/database.json"), &c)

	return &c
}
