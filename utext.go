// Copyright © 2024 Mark Summerfield. All rights reserved.

package utext

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/mark-summerfield/unumber"
)

//go:embed Version.dat
var Version string

func CleanWhitespace(text string) string {
	return strings.Join(strings.Fields(text), " ")
}

func Commas[I unumber.Integer](i I) string {
	sign := ""
	value := fmt.Sprint(i) // Can't use Itoa() with Integer
	if value[0] == '-' {
		sign = "-"
		value = value[1:]
	}
	for i := len(value) - 3; i >= 0; i -= 3 {
		value = value[:i] + "," + value[i:]
	}
	return sign + strings.TrimPrefix(value, ",")
}

// LessFold returns true if string a is case-insensitively less than string
// b; otherwise returns false.
// This function can also be used to sort a slice of strings, e.g.,
// `slices.SortFunc(mystrings, gong.LessFold)`.
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

// StringForSlice returns a string of space-separated items.
// Mostly useful for tests.
func StringForSlice[T any](x []T) string {
	items := make([]string, len(x))
	for _, n := range x {
		items = append(items, fmt.Sprintf("%v ", n))
	}
	return strings.TrimSpace(strings.Join(items, ""))
}
