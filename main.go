package main

import (
	"fmt"
	"os"
	"time"

	"radimsuckr/rumgo/client"
	"radimsuckr/rumgo/config"
	"radimsuckr/rumgo/rules"
)

const (
	ERROR_LOADING_CONFIG     = 123
	ERROR_CREATING_WATCHLIST = 124
)

func Scrape(item rules.WatchlistItem) {
	resp, resp_err := client.SendRequest(item.URL)
	if resp_err != nil {
		fmt.Println(resp_err)
	}

	for _, rule := range item.Rules {
		result, rule_err := rule.Evaluate(&resp.Content)
		if rule_err != nil {
			fmt.Printf("Rule error: %s\n", rule_err)
			continue
		}
		fmt.Printf("%s (%s): %t\n", item.URL, rule.GetType(), *result)
	}
}

func main() {
	fmt.Println("rumgo starting...")

	cfg, cfg_err := config.LoadConfig("./config.json") // TODO: get config path from a env var
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

		time.Sleep(1 * time.Second)
	}
}
