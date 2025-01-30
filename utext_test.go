// Copyright © 2024 Mark Summerfield. All rights reserved.

package utext

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

const (
	Text     = "  Text\twith   no\r\n   excess\nspace  €$•…   etc   \n"
	Expected = "Text with no excess space €$•… etc"
)

func Test_CleanWhitespace(t *testing.T) {
	cleaned := CleanWhitespace(Text)
	if Expected != cleaned {
		t.Errorf("expected %q,\ngot %q", Expected, cleaned)
	}
	cleaned = CleanWhitespace(cleaned)
	if Expected != cleaned {
		t.Errorf("expected %q,\ngot %q", Expected, cleaned)
	}
}

func Benchmark_CleanWhitespace(b *testing.B) {
	text1 := Text
	text2 := "  Text\twith   no\r\n\t   excess\nspace   "
	for range b.N {
		text1c := CleanWhitespace(text1)
		_ = CleanWhitespace(text1c)
		text2c := CleanWhitespace(text2)
		_ = CleanWhitespace(text2c)
	}
}

func TestLcPrefix1(t *testing.T) {
	items := []string{
		"/home/mark/app/go/utext",
		"/home/mark/app/py/accelhints", "/home/mark/app/rs",
	}
	if runtime.GOOS == "windows" {
		for i := range len(items) {
			items[i] = filepath.FromSlash(items[i])
		}
		prefix := LongestCommonPrefix(items)
		if prefix != `\home\mark\app\` {
			t.Errorf(`expected \home\mark\app\ got %q`, prefix)
		}
	} else {
		prefix := LongestCommonPrefix(items)
		if prefix != "/home/mark/app/" {
			t.Errorf("expected /home/mark/app/ got %q", prefix)
		}
	}
}

func TestLcPrefix2(t *testing.T) {
	items := []string{
		"/users/mark/app/go/utext",
		"/Users/mark/app/py/accelhints", "/home/mark/app/rs",
	}
	if runtime.GOOS == "windows" {
		for i := range len(items) {
			items[i] = filepath.FromSlash(items[i])
		}
		prefix := LongestCommonPrefix(items)
		if prefix != `\` {
			t.Errorf("expected \\ got %q", prefix)
		}
	} else {
		prefix := LongestCommonPrefix(items)
		if prefix != "/" {
			t.Errorf("expected / got %q", prefix)
		}
	}
}

func TestLcPrefix3(t *testing.T) {
	items := []string{
		"C:\\users\\mark\\app\\go\\utext",
		"/Users/mark/app/py/accelhints",
	}
	prefix := LongestCommonPrefix(items)
	if prefix != "" {
		t.Errorf("expected \"\" got %s", prefix)
	}
}

func TestLcPrefix4(t *testing.T) {
	items := []string{
		"mark/app/go/utext", "mark/app/py/accelhints",
		"mark/app/rs",
	}
	if runtime.GOOS == "windows" {
		for i := range len(items) {
			items[i] = filepath.FromSlash(items[i])
		}
		prefix := LongestCommonPrefix(items)
		if prefix != `mark\app\` {
			t.Errorf(`expected mark\app\ got %q`, prefix)
		}
	} else {
		prefix := LongestCommonPrefix(items)
		if prefix != "mark/app/" {
			t.Errorf("expected mark/app/ got %q", prefix)
		}
	}
}

func TestLcPrefix5(t *testing.T) {
	items := []string{"fan", "fate", "fame"}
	prefix := LongestCommonPrefix(items)
	if prefix != "fa" {
		t.Errorf("expected fa got %s", prefix)
	}
	items = []string{"elefan", "elefate", "elefame", "elefa"}
	prefix = LongestCommonPrefix(items)
	if prefix != "elefa" {
		t.Errorf("expected fa got %s", prefix)
	}
}

func TestLcPrefix6(t *testing.T) {
	items := []string{"bat", "vat", "cat"}
	prefix := LongestCommonPrefix(items)
	if prefix != "" {
		t.Errorf("expected \"\" got %s", prefix)
	}
}

func TestLessFold(t *testing.T) {
	if !LessFold("ABC", "abd") {
		t.Errorf("LessFold error #1")
	}
	if LessFold("ABD", "abc") {
		t.Errorf("LessFold error #2")
	}
	if LessFold("Able", "Ability") {
		t.Errorf("LessFold error #3")
	}
	if !LessFold("Ability", "Able") {
		t.Errorf("LessFold error #4")
	}
}

func TestStringForSlice(t *testing.T) {
	items := []int{1, 2, 4, 8, 16, -9, -7, 0, 12}
	expected := "1 2 4 8 16 -9 -7 0 12"
	actual := StringForSlice(items)
	if actual != expected {
		t.Errorf("expected %s got %s", expected, actual)
	}
}

func TestTitleCase(t *testing.T) {
	expected := "This And That Or Else"
	actual := TitleCase("THIS AND THAT OR ELSE")
	if actual != expected {
		t.Errorf("expected %q got %q", expected, actual)
	}
	actual = TitleCase("this and that or else")
	if actual != expected {
		t.Errorf("expected %q got %q", expected, actual)
	}
	actual = TitleCase(expected)
	if actual != expected {
		t.Errorf("expected %q got %q", expected, actual)
	}
}

func ExampleCentered() {
	s := Centered("The Title", ' ', 15)
	fmt.Printf("%q\n", s)
	s = Centered(" Heading ", '*', 15)
	fmt.Printf("%q\n", s)
	s = Centered(" Heading ", '=', 16)
	fmt.Printf("%q\n", s)
	s = Centered(" Heading ", '-', 17)
	fmt.Printf("%q\n", s)
	s = Centered("Too wide to center", '#', 12)
	fmt.Printf("%q\n", s)
	// Output:
	// "   The Title   "
	// "*** Heading ***"
	// "=== Heading ===="
	// "---- Heading ----"
	// "Too wide to center"
}

func ExampleElideMiddle() {
	s := ElideMiddle("This is short enough", 24)
	fmt.Printf("%q\n", s)
	t := "This is now far too long"
	for i := 14; i < 20; i++ {
		s = ElideMiddle(t, i)
		fmt.Printf("%d: %q\n", i, s)
	}
	// Output:
	// "This is short enough"
	// 14: "This is…o long"
	// 15: "This is …o long"
	// 16: "This is …oo long"
	// 17: "This is n…oo long"
	// 18: "This is n…too long"
	// 19: "This is no…too long"
}
