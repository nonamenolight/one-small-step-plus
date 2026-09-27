package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConfigHandler(t *testing.T) {
	config := Config{
		SiteName:  "Test Site",
		StartDate: "2026-01-01",
	}

	req := httptest.NewRequest(http.MethodGet, "/api/config", nil)
	rec := httptest.NewRecorder()

	handler := configHandler(config)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var got Config

	err := json.NewDecoder(rec.Body).Decode(&got)
	if err != nil {
		t.Fatal(err)
	}

	if got.SiteName != config.SiteName {
		t.Errorf("expected site name %q, got %q", config.SiteName, got.SiteName)
	}

	if got.StartDate != config.StartDate {
		t.Errorf("expected start date %q, got %q", config.StartDate, got.StartDate)
	}
}
