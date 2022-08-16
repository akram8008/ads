package config

import (
	"ads/internal/domain"
	"encoding/json"
	"io/ioutil"
	"log"
)

func New(F string) *domain.Config {
	byteValue, err := ioutil.ReadFile(F)
	if err != nil {
		log.Fatalf("%v", err)
	}
	var config *domain.Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return config
}
