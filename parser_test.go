package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseJSON_ValidJSON_Empty_ReturnEmpty(t *testing.T) {
	// Test for valid JSON
	jsonStr := `{}`
	result, err := parseJSON(jsonStr)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{}, result)
}

func TestParseJSON_InvalidJSON_MissingClosingBracket_ThrowError(t *testing.T) {
	// Test for invalid JSON
	jsonStr := `{`
	_, err := parseJSON(jsonStr)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "Invalid JSON. Missing `}` at the end")
}

func TestParseJSON_ValidJSON_WithKeyValue_ReturnJSONWithKeyValue(t *testing.T) {
	// Test for valid JSON with one key-value pair
	jsonStr := `{"key": "value"}`
	result, err := parseJSON(jsonStr)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"key": "value"}, result)

	// Test for valid JSON with multiple key-value pairs
	jsonStr = `{"name": "John", "age": "30"}`
	result, err = parseJSON(jsonStr)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"name": "John", "age": "30"}, result)
}

func TestParseJSON_InvalidJSON_MissingQuotesAroundValue(t *testing.T) {
	// Test for invalid JSON: missing quotes around key
	invalidJSON := `{key: "value"}`
	_, err := parseJSON(invalidJSON)
	assert.Error(t, err)
}

func TestParseJSON_InvalidJSON_MissingColon(t *testing.T) {
	// Test for invalid JSON: missing colon
	invalidJSON := `{"key" "value"}`
	_, err := parseJSON(invalidJSON)
	assert.Error(t, err)
}

func TestParseJSON_InvalidJSON_MissingClosingBrace(t *testing.T) {
	// Test for invalid JSON: missing closing brace
	invalidJSON := `{"key": "value"`
	_, err := parseJSON(invalidJSON)
	assert.Error(t, err)
}
