package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// Default Configuration
const (
	DefaultWidth   = 40
	DefaultMessage = "I am Waddles"
	Footer         = " ~~~~~~~~~~~~~~~~~~-->\n"
)

// Style types
type Style int

const (
	StyleMono Style = iota
	StyleTwitter
)

type DuckLayout struct {
	Header          string
	SpeechPrefix    string
	Body            string
	ContinuationGap string
}

var TwitterLayout = DuckLayout{
	Header:          "<!--       _\n",
	SpeechPrefix:    "        .__( . )< ",
	Body:            "         \\___)",
	ContinuationGap: "     ",
}

var MonoLayout = DuckLayout{
	Header:          "<!--        _\n",
	SpeechPrefix:    "        .__(.)< ",
	Body:            "         \\___)",
	ContinuationGap: "    ",
}

func (s Style) Layout() DuckLayout {
	switch s {
	case StyleTwitter:
		return TwitterLayout
	case StyleMono:
		fallthrough
	default:
		return MonoLayout
	}
}

func main() {
	// Replicating clap CLI flags
	widthFlag := flag.Int("width", DefaultWidth, "Wrap speech text at N columns")
	flag.IntVar(widthFlag, "w", DefaultWidth, "Wrap speech text at N columns (shorthand)")
	twitterFlag := flag.Bool("twitter", false, "Use Twitter-compatible output")
	napFlag := flag.Bool("nap", false, "Waddles takes a nap")
	sleepFlag := flag.Bool("sleep", false, "Waddles takes a nap (alias for -nap)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "ducksay - Make Waddles say things.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n  ducksay [flags] [MESSAGE...]\n\nFlags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *napFlag || *sleepFlag {
		fmt.Print("# Waddles takin a nap\n    _\n.__(_)<\n \\___)\n")
		return
	}

	// Handle trailing positional arguments as the message
	message := strings.Join(flag.Args(), " ")
	if strings.TrimSpace(message) == "" {
		message = DefaultMessage
	}

	// Enforce width boundary (> 0)
	width := *widthFlag
	if width < 1 {
		width = 1
	}

	style := StyleMono
	if *twitterFlag {
		style = StyleTwitter
	}

	fmt.Print(RenderWithStyle(message, width, style))
}

// Render with default Mono style
func Render(message string, width int) string {
	return RenderWithStyle(message, width, StyleMono)
}

// Render with custom layout style
func RenderWithStyle(message string, width int, style Style) string {
	lines := wrapMessage(message, width)
	return renderDuck(lines, style.Layout())
}

func renderDuck(lines []string, layout DuckLayout) string {
	var output strings.Builder

	output.WriteString(layout.Header)

	switch len(lines) {
	case 1:
		writeSingleLineMessage(&output, layout, lines[0])
	default:
		if len(lines) > 1 {
			writeWrappedMessage(&output, layout, lines[0], lines[1:])
		} else {
			writeSingleLineMessage(&output, layout, "")
		}
	}

	output.WriteString(Footer)
	return output.String()
}

func writeSingleLineMessage(output *strings.Builder, layout DuckLayout, line string) {
	output.WriteString(layout.SpeechPrefix)
	output.WriteByte('(')
	output.WriteString(line)
	output.WriteString(")\n")
	output.WriteString(layout.Body)
	output.WriteByte('\n')
}

func writeWrappedMessage(output *strings.Builder, layout DuckLayout, first string, rest []string) {
	output.WriteString(layout.SpeechPrefix)
	output.WriteByte('(')
	output.WriteString(first)
	output.WriteByte('\n')
	output.WriteString(layout.Body)
	output.WriteString(layout.ContinuationGap)
	writeContinuationLines(output, rest)
	output.WriteString(")\n")
}

func writeContinuationLines(output *strings.Builder, lines []string) {
	for i, line := range lines {
		if i > 0 {
			output.WriteByte(' ')
		}
		output.WriteString(line)
	}
}

func wrapMessage(message string, width int) []string {
	if width < 1 {
		width = 1
	}
	var wrappedLines []string

	// Split by newlines in source message
	sourceLines := strings.Split(message, "\n")
	for _, sourceLine := range sourceLines {
		wrapSourceLine(sourceLine, width, &wrappedLines)
	}

	if len(wrappedLines) == 0 {
		wrappedLines = append(wrappedLines, "")
	}

	return wrappedLines
}

func wrapSourceLine(sourceLine string, width int, wrappedLines *[]string) {
	if sourceLine == "" {
		*wrappedLines = append(*wrappedLines, "")
		return
	}

	var currentLine strings.Builder
	words := strings.Fields(sourceLine)

	for _, word := range words {
		if textWidth(word) > width {
			finishLine(&currentLine, wrappedLines)
			*wrappedLines = append(*wrappedLines, splitLongWord(word, width)...)
			continue
		}

		if !wordFitsOnLine(currentLine.String(), word, width) {
			finishLine(&currentLine, wrappedLines)
		}

		appendWord(&currentLine, word)
	}

	finishLine(&currentLine, wrappedLines)
}

func wordFitsOnLine(line string, word string, width int) bool {
	return line == "" || textWidth(line)+1+textWidth(word) <= width
}

func appendWord(line *strings.Builder, word string) {
	if line.Len() > 0 {
		line.WriteByte(' ')
	}
	line.WriteString(word)
}

func finishLine(line *strings.Builder, wrappedLines *[]string) {
	if line.Len() > 0 {
		*wrappedLines = append(*wrappedLines, line.String())
		line.Reset()
	}
}

func splitLongWord(word string, width int) []string {
	var chunks []string
	var current strings.Builder

	for _, ch := range word {
		current.WriteRune(ch)
		if textWidth(current.String()) >= width {
			chunks = append(chunks, current.String())
			current.Reset()
		}
	}

	if current.Len() > 0 {
		chunks = append(chunks, current.String())
	}

	return chunks
}

func textWidth(value string) int {
	return utf8.RuneCountInString(value)
}
