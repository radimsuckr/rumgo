package rules

import (
	"errors"

	"rumgo/internal/config"
)

// WatchlistItem is a single item to watch with associated rules
type WatchlistItem struct {
	URL   string
	Rules []rule
}

// Watchlist holds all items to watch
type Watchlist struct {
	Items []WatchlistItem
}

const invalidstringError = "invalid rule type"

func createRule(configRule config.Rule) (rule, error) {
	var output rule
	switch configRule.Type {
	case ruleTypeContains:
		output = newContainsRule(configRule.Value)
	case ruleTypeXPathContains:
		output = newXPathContainsRule(configRule.Path, configRule.Value)
	default:
		return nil, errors.New(invalidstringError)
	}
	if configRule.Invert {
		output = newNot(output)
	}
	return output, nil
}

// NewWatchlistFromConfig parses a watchlist from program's configuration
func NewWatchlistFromConfig(config config.Config) (Watchlist, error) {
	wl := Watchlist{}

	for _, item := range config.Watchlist {
		wli := WatchlistItem{URL: item.URL}

		for _, rule := range item.Rules {
			wlRule, err := createRule(rule)
			if err != nil {
				return wl, err
			}
			wli.Rules = append(wli.Rules, wlRule)
		}

		wl.Items = append(wl.Items, wli)
	}

	return wl, nil
}
