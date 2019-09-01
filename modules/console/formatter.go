package console

import (
	"fmt"
	"regexp"
	"strings"
)

// Formatter ...
type Formatter struct {
	colors      map[string]string
	colorPatter *regexp.Regexp
}

// GetFormatterInstance ...
func GetFormatterInstance() *Formatter {
	return &Formatter{
		colors: map[string]string{
			"gray":   "\033[1;37m%s\033[0m",
			"black":  "\033[1;38m%s\033[0m",
			"blue":   "\033[1;34m%s\033[0m",
			"yellow": "\033[1;33m%s\033[0m",
			"red":    "\033[1;31m%s\033[0m",
			"green":  "\033[0;32m%s\033[0m",
		},
		colorPatter: regexp.MustCompile(`<(\w+)>(.*?)<\/>`),
	}
}

// Write ...
func (f *Formatter) Write(text string) {
	fmt.Print(f.colorize(text))
}

// Writeln ...
func (f *Formatter) Writeln(text string) {
	fmt.Println(f.colorize(text))
}

// NewLine ...
func (f *Formatter) NewLine() {
	fmt.Println()
}

// Text ...
func (f *Formatter) Text(text string) {
	fmt.Println("  " + f.colorize(text))
}

func (f *Formatter) colorize(text string) string {

	matches := f.colorPatter.FindAllStringSubmatch(text, -1)
	for _, group := range matches {
		if pattern, ok := f.colors[group[1]]; ok {
			text = strings.Replace(text, group[0], fmt.Sprintf(pattern, group[2]), -1)
		}
	}

	return text
}
