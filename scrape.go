package main

import (
	"fmt"

	"radimsuckr/rumgo/client"
	"radimsuckr/rumgo/rules"
)

func Scrape(item rules.WatchlistItem) {
	resp, err := client.SendRequest(item.URL)
	if err != nil {
		fmt.Printf("Failed sending request to: %s\n", item.URL)
	} else {
		for _, rule := range item.Rules {
			result, rule_err := rule.Evaluate(&resp.Content)
			if rule_err != nil {
				fmt.Printf("Rule error: %s\n", rule_err)
				continue
			}
			fmt.Printf("%s (%s): %t\n", item.URL, rule.GetType(), result)
		}
	}
}
