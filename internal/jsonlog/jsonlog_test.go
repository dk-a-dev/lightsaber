package jsonlog

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf, LevelInfo)

	if logger == nil {
		t.Error("expected logger to be created")
	}
}

func TestLogLevels(t *testing.T) {
	tests := []struct {
		name     string
		level    Level
		expected string
	}{
		{"Info level", LevelInfo, "INFO"},
		{"Error level", LevelError, "ERROR"},
		{"Fatal level", LevelFatal, "FATAL"},
		{"Off level", LevelOff, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.level.String() != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, tt.level.String())
			}
		})
	}
}

func TestPrintInfo(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf, LevelInfo)

	logger.PrintInfo("test message", map[string]string{"key": "value"})

	var logEntry map[string]interface{}
	err := json.Unmarshal(buf.Bytes(), &logEntry)
	if err != nil {
		t.Fatalf("failed to unmarshal log entry: %v", err)
	}

	if logEntry["level"] != "INFO" {
		t.Errorf("expected level INFO, got %v", logEntry["level"])
	}

	if logEntry["message"] != "test message" {
		t.Errorf("expected message 'test message', got %v", logEntry["message"])
	}
}

func TestPrintError(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf, LevelError)

	err := bytes.ErrTooLarge
	logger.PrintError(err, map[string]string{"context": "test"})

	var logEntry map[string]interface{}
	unmarshalErr := json.Unmarshal(buf.Bytes(), &logEntry)
	if unmarshalErr != nil {
		t.Fatalf("failed to unmarshal log entry: %v", unmarshalErr)
	}

	if logEntry["level"] != "ERROR" {
		t.Errorf("expected level ERROR, got %v", logEntry["level"])
	}

	if logEntry["message"] != err.Error() {
		t.Errorf("expected message '%s', got %v", err.Error(), logEntry["message"])
	}
}

func TestLogFiltering(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf, LevelError)

	// Info message should not be logged when level is Error
	logger.PrintInfo("this should not appear", nil)

	if buf.Len() > 0 {
		t.Error("expected no output when logging below threshold level")
	}

	// Error message should be logged
	logger.PrintError(bytes.ErrTooLarge, nil)

	if buf.Len() == 0 {
		t.Error("expected output when logging at or above threshold level")
	}
}
