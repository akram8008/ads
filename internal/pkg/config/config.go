package config

import (
	"ads/internal/model"
	"encoding/json"
	"io/ioutil"
	"log"
)

func New(F string) *model.Config {
	byteValue, err := ioutil.ReadFile(F)
	if err != nil {
		log.Fatalf("%v", err)
	}
	var config *model.Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return config
}
