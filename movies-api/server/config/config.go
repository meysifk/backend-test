package config

import (
	"log"
	"sync"

	"github.com/BurntSushi/toml"
)

var directories = []string{
	"./files/opt/stockbit/bin", // used at local dev
	".",                        // used at staging and production
}

type Config struct {
	Server Server
	DB     DB
	Omdb   Omdb `toml:"omdb"`
}

var config *Config
var once sync.Once
var appEnvName string

func InitConfig(env string) (*Config, error) {
	appEnvName = env
	var err error
	if config == nil {
		c := Config{}
		once.Do(func() {
			err = parseConfig(env, &c, "config", directories...)
		})
		if err != nil {
			return nil, err
		}
		config = &c
	}

	return config, nil
}

// GetConfig is getter for config instance
func GetConfig() *Config {
	if config != nil {
		return config
	}

	cfg, err := InitConfig(appEnvName)
	if err != nil {
		log.Fatal("error when InitConfig, error: ", err.Error())
	}

	return cfg
}

func parseConfig(env string, conf *Config, moduleName string, path ...string) error {

	err := readModuleConfig(conf, env, moduleName, path...)
	if err != nil {
		return err
	}

	return nil
}

func readModuleConfig(config *Config, env, module string, path ...string) error {
	var err error
	for _, val := range path {
		pathToFile := val + "/" + module + "-" + env + ".toml"

		_, err = toml.DecodeFile(pathToFile, &config)
		if err == nil {
			break
		}
	}
	return err
}
