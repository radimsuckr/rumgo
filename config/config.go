package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

type Rule struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Path  string `json:"path,omitempty"`
}

type WatchlistItem struct {
	URL   string `json:"url"`
	Rules []Rule `json:"rules"`
}

type Config struct {
	Vesion    string          `json:"vesion"`
	Watchlist []WatchlistItem `json:"watchlist"`
}

func LoadConfig(path string) (config *Config, error error) {
	file, file_err := os.Open(path)
	if file_err != nil {
		return nil, file_err
	}

	content, content_err := io.ReadAll(file)
	file.Close()
	if content_err != nil {
		return nil, content_err
	}

	json.Unmarshal(content, &config)
	if config.Vesion != "0.1.0" {
		return nil, errors.New("config uses unsupported version")
	}

	return config, nil
}
