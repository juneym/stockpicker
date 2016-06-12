package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Symbols       []string
	StorageFile   string
	FetchInterval int
}

func Load(configFile string) (Config, error) {

	file, _ := os.Open(configFile)
	decoder := json.NewDecoder(file)

	conf := Config{}
	err := decoder.Decode(&conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
