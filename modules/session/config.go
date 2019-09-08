package session

import (
	"encoding/json"
	"gometer/modules/core"
)

// Config ...
type Config struct {
	core.ConfigurationLoader
	Driver   string `json:"driver"`
	Lifetime int    `json:"lifetime"`
	Cookie   string `json:"cookie"`
	Domain   string `json:"domain"`
	Path     string `json:"path"`
	Secure   bool   `json:"secure"`
	HttpOnly bool   `json:"http_only"`
}

// GetConfig ...
func GetConfig(configPath string) *Config {

	c := Config{}
	_ = json.Unmarshal(c.ConfigurationLoader.LoadFromFile(configPath+"/session.json"), &c)

	return &c
}
