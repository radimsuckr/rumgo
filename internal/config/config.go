// Package config handles all program's configuration
package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"time"
)

const defaultLoopInterval = 60

// Rule is a single rule configuration of what to do with the crawler response
type Rule struct {
	Type   string `json:"type"`
	Invert bool   `json:"invert"`
	Value  string `json:"value"`
	Path   string `json:"path,omitempty"`
}

// Telegram holds configuration for Telegram bot
type Telegram struct {
	Channel string `json:"channel"`
	Token   string `json:"token"`
}

// WatchlistItem is a single URL with associated rules
type WatchlistItem struct {
	URL   string `json:"url"`
	Rules []Rule `json:"rules"`
}

// Config is a whole configuration struct
type Config struct {
	Version      string          `json:"version"`
	Telegram     Telegram        `json:"telegram"`
	Watchlist    []WatchlistItem `json:"watchlist"`
	LoopInterval time.Duration   `json:"loopInterval,omitempty"`
}

// ReadConfigFile loads config from a file
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

// NewConfig creates a new Config struct from file content
func NewConfig(content []byte) (config Config, err error) {
	config = Config{LoopInterval: defaultLoopInterval}
	if json.Unmarshal(content, &config) != nil {
		return Config{}, errors.New("file does not contain valid JSON")
	}
	if config.Version != "0.1.0" {
		return Config{}, errors.New("config uses unsupported version")
	}
	return config, nil
}
