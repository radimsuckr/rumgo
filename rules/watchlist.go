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

const InvalidRuleTypeError = "invalid rule type"

func NewWatchlistFromConfig(config config.Config) (Watchlist, error) {
	wl := Watchlist{}

	for _, item := range config.Watchlist {
		wli := WatchlistItem{URL: item.URL}

		for _, rule := range item.Rules {
			switch rule.Type {
			case RuleTypeContains:
				wli.Rules = append(wli.Rules, NewContainsRule(rule.Value))
			case RuleTypeXPathContains:
				wli.Rules = append(wli.Rules, NewXPathContainsRule(rule.Path, rule.Value))
			default:
				return wl, errors.New(InvalidRuleTypeError)
			}
		}

		wl.Items = append(wl.Items, wli)
	}

	return wl, nil
}
