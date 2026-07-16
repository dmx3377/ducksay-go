package main

import (
	"strings"
	"testing"
)


// gotta make sure it works, innit?
func TestTextWidth(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello", 5},
		{"", 0},
		{"I am waddles", 1}, // Emoji test (UTF-8 multi-byte character, but 1 rune wide)
		{"hello waddles", 7},
	}

	for _, tt := range tests {
		actual := textWidth(tt.input)
		if actual != tt.expected {
			t.Errorf("textWidth(%q) = %d; want %d", tt.input, actual, tt.expected)
		}
	}
}

func TestWrapMessage(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		width    int
		expected []string
	}{
		{
			name:     "simple wrap",
			message:  "hello world",
			width:    5,
			expected: []string{"hello", "world"},
		},
		{
			name:     "fits perfectly",
			message:  "hello",
			width:    10,
			expected: []string{"hello"},
		},
		{
			name:     "empty message",
			message:  "",
			width:    5,
			expected: []string{""},
		},
		{
			name:     "long word split",
			message:  "supercalifragilistic",
			width:    5,
			expected: []string{"super", "calif", "ragil", "istic"},
		},
		{
			name:     "newlines preserved",
			message:  "line one\nline two",
			width:    20,
			expected: []string{"line one", "line two"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := wrapMessage(tt.message, tt.width)
			if len(actual) != len(tt.expected) {
				t.Fatalf("wrapMessage returned %d lines; want %d. Got: %v", len(actual), len(tt.expected), actual)
			}
			for i := range actual {
				if actual[i] != tt.expected[i] {
					t.Errorf("line %d = %q; want %q", i, actual[i], tt.expected[i])
				}
			}
		})
	}
}

func TestRender(t *testing.T) {
	// Test basic Mono rendering
	result := Render("hello", 10)
	if !strings.Contains(result, "(hello)") {
		t.Errorf("Expected render to contain speech bubble '(hello)', got:\n%s", result)
	}
	if !strings.Contains(result, " .__(.)< ") {
		t.Errorf("Expected render to contain default duck layout, got:\n%s", result)
	}

	// Test Twitter rendering
	resultTwitter := RenderWithStyle("hello", 10, StyleTwitter)
	if !strings.Contains(resultTwitter, " .__( . )< ") {
		t.Errorf("Expected render to contain Twitter duck layout, got:\n%s", resultTwitter)
	}
}
