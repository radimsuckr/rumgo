package rules

import (
	"errors"

	"radimsuckr/rumgo/config"
)

type WatchlistItem struct {
	URL   string
	Rules []Rule
}

type Watchlist struct {
	Items []WatchlistItem
}

func NewWatchlistFromConfig(config *config.Config) (*Watchlist, error) {
	wl := &Watchlist{}

	for _, item := range config.Watchlist {
		wli := WatchlistItem{URL: item.URL}

		for _, rule := range item.Rules {
			switch rule.Type {
			case "contains":
				wli.Rules = append(wli.Rules, &ContainsRule{Value: rule.Value})
			case "xpath-contains":
				wli.Rules = append(wli.Rules, &XPathContains{Value: rule.Value})
			default:
				return nil, errors.New("Invalid rule type")
			}
		}

		wl.Items = append(wl.Items, wli)
	}

	return wl, nil
}
