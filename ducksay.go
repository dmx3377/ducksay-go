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
	Header:          "<!--          _\n",
	SpeechPrefix:    "        .__( . )< ",
	Body:            "         \\___)",
	ContinuationGap: "     ",
}

var MonoLayout = DuckLayout{
	Header:          "<!--           _\n",
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
	styleFlag := flag.String("style", "html", "Comment style (html, go, js, py, sql, none, etc.)")
	colorDuckFlag := flag.String("color-duck", "", "Color of the duck (red, green, yellow, blue, magenta, cyan, white, devgreen, or hex code)")
	colorBubbleFlag := flag.String("color-bubble", "", "Color of the speech bubble (red, green, yellow, blue, magenta, cyan, white, devgreen, or hex code)")

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

	fmt.Print(RenderCustom(message, width, style, *styleFlag, *colorDuckFlag, *colorBubbleFlag))
}

// Render with default Mono style
func Render(message string, width int) string {
	return RenderWithStyle(message, width, StyleMono)
}

// Render with custom layout style
func RenderWithStyle(message string, width int, style Style) string {
	return RenderCustom(message, width, style, "html", "", "")
}

// RenderCustom renders the duck with custom style, comment wrappers, and colors
func RenderCustom(message string, width int, duckStyle Style, commentStyle string, duckColorName string, bubbleColorName string) string {
	lines := wrapMessage(message, width)
	prefix, suffix, isBlock := getCommentConfig(commentStyle)
	duckColor := getColorCode(duckColorName)
	bubbleColor := getColorCode(bubbleColorName)
	return renderDuckCustom(lines, duckStyle, prefix, suffix, isBlock, duckColor, bubbleColor)
}

func renderDuckCustom(lines []string, duckStyle Style, prefix, suffix string, isBlock bool, duckColor, bubbleColor string) string {
	var output strings.Builder

	eye := "."
	if duckStyle == StyleTwitter {
		eye = " . "
	}

	duckHeadStr := ".__(" + eye + ")<"
	duckBodyStr := "\\___)"

	// Header
	pLen := utf8.RuneCountInString(prefix)
	var headerLen int
	if duckStyle == StyleTwitter {
		headerLen = 13 - pLen
	} else {
		headerLen = 12 - pLen
	}
	if headerLen < 0 {
		headerLen = 0
	}
	headerSpaces := strings.Repeat(" ", headerLen)

	output.WriteString(colorText(prefix, bubbleColor))
	output.WriteString(headerSpaces)
	output.WriteString(colorText("_", duckColor))
	output.WriteByte('\n')

	// Line-by-line helper values
	var headIndent string
	var bodyIndent string
	var contGap string

	if isBlock {
		headIndent = "        "  // 8 spaces
		bodyIndent = "         " // 9 spaces
	} else {
		pLen := utf8.RuneCountInString(prefix)
		headLen := 8 - pLen
		if headLen < 0 {
			headLen = 0
		}
		headIndent = strings.Repeat(" ", headLen)

		bodyLen := 9 - pLen
		if bodyLen < 0 {
			bodyLen = 0
		}
		bodyIndent = strings.Repeat(" ", bodyLen)
	}

	if duckStyle == StyleTwitter {
		contGap = "     " // 5 spaces
	} else {
		contGap = "    " // 4 spaces
	}

	// Write bubble and duck
	firstLine := ""
	if len(lines) > 0 {
		firstLine = lines[0]
	}

	if isBlock {
		output.WriteString(headIndent)
	} else {
		output.WriteString(colorText(prefix, bubbleColor))
		output.WriteString(headIndent)
	}
	output.WriteString(colorText(duckHeadStr, duckColor))
	output.WriteString(" ")
	output.WriteString(colorText("(", bubbleColor))
	output.WriteString(colorText(firstLine, bubbleColor))

	if len(lines) > 1 {
		output.WriteByte('\n')
		if isBlock {
			output.WriteString(bodyIndent)
		} else {
			output.WriteString(colorText(prefix, bubbleColor))
			output.WriteString(bodyIndent)
		}
		output.WriteString(colorText(duckBodyStr, duckColor))
		output.WriteString(contGap)

		var rest strings.Builder
		for i, line := range lines[1:] {
			if i > 0 {
				rest.WriteByte(' ')
			}
			rest.WriteString(line)
		}
		output.WriteString(colorText(rest.String(), bubbleColor))
		output.WriteString(colorText(")\n", bubbleColor))
	} else {
		output.WriteString(colorText(")\n", bubbleColor))
		if isBlock {
			output.WriteString(bodyIndent)
		} else {
			output.WriteString(colorText(prefix, bubbleColor))
			output.WriteString(bodyIndent)
		}
		output.WriteString(colorText(duckBodyStr, duckColor))
		output.WriteByte('\n')
	}

	// Footer
	if isBlock {
		output.WriteString(colorText(" ~~~~~~~~~~~~~~~~~~", bubbleColor))
		output.WriteString(colorText(suffix, bubbleColor))
		output.WriteByte('\n')
	} else {
		output.WriteString(colorText(prefix, bubbleColor))
		output.WriteString(colorText(" ~~~~~~~~~~~~~~~~~~\n", bubbleColor))
	}

	return output.String()
}

func getCommentConfig(styleName string) (string, string, bool) {
	switch strings.ToLower(strings.TrimSpace(styleName)) {
	case "html", "xml", "default":
		return "<!--", "-->", true
	case "go", "js", "ts", "cpp", "c", "java", "rust", "swift", "cs", "kotlin", "scala":
		return "//", "", false
	case "py", "python", "sh", "bash", "rb", "ruby", "pl", "perl", "yaml", "yml", "make", "docker", "toml":
		return "#", "", false
	case "sql", "lua", "hs", "haskell", "ada":
		return "--", "", false
	case "none", "raw":
		return "", "", false
	default:
		return "<!--", "-->", true
	}
}

func getColorCode(colorName string) string {
	colorName = strings.ToLower(strings.TrimSpace(colorName))
	if colorName == "" {
		return ""
	}
	switch colorName {
	case "red":
		return "\x1b[31m"
	case "green":
		return "\x1b[32m"
	case "yellow":
		return "\x1b[33m"
	case "blue":
		return "\x1b[34m"
	case "magenta":
		return "\x1b[35m"
	case "cyan":
		return "\x1b[36m"
	case "white":
		return "\x1b[37m"
	case "devgreen", "dev-green", "dev green":
		return "\x1b[38;2;0;229;130m"
	}
	if strings.HasPrefix(colorName, "#") {
		if code, ok := parseHexColor(colorName); ok {
			return code
		}
	}
	return ""
}

func parseHexColor(hex string) (string, bool) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) == 3 {
		var r, g, b byte
		_, err := fmt.Sscanf(hex, "%1x%1x%1x", &r, &g, &b)
		if err == nil {
			return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r*17, g*17, b*17), true
		}
	} else if len(hex) == 6 {
		var r, g, b byte
		_, err := fmt.Sscanf(hex, "%2x%2x%2x", &r, &g, &b)
		if err == nil {
			return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b), true
		}
	}
	return "", false
}

func colorText(text string, colorCode string) string {
	if colorCode == "" || text == "" {
		return text
	}
	return colorCode + text + "\x1b[0m"
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
