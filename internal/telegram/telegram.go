// Package telegram handles the Telegram bot functions
package telegram

import (
	"log"
	"net/http"

	"github.com/radimsuckr/rumgo/internal/config"
)

// SendTelegramMessage sends a simple message via Telegram's sendMessage API
func SendTelegramMessage(telegram config.Telegram, text string) {
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
