package config

import (
	"os"
	"testing"
)

func TestNewConfigWithInvalidJSONShoudFail(t *testing.T) {
	configStr := "invalid json"

	_, err := NewConfig([]byte(configStr))

	if err == nil {
		t.Fatal("reading config with invalid JSON should fail")
	}
}

func TestNewConfigWithValidJSONAndVersionShouldPass(t *testing.T) {
	configStr := "{\"version\":\"0.1.0\"}"

	_, err := NewConfig([]byte(configStr))

	if err != nil {
		t.Fatalf("reading config with valid JSON should pass, %s", err)
	}
}

func TestNewConfigWithValidJSONButInvalidVersionShouldFail(t *testing.T) {
	configStr := "{\"version\":\"0.0.0\"}"

	_, err := NewConfig([]byte(configStr))

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

	expectedStr := "hello world"
	if _, err := file.WriteString(expectedStr); err != nil {
		t.Fatal("failed to prepare test file")
	}
	file.Close()

	path := file.Name()

	contentBytes, err := ReadConfigFile(path)
	if err != nil {
		t.Fatalf("reading a file from existing path %s should pass", path)
	}

	contentStr := string(contentBytes)
	if contentStr != expectedStr {
		t.Fatalf("content should be \"%s\" but is \"%s\"", expectedStr, contentStr)
	}
}
