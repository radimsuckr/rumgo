// main package
package main

import (
	"log"
	"os"
	"time"

	"rumgo/internal/client"
	"rumgo/internal/config"
	"rumgo/internal/rules"
	"rumgo/internal/telegram"
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

func scrape(telegramConfig config.Telegram, item rules.WatchlistItem) {
	resp, err := client.SendRequest(item.URL)
	if err != nil {
		log.Printf("Failed sending request to: %s\n", item.URL)
	} else {
		for _, rule := range item.Rules {
			result, ruleErr := rule.Evaluate(&resp.Content)
			if ruleErr != nil {
				log.Printf("Rule error: %s\n", ruleErr)
				continue
			}

			if result {
				text := "Rule " + string(rule.GetType()) + " matched for " + item.URL
				telegram.SendTelegramMessage(telegramConfig, text)
			}

			log.Printf("%s (%s): %t\n", item.URL, rule.GetType(), result)
		}
	}
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
