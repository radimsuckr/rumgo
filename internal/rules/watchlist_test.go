package rules

import (
	"testing"

	"github.com/radimsuckr/rumgo/internal/config"
)

const (
	validConfigString   = "{\"version\":\"0.1.0\",\"loopInterval\": 10,\"watchlist\":[{\"url\":\"https://radimsuckr.cz\",\"rules\":[{\"type\":\"contains\",\"value\":\"Sückr\"},{\"type\":\"xpath-contains\",\"path\":\"//p\",\"value\":\"cloud\"},{\"type\":\"contains\",\"invert\":true,\"value\":\"Sückr\"}]},{\"url\":\"http://localhost:8000\",\"rules\":[{\"type\":\"contains\",\"value\":\"Sückr\"}]}]}"
	invalidConfigString = "{\"version\":\"0.1.0\",\"loopInterval\": 10,\"watchlist\":[{\"url\":\"https://radimsuckr.cz\",\"rules\":[{\"type\":\"abc-xyz-123\",\"path\":\"//p\",\"value\":\"cloud\"}]},{\"url\":\"http://localhost:8000\",\"rules\":[{\"type\":\"contains\",\"value\":\"Sückr\"}]}]}"
)

func TestNewWatchlistFromCompletelyValidConfigShouldPass(t *testing.T) {
	const expectedItems = 2
	const expectedRules = 3

	config, err := config.NewConfig([]byte(validConfigString))
	if err != nil {
		t.Fatal(err)
	}

	wl, err := NewWatchlistFromConfig(config)
	if err != nil {
		t.Fatal(err)
	}

	if l := len(wl.Items); l != expectedItems {
		t.Fatalf("watchlist should contain %d items, it has %d", expectedItems, l)
	}

	if wl.Items[0].URL != "https://radimsuckr.cz" {
		t.Fatal("url from config does not match watchlist")
	}

	if l := len(wl.Items[0].Rules); l != expectedRules {
		t.Fatalf("first item should have %d rules, it has %d", expectedRules, l)
	}

	if wl.Items[0].Rules[0].GetType() != ruleTypeContains {
		t.Fatalf("first rule shoud be of type %s", ruleTypeContains)
	}

	if wl.Items[0].Rules[1].GetType() != ruleTypeXPathContains {
		t.Fatalf("second rule shoud be of type %s", ruleTypeXPathContains)
	}

	if wl.Items[0].Rules[2].GetType() != getNotType(ruleTypeContains) {
		t.Fatalf("third rule shoud be of type %s", getNotType(ruleTypeContains))
	}
}

func TestNewWatchlistFromConfigWithInvalidstringShouldFail(t *testing.T) {
	config, err := config.NewConfig([]byte(invalidConfigString))
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewWatchlistFromConfig(config)
	if err == nil {
		t.Fatal("creating a watchlist from a config with invalid rule should return an error")
	}

	if err.Error() != invalidstringError {
		t.Fatal("invalid error returned")
	}
}
