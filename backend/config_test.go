package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	dir := t.TempDir()

	configPath := filepath.Join(dir, "config.yaml")

	content := []byte(`
site_name: "Test Site"
start_date: "2026-01-01"
`)

	err := os.WriteFile(configPath, content, 0644)
	if err != nil {
		t.Fatal(err)
	}

	config, err := loadConfig(configPath)
	if err != nil {
		t.Fatal(err)
	}

	if config.SiteName != "Test Site" {
		t.Errorf("expected site name %q, got %q", "Test Site", config.SiteName)
	}

	if config.StartDate != "2026-01-01" {
		t.Errorf("expected start date %q, got %q", "2026-01-01", config.StartDate)
	}
}
