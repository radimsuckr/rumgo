package rules

import (
	"testing"

	"radimsuckr/rumgo/config"
)

const (
	validConfigString   = "{\"version\":\"0.1.0\",\"loopInterval\": 10,\"watchlist\":[{\"url\":\"https://radimsuckr.cz\",\"rules\":[{\"type\":\"contains\",\"value\":\"Sückr\"},{\"type\":\"xpath-contains\",\"path\":\"//p\",\"value\":\"cloud\"}]},{\"url\":\"http://localhost:8000\",\"rules\":[{\"type\":\"contains\",\"value\":\"Sückr\"}]}]}"
	invalidConfigString = "{\"version\":\"0.1.0\",\"loopInterval\": 10,\"watchlist\":[{\"url\":\"https://radimsuckr.cz\",\"rules\":[{\"type\":\"abc-xyz-123\",\"path\":\"//p\",\"value\":\"cloud\"}]},{\"url\":\"http://localhost:8000\",\"rules\":[{\"type\":\"contains\",\"value\":\"Sückr\"}]}]}"
)

func TestNewWatchlistFromCompletelyValidConfigShouldPass(t *testing.T) {
	config, err := config.NewConfig([]byte(validConfigString))
	if err != nil {
		t.Fatal(err)
	}

	wl, err := NewWatchlistFromConfig(config)
	if err != nil {
		t.Fatal(err)
	}

	if len(wl.Items) != 2 {
		t.Fatal("watchlist should contain 2 items")
	}

	if wl.Items[0].URL != "https://radimsuckr.cz" {
		t.Fatal("url from config does not match watchlist")
	}

	if wl.Items[0].Rules[0].GetType() != RuleTypeContains {
		t.Fatalf("first rule shoud be of type %s", RuleTypeContains)
	}

	if wl.Items[0].Rules[1].GetType() != RuleTypeXPathContains {
		t.Fatalf("first rule shoud be of type %s", RuleTypeXPathContains)
	}
}

func TestNewWatchlistFromConfigWithInvalidRuleTypeShouldFail(t *testing.T) {
	config, err := config.NewConfig([]byte(invalidConfigString))
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewWatchlistFromConfig(config)
	if err == nil {
		t.Fatal("creating a watchlist from a config with invalid rule should return an error")
	}

	if err.Error() != InvalidRuleTypeError {
		t.Fatal("invalid error returned")
	}
}
