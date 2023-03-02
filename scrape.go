package main

import (
	"fmt"

	"radimsuckr/rumgo/client"
	"radimsuckr/rumgo/rules"
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
