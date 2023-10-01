package main

import (
	"log"
	"net/http"

	"radimsuckr/rumgo/client"
	"radimsuckr/rumgo/config"
	"radimsuckr/rumgo/rules"
)

func sendTelegramMessage(telegram config.Telegram, text string) {
	resp, err := http.Get("https://api.telegram.org/bot" + telegram.Token + "/sendMessage?chat_id=" + telegram.Channel + "&text=" + text)
	if err != nil {
		log.Printf("Failed sending Telegram message: %s\n", err)
	}
	if resp.StatusCode != 200 {
		log.Printf("Failed sending Telegram message, status code = %d\n", resp.StatusCode)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
}

func Scrape(telegram config.Telegram, item rules.WatchlistItem) {
	resp, err := client.SendRequest(item.URL)
	if err != nil {
		log.Printf("Failed sending request to: %s\n", item.URL)
	} else {
		for _, rule := range item.Rules {
			result, rule_err := rule.Evaluate(&resp.Content)
			if rule_err != nil {
				log.Printf("Rule error: %s\n", rule_err)
				continue
			}

			if result {
				text := "Rule " + string(rule.GetType()) + " matched for " + item.URL
				sendTelegramMessage(telegram, text)
			}

			log.Printf("%s (%s): %t\n", item.URL, rule.GetType(), result)
		}
	}
}
