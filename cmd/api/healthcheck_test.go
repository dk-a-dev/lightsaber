package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"lightsaber.dkadev.xyz/internal/jsonlog"
)

func TestHealthCheckHandler(t *testing.T) {
	app := &application{
		logger: jsonlog.New(nil, jsonlog.LevelOff),
		config: config{
			env: "test",
		},
	}

	req := httptest.NewRequest("GET", "/v1/healthcheck", nil)
	w := httptest.NewRecorder()

	app.healthCheckHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
	}

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}

	if response["status"] != "available" {
		t.Errorf("expected status 'available', got %v", response["status"])
	}

	systemInfo, ok := response["system_info"].(map[string]interface{})
	if !ok {
		t.Errorf("expected system_info to be a map")
		return
	}

	if systemInfo["environment"] != "test" {
		t.Errorf("expected environment 'test', got %v", systemInfo["environment"])
	}
}
