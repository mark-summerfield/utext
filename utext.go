// Copyright © 2024 Mark Summerfield. All rights reserved.

// ([TOC]) This package provides some text-related functions.
//
// [TOC]: file:///home/mark/app/golib/doc/index.html
package utext

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

//go:embed Version.dat
var Version string

// Centered returns s centered between runs of the given pad to make it
// width long, or returns s as-is if it is >= width runes long.
func Centered(s string, pad rune, width int) string {
	size := utf8.RuneCountInString(s)
	if size >= width {
		return s
	}
	remainder := width - size
	left := remainder / 2
	right := remainder - left
	spad := string(pad)
	return strings.Repeat(spad, left) + s + strings.Repeat(spad, right)
}

func CleanWhitespace(text string) string {
	return strings.Join(strings.Fields(text), " ")
}

// ElideMiddle returns s at most width runes long. If s is longer than
// width, splits it in the middle, inserts an ellipsis, and removes runes so
// that it fits width.
func ElideMiddle(s string, width int) string {
	chars := []rune(s)
	size := len(chars)
	if size <= width { // it fits
		return s
	}
	diff := size - width
	left := diff / 2
	right := diff - left + 1
	mid := size / 2
	return string(chars[:mid-left]) + "…" + string(chars[mid+right:])
}

// LessFold returns true if string a is case-insensitively less than string
// b; otherwise returns false.
// This function can also be used to sort a slice of strings, e.g.,
// `slices.SortFunc(mystrings, utext.LessFold)`.
func LessFold(a, b string) bool {
	return strings.ToUpper(a) < strings.ToUpper(b)
}

// LongestCommonPrefix returns the longest common prefix (which could be ""
// if there isn't one).
// See also [LongestCommonPath]
func LongestCommonPrefix(lines []string) string {
	if len(lines) == 0 {
		return ""
	} else if len(lines) == 1 {
		return lines[0]
	}
	first := []rune(lines[0])
	prefix := make([]rune, 0, len(first))
outer:
	for i := range len(first) {
		r := first[i]
		for j := 1; j < len(lines); j++ {
			line := []rune(lines[j])
			if len(line)-1 < i || line[i] != r {
				break outer
			}
		}
		prefix = append(prefix, r)
	}
	return string(prefix)
}

// StringForSlice returns a string of sep-separated items.
// Mostly useful for tests.
func StringForSlice[T any](x []T, sep string) string {
	items := make([]string, 0, len(x))
	for _, v := range x {
		items = append(items, fmt.Sprint(v))
	}
	return strings.TrimSpace(strings.Join(items, sep))
}

// TitleCase returns the given text with the first letter of every "word"
// uppercased and the rest lowercased.
func TitleCase(text string) string {
	rx := regexp.MustCompile(`\b(\pL)`)
	return rx.ReplaceAllStringFunc(strings.ToLower(text),
		func(x string) string {
			return strings.ToUpper(x)
		})
}
