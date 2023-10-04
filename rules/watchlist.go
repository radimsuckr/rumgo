package rules

import (
	"errors"

	"rumgo/config"
)

type WatchlistItem struct {
	URL   string
	Rules []Rule
}

type Watchlist struct {
	Items []WatchlistItem
}

const InvalidRuleTypeError = "invalid rule type"

func createRule(rule config.Rule) (Rule, error) {
	var output Rule
	switch rule.Type {
	case RuleTypeContains:
		output = NewContainsRule(rule.Value)
	case RuleTypeXPathContains:
		output = NewXPathContainsRule(rule.Path, rule.Value)
	default:
		return nil, errors.New(InvalidRuleTypeError)
	}
	if rule.Invert {
		output = NewNot(output)
	}
	return output, nil
}

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
