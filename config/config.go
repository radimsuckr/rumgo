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

func ReadConfigFile(path string) (content []byte, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err = io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func NewConfig(content []byte) (config *Config, err error) {
	config = &Config{LoopInterval: DEFAULT_LOOP_INTERVAL}
	if json.Unmarshal(content, &config) != nil {
		return nil, errors.New("file does not contain valid JSON")
	}
	if config.Version != "0.1.0" {
		return nil, errors.New("config uses unsupported version")
	}
	return config, nil
}
