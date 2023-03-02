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
			case CONTAINS_RULE_TYPE:
				wli.Rules = append(wli.Rules, &ContainsRule{Value: rule.Value})
			case XPATH_CONTAINS_RULE_TYPE:
				wli.Rules = append(wli.Rules, &XPathContains{Path: rule.Path, Value: rule.Value})
			default:
				return nil, errors.New("Invalid rule type")
			}
		}

		wl.Items = append(wl.Items, wli)
	}

	return wl, nil
}
