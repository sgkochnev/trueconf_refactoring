package config

import (
	"errors"
	"os"
)

type Config struct {
	HTTP
	JsonStore
}

type HTTP struct {
	Port string
}

type JsonStore struct {
	Name string
}

var ErrJsonStoreNotFound = errors.New("json store not found")

func New() (*Config, error) {
	cfg := &Config{
		HTTP:      HTTP{Port: os.Getenv("HTTP_PORT")},
		JsonStore: JsonStore{Name: os.Getenv("JSON_STORE_NAME")},
	}

	if cfg.JsonStore.Name == "" {
		return nil, ErrJsonStoreNotFound
	}

	return cfg, nil
}
