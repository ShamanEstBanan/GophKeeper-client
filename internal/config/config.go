package config

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

const (
	tempDir   = "keeper"
	tokenFile = "token.txt"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	Directory     string
	TokenFile     string
}

var cfg Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {

		cfg.Directory = os.TempDir() + tempDir
		cfg.TokenFile = tokenFile

		flag.StringVar(
			&cfg.ServerAddress,
			"server",
			":3080",
			"адрес и порт сервера",
		)

		err := os.Mkdir(cfg.Directory, os.FileMode(0777))
		if err != nil {
			if !errors.Is(err, os.ErrExist) {
				log.Fatal("Error creating config file in temp:", err)
			}
		}
		data, err := json.Marshal(cfg)
		if err != nil {
			log.Fatal("ERROR while marshal storage")
		}
		configFile := cfg.Directory + "/config.json"
		err = ioutil.WriteFile(configFile, data, 0o600)
		if err != nil {
			log.Fatal("Error init config file")
		}
	})
	return &cfg
}
