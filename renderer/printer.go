package renderer

import (
	"fmt"
	"jager/fast_reader/processor"
	"log"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/x/term"
)

func PrintTokenStream(tokens chan processor.Token, base_wpm int64) {
	baseDelay := time.Minute / time.Duration(base_wpm)

	Initialise()
	defer CleanUp()

	for token := range tokens {
		PrintToken(token)

		pauseTime := float64(baseDelay) * processor.PauseLenMult[token.Pause]
		time.Sleep(time.Duration(pauseTime))
	}
}

func PrintToken(t processor.Token) {
	focused := applyFocus(t.Word, t.Focus)

	padding, e := get_padding(t.Focus, t.Width)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Print(CarriageReturn)

	fmt.Print(padding, focused)
}

func Initialise() {
	// Hide cursor
	fmt.Print(HideCursor)
}

func CleanUp() {
	// Show cursor
	fmt.Print(ShowCursor)

	fmt.Println(Reset)
}

func applyFocus(word string, focus int) string {
	runes := []rune(word)

	var b strings.Builder
	b.Grow(len(word) + 16) // small optimization for ANSI codes

	for i, r := range runes {
		if i == focus {
			b.WriteString(Bold)
			b.WriteString(Red)
			b.WriteRune(r)
			b.WriteString(Reset)
		} else {
			b.WriteRune(r)
		}
	}

	return b.String()
}

func get_padding(focus int, width int) (string, error) {
	termWidth, _, err := term.GetSize(uintptr(os.Stdout.Fd()))
	if err != nil {
		return "", err
	}

	leftPadding := (termWidth / 2) - focus
	return strings.Repeat(" ", leftPadding), nil
}
