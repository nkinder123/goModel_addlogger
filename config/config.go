package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Mode       string `json:"mode"`
	Port       int    `json:"port"`
	*LogConfig `json:"log"`
}

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

var Conf = new(Config)

func Init(url string) (err error) {
	b, errs := os.ReadFile(url)
	if errs != nil {
		return errs
	}
	return json.Unmarshal(b, Conf)
}
