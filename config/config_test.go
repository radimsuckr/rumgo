package config

import (
	"os"
	"testing"
)

func TestNewConfigWithInvalidJSONShoudFail(t *testing.T) {
	config_str := "invalid json"

	_, err := NewConfig([]byte(config_str))

	if err == nil {
		t.Fatal("reading config with invalid JSON should fail")
	}
}

func TestNewConfigWithValidJSONAndVersionShouldPass(t *testing.T) {
	config_str := "{\"version\":\"0.1.0\"}"

	_, err := NewConfig([]byte(config_str))

	if err != nil {
		t.Fatalf("reading config with valid JSON should pass, %s", err)
	}
}

func TestNewConfigWithValidJSONButInvalidVersionShouldFail(t *testing.T) {
	config_str := "{\"version\":\"0.0.0\"}"

	_, err := NewConfig([]byte(config_str))

	if err == nil {
		t.Fatalf("reading config with valid JSON but invalid version should fail, %s", err)
	}
}

func TestReadConfigFileWithInvalidPathShouldFail(t *testing.T) {
	path := "/a/b/c/lol"

	_, err := ReadConfigFile(path)

	if err == nil {
		t.Fatalf("reading a file from non-existent path %s should fail", path)
	}
}

func TestReadConfigFileWithValidPathShouldPass(t *testing.T) {
	file, err := os.CreateTemp("", "*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	expected_str := "hello world"
	file.WriteString(expected_str)
	file.Close()

	path := file.Name()

	content_bytes, err := ReadConfigFile(path)
	if err != nil {
		t.Fatalf("reading a file from existing path %s should pass", path)
	}

	content_str := string(content_bytes)
	if content_str != expected_str {
		t.Fatalf("content should be \"%s\" but is \"%s\"", expected_str, content_str)
	}
}
