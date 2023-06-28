package main

import (
	"fmt"
	"os"
	"time"

	"radimsuckr/rumgo/config"
	"radimsuckr/rumgo/rules"
)

const (
	ErrorLoadingConfig      = 123
	ErrorCreatingWatchlist  = 124
	ErrorReadingRumgoConfig = 125
)

func GetConfigPath() (cfg_path string) {
	cfg_path = os.Getenv("RUMGO_CONFIG")
	if len(cfg_path) < 1 {
		fmt.Println("set path to config file in \"RUMGO_CONFIG\" env var")
		os.Exit(ErrorReadingRumgoConfig)
	}
	return cfg_path
}

func LoadConfig(cfg_path string) (cfg *config.Config) {
	cfg, err := config.LoadConfig(cfg_path)
	if err != nil {
		fmt.Println(err)
		os.Exit(ErrorLoadingConfig)
	}
	return cfg
}

func CreateWatchlist(cfg *config.Config) (watchlist *rules.Watchlist) {
	watchlist, err := rules.NewWatchlistFromConfig(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(ErrorCreatingWatchlist)
	}
	return watchlist
}

func main() {
	cfg_path := GetConfigPath()
	cfg := LoadConfig(cfg_path)
	watchlist := CreateWatchlist(cfg)

	for {
		for _, item := range watchlist.Items {
			go Scrape(item)
		}

		time.Sleep(cfg.LoopInterval * time.Second)
	}
}
