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

func main() {
	cfg_path := os.Getenv("RUMGO_CONFIG")
	if len(cfg_path) < 1 {
		fmt.Println("set path to config file in \"RUMGO_CONFIG\" env var")
		os.Exit(ErrorReadingRumgoConfig)
	}

	cfg, err := config.LoadConfig(cfg_path)
	if err != nil {
		fmt.Println(err)
		os.Exit(ErrorLoadingConfig)
	}

	watchlist, err := rules.NewWatchlistFromConfig(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(ErrorCreatingWatchlist)
	}

	for {
		for _, item := range watchlist.Items {
			go Scrape(item)
		}

		time.Sleep(cfg.LoopInterval * time.Second)
	}
}
