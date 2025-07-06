package validator

import (
	"regexp"
	"testing"
)

func TestNew(t *testing.T) {
	v := New()
	if v.Errors == nil {
		t.Error("expected Errors to be initialized")
	}
	if len(v.Errors) != 0 {
		t.Error("expected Errors to be empty initially")
	}
}

func TestValid(t *testing.T) {
	v := New()
	if !v.Valid() {
		t.Error("expected new validator to be valid")
	}

	v.AddError("test", "test error")
	if v.Valid() {
		t.Error("expected validator with errors to be invalid")
	}
}

func TestAddError(t *testing.T) {
	v := New()
	v.AddError("field1", "error message 1")
	v.AddError("field2", "error message 2")

	if len(v.Errors) != 2 {
		t.Errorf("expected 2 errors, got %d", len(v.Errors))
	}

	if v.Errors["field1"] != "error message 1" {
		t.Errorf("expected 'error message 1', got '%s'", v.Errors["field1"])
	}

	if v.Errors["field2"] != "error message 2" {
		t.Errorf("expected 'error message 2', got '%s'", v.Errors["field2"])
	}
}

func TestCheck(t *testing.T) {
	v := New()

	// Test with passing condition
	v.Check(true, "field1", "should not add error")
	if len(v.Errors) != 0 {
		t.Error("expected no errors when condition is true")
	}

	// Test with failing condition
	v.Check(false, "field2", "should add error")
	if len(v.Errors) != 1 {
		t.Errorf("expected 1 error, got %d", len(v.Errors))
	}

	if v.Errors["field2"] != "should add error" {
		t.Errorf("expected 'should add error', got '%s'", v.Errors["field2"])
	}
}

func TestIn(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		list     []string
		expected bool
	}{
		{
			name:     "value in list",
			value:    "apple",
			list:     []string{"apple", "banana", "cherry"},
			expected: true,
		},
		{
			name:     "value not in list",
			value:    "grape",
			list:     []string{"apple", "banana", "cherry"},
			expected: false,
		},
		{
			name:     "empty list",
			value:    "apple",
			list:     []string{},
			expected: false,
		},
		{
			name:     "empty value",
			value:    "",
			list:     []string{"apple", "banana", ""},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := In(tt.value, tt.list...)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestMatches(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		rx       *regexp.Regexp
		expected bool
	}{
		{
			name:     "valid email",
			value:    "test@example.com",
			rx:       EmailRX,
			expected: true,
		},
		{
			name:     "invalid email",
			value:    "invalid-email",
			rx:       EmailRX,
			expected: false,
		},
		{
			name:     "empty value",
			value:    "",
			rx:       EmailRX,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Matches(tt.value, tt.rx)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name     string
		values   []string
		expected bool
	}{
		{
			name:     "all unique values",
			values:   []string{"apple", "banana", "cherry"},
			expected: true,
		},
		{
			name:     "duplicate values",
			values:   []string{"apple", "banana", "apple"},
			expected: false,
		},
		{
			name:     "empty slice",
			values:   []string{},
			expected: true,
		},
		{
			name:     "single value",
			values:   []string{"apple"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Unique(tt.values)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
