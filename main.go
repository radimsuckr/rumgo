// main package
package main

import (
	"log"
	"os"
	"time"

	"rumgo/internal/config"
	"rumgo/internal/rules"
)

const (
	errorCreatingWatchlist  = 100
	errorLoadingConfig      = 101
	errorReadingConfigFile  = 102
	errorReadingRumgoConfig = 103
)

func getConfigPath() (cfgPath string) {
	cfgPath = os.Getenv("RUMGO_CONFIG")
	if len(cfgPath) < 1 {
		log.Println("set path to config file in \"RUMGO_CONFIG\" env var")
		os.Exit(errorReadingRumgoConfig)
	}
	return cfgPath
}

func createWatchlist(cfg config.Config) (watchlist rules.Watchlist) {
	watchlist, err := rules.NewWatchlistFromConfig(cfg)
	if err != nil {
		log.Println(err)
		os.Exit(errorCreatingWatchlist)
	}
	return watchlist
}

func main() {
	cfgPath := getConfigPath()
	cfgContent, err := config.ReadConfigFile(cfgPath)
	if err != nil {
		log.Println(err)
		os.Exit(errorReadingConfigFile)
	}
	cfg, err := config.NewConfig(cfgContent)
	if err != nil {
		log.Println(err)
		os.Exit(errorLoadingConfig)
	}
	watchlist := createWatchlist(cfg)

	for {
		for _, item := range watchlist.Items {
			go scrape(cfg.Telegram, item)
		}

		time.Sleep(cfg.LoopInterval * time.Second)
	}
}
