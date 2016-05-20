package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type configFile struct {
	DBPath string `json:"db_path"`
}

func readConfig(path string) (*configFile, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	conf := configFile{}
	if err := json.Unmarshal(b, &conf); err != nil {
		return nil, err
	}

	return &conf, err
}
