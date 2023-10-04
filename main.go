package main

import (
	"log"
	"os"
	"time"

	"rumgo/internal/config"
	"rumgo/internal/rules"
)

const (
	ErrorCreatingWatchlist  = 100
	ErrorLoadingConfig      = 101
	ErrorReadingConfigFile  = 102
	ErrorReadingRumgoConfig = 103
)

func GetConfigPath() (cfg_path string) {
	cfg_path = os.Getenv("RUMGO_CONFIG")
	if len(cfg_path) < 1 {
		log.Println("set path to config file in \"RUMGO_CONFIG\" env var")
		os.Exit(ErrorReadingRumgoConfig)
	}
	return cfg_path
}

func CreateWatchlist(cfg config.Config) (watchlist rules.Watchlist) {
	watchlist, err := rules.NewWatchlistFromConfig(cfg)
	if err != nil {
		log.Println(err)
		os.Exit(ErrorCreatingWatchlist)
	}
	return watchlist
}

func main() {
	cfg_path := GetConfigPath()
	cfg_content, err := config.ReadConfigFile(cfg_path)
	if err != nil {
		log.Println(err)
		os.Exit(ErrorReadingConfigFile)
	}
	cfg, err := config.NewConfig(cfg_content)
	if err != nil {
		log.Println(err)
		os.Exit(ErrorLoadingConfig)
	}
	watchlist := CreateWatchlist(cfg)

	for {
		for _, item := range watchlist.Items {
			go Scrape(cfg.Telegram, item)
		}

		time.Sleep(cfg.LoopInterval * time.Second)
	}
}
