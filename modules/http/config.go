package http

import (
	"encoding/json"
	"gometer/modules/core"
)

// Config ...
type Config struct {
	core.ConfigurationLoader
	Host string `json:"host"`
	Port int    `json:"port"`
}

// GetConfig ...
func GetConfig(configPath string) *Config {

	c := Config{}
	_ = json.Unmarshal(c.ConfigurationLoader.LoadFromFile(configPath+"/server.json"), &c)

	return &c
}
