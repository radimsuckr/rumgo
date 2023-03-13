package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"time"
)

const DEFAULT_LOOP_INTERVAL = 60

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
	Version      string          `json:"version"`
	Watchlist    []WatchlistItem `json:"watchlist"`
	LoopInterval time.Duration   `json:"loopInterval,omitempty"`
}

func LoadConfig(path string) (*Config, error) {
	file, file_err := os.Open(path)
	if file_err != nil {
		return nil, file_err
	}

	content, content_err := io.ReadAll(file)
	file.Close()
	if content_err != nil {
		return nil, content_err
	}

	config := &Config{LoopInterval: DEFAULT_LOOP_INTERVAL}
	json.Unmarshal(content, &config)
	if config.Version != "0.1.0" {
		return nil, errors.New("config uses unsupported version")
	}

	return config, nil
}
