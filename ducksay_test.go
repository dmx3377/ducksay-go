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
		{"I am waddles",12}, 
		{"hello waddles", 13},
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

func TestRenderCustom(t *testing.T) {
	// Test Go comments line styling
	resultGo := RenderCustom("hello", 10, StyleMono, "go", "", "")
	if !strings.Contains(resultGo, "//          _") {
		t.Errorf("Expected RenderCustom with go style header to be aligned, got:\n%s", resultGo)
	}
	if !strings.Contains(resultGo, "//      .__(.)< (hello)") {
		t.Errorf("Expected RenderCustom with go style to prefix with //, got:\n%s", resultGo)
	}
	if !strings.Contains(resultGo, "// ~~~~~~~~~~~~~~~~~~") {
		t.Errorf("Expected RenderCustom with go style footer to be // ~~~~~~~~~~~~~~~~~~, got:\n%s", resultGo)
	}

	// Test Python comments styling
	resultPy := RenderCustom("hello", 10, StyleMono, "py", "", "")
	if !strings.Contains(resultPy, "#           _") {
		t.Errorf("Expected RenderCustom with py style header to be aligned, got:\n%s", resultPy)
	}
	if !strings.Contains(resultPy, "#       .__(.)< (hello)") {
		t.Errorf("Expected RenderCustom with py style to prefix with #, got:\n%s", resultPy)
	}

	// Test Raw/None styling
	resultNone := RenderCustom("hello", 10, StyleMono, "none", "", "")
	if !strings.Contains(resultNone, "            _") {
		t.Errorf("Expected RenderCustom with none style header to be aligned, got:\n%s", resultNone)
	}
	if !strings.Contains(resultNone, "        .__(.)< (hello)") {
		t.Errorf("Expected RenderCustom with none style to have no comment prefix, got:\n%s", resultNone)
	}

	// Test Color code rendering
	resultColor := RenderCustom("hello", 10, StyleMono, "go", "red", "blue")
	if !strings.Contains(resultColor, "\x1b[31m") {
		t.Errorf("Expected RenderCustom output to contain red ANSI code, got:\n%s", resultColor)
	}
	if !strings.Contains(resultColor, "\x1b[34m") {
		t.Errorf("Expected RenderCustom output to contain blue ANSI code, got:\n%s", resultColor)
	}

	// Test DevGreen rendering
	resultDevGreen := RenderCustom("hello", 10, StyleMono, "go", "devgreen", "")
	if !strings.Contains(resultDevGreen, "\x1b[38;2;0;229;130m") {
		t.Errorf("Expected RenderCustom output to contain devgreen truecolor ANSI code, got:\n%s", resultDevGreen)
	}

	// Test Custom Hex rendering
	resultHex := RenderCustom("hello", 10, StyleMono, "go", "#aabbcc", "")
	if !strings.Contains(resultHex, "\x1b[38;2;170;187;204m") {
		t.Errorf("Expected RenderCustom output to contain parsed hex truecolor ANSI code, got:\n%s", resultHex)
	}
}

func TestParseHexColor(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		ok       bool
	}{
		{"#00E582", "\x1b[38;2;0;229;130m", true},
		{"#fff", "\x1b[38;2;255;255;255m", true},
		{"invalid", "", false},
	}

	for _, tt := range tests {
		actual, ok := parseHexColor(tt.input)
		if ok != tt.ok || actual != tt.expected {
			t.Errorf("parseHexColor(%q) = (%q, %t); want (%q, %t)", tt.input, actual, ok, tt.expected, tt.ok)
		}
	}
}
