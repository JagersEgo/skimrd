package processor

import (
	"unicode"
	"unicode/utf8"
)

const MaxFocus = 3
const MinFocus = 0

func Tokenise(word string) Token {
	rl := utf8.RuneCountInString(word)

	t := Token{
		Word:  word,
		Focus: get_focus(word, rl),
		Width: utf8.RuneCountInString(word),
		Pause: get_pause(word),
	}

	return t
}

func get_focus(word string, width int) int {
	r, _ := (utf8.DecodeLastRuneInString(word))
	if unicode.IsPunct(r) {
		width -= 1
	}

	return min((width-1)/2, len(word))
}

func get_pause(word string) PauseLength {
	r, _ := utf8.DecodeLastRuneInString(word)

	switch r {
	case '.':
		return Long
	case ',':
		return Medium
	case '"':
		return Medium
	default:
		return Normal
	}
}
