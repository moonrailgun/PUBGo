package config

import (
	"github.com/jinzhu/configor"
	"path"
	"path/filepath"
	"runtime"
)

var Config = struct {
	DB struct {
		Name string `env:"DBName" default:"pubgo"`
		Host string `env:"DBHost" default:"localhost"`
		Port string `env:"DBPort" default:"3306"`
		User string `env:"DBUser"`
		Pass string `env:"DBPassword"`
	}
	PUBG struct {
		API string `env:"PUBG_API"`
	}
}{}

func init() {
	configPath := filepath.Join(getCurrentPath(), "./config.yml")
	if err := configor.Load(&Config, configPath); err != nil {
		panic(err)
	}
}

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
