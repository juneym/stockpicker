package config

import (
	"encoding/json"
	"os"
)

func Load(configFile string) (Config, error) {

	file, _ := os.Open(configFile)
	decoder := json.NewDecoder(file)

	conf := Config{}
	err := decoder.Decode(&conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
