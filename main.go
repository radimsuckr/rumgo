package main

import (
	"fmt"
	"os"
	"time"

	"radimsuckr/rumgo/config"
	"radimsuckr/rumgo/rules"
)

const (
	ERROR_LOADING_CONFIG       = 123
	ERROR_CREATING_WATCHLIST   = 124
	ERROR_READING_RUMGO_CONFIG = 125
)

func main() {
	cfg_path := os.Getenv("RUMGO_CONFIG")
	if len(cfg_path) < 1 {
		fmt.Println("set path to config file in \"RUMGO_CONFIG\" env var")
		os.Exit(ERROR_READING_RUMGO_CONFIG)
	}

	cfg, cfg_err := config.LoadConfig(cfg_path)
	if cfg_err != nil {
		fmt.Println(cfg_err)
		os.Exit(ERROR_LOADING_CONFIG)
	}

	watchlist, watchlist_error := rules.NewWatchlistFromConfig(cfg)
	if watchlist_error != nil {
		fmt.Println(watchlist_error)
		os.Exit(ERROR_CREATING_WATCHLIST)
	}

	for {
		for _, item := range watchlist.Items {
			go Scrape(item)
		}

		time.Sleep(cfg.LoopInterval * time.Second)
	}
}
