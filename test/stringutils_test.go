package test

import (
	"testing"

	"github.com/printfcoder/goutils/stringutils"
)

func TestIsEmpty(t *testing.T) {
	testStr := "999ss999asf2323323asd"

	out := stringutils.IsEmpty(testStr)
	if out {
		t.Error(out)
	}

	t.Log(out)
}

func TestIsAnyEmpty(t *testing.T) {

	out := stringutils.IsAnyEmpty("", "2")
	if !out {
		t.Error(out)
	}

	t.Log(out)
}

func TestIsNoneEmpty(t *testing.T) {

	out := stringutils.IsNoneEmpty("", "2")
	if out {
		t.Error(out)
	}

	t.Log(out)
}

func TestIsBlank(t *testing.T) {

	out := stringutils.IsBlank("")
	if !out {
		t.Error(out)
	}

	out = stringutils.IsBlank(" ")
	if !out {
		t.Error(out)
	}

	t.Log(out)
}

func TestIsNotBlank(t *testing.T) {

	out := stringutils.IsNotBlank("")
	if out {
		t.Error(out)
	}

	out = stringutils.IsNotBlank(" ")
	if out {
		t.Error(out)
	}

	t.Log(out)
}

func TestIsAnyBlank(t *testing.T) {

	out := stringutils.IsAnyBlank(" ", "d")
	if !out {
		t.Error(out)
	}

	out = stringutils.IsAnyBlank("dd", "ee")
	if out {
		t.Error(out)
	}

	t.Log(out)
}

func TestTruncate(t *testing.T) {

	out := stringutils.Truncate("012345678", 3)
	if out != "012" {
		t.Error(out)
	}

	out = stringutils.Truncate("012345678", 9)
	if out != "012345678" {
		t.Error(out)
	}

	out = stringutils.Truncate("012345678", 8)
	if out != "01234567" {
		t.Error(out)
	}

	out = stringutils.Truncate("012345678", 0)
	if out != "" {
		t.Error(out)
	}

	t.Log(out)
}

func TestTruncateFromWithMaxWith(t *testing.T) {

	out, _ := stringutils.TruncateFromWithMaxWith("0123456789", 3, 7)
	if out != "3456789" {
		t.Error(out)
	}

	t.Log(out)
}

func TestSubString(t *testing.T) {

	out, _ := stringutils.SubString("0123456789", 3)
	if out != "3456789" {
		t.Error(out)
	}

	t.Log(out)
}

func TestSubStringBetween(t *testing.T) {

	out, err := stringutils.SubStringBetween("0123456789", 0, 9)
	if out != "012345678" {
		t.Error(out)
	}

	if err != nil {
		t.Error(err)
	}

	out, err = stringutils.SubStringBetween("120.00", 0, 2)
	if out != "12" {
		t.Error(out)
	}

	out, err = stringutils.SubStringBetween("120.00", 0, 0)
	if out != "" {
		t.Error(out)
	}

	t.Log(out)
}

func TestCharAt(t *testing.T) {

	out, err := stringutils.CharAt("0123456789", 2)
	if out != "2" {
		t.Error(out)
	}

	if err != nil {
		t.Error(err)
	}

	out, err = stringutils.CharAt("01中国456789", 2)
	if out != "中" {
		t.Error(out)
	}

	if err != nil {
		t.Error(err)
	}

	out, err = stringutils.CharAt("01中国456789", 4)
	if out != "4" {
		t.Error(out)
	}

	if err != nil {
		t.Error(err)
	}

	t.Log(out)
}

func TestIndexOf(t *testing.T) {

	out := stringutils.IndexOf("0123456789", "2")
	if out != 2 {
		t.Error(out)
	}

	out = stringutils.IndexOf("01中国456789", "中")
	if out != 2 {
		t.Error(out)
	}

	out = stringutils.IndexOf("01中国456789", "d中")
	if out != -1 {
		t.Error(out)
	}

	t.Log(out)
}

func TestStripStart(t *testing.T) {

	out := stringutils.StripStart("0123456789", "d081d2d348")
	if out != "56789" {
		t.Error(out)
	}

	out = stringutils.StripStart("01中国456789", "01中国456789")
	if out != "" {
		t.Error(out)
	}

	out = stringutils.StripStart("yxabc  ", "xyz")
	if out != "abc  " {
		t.Error(out)
	}

	t.Log(out)
}

func TestStripEnd(t *testing.T) {

	out := stringutils.StripEnd("0123456789", "d081d2d348")
	if out != "0123456789" {
		t.Error(out)
	}

	out = stringutils.StripEnd("01中国456789", "01中国456789")
	if out != "" {
		t.Error(out)
	}

	out = stringutils.StripEnd("120.00", ".0")
	if out != "12" {
		t.Error(out)
	}

	t.Log(out)
}

func TestStripWithChar(t *testing.T) {
	out := stringutils.StripWithChar("  abcyx", "xyz")
	if out != "  abc" {
		t.Error(out)
	}
}

func TestStrip(t *testing.T) {

	out := stringutils.Strip("")
	if out != "" {
		t.Error(out)
	}

	out = stringutils.Strip("   ")
	if out != "" {
		t.Error(out)
	}

	out = stringutils.Strip("abc")
	if out != "abc" {
		t.Error(out)
	}

	out = stringutils.Strip("  abc")
	if out != "abc" {
		t.Error(out)
	}

	out = stringutils.Strip("abc  ")
	if out != "abc" {
		t.Error(out)
	}

	out = stringutils.Strip(" abc ")
	if out != "abc" {
		t.Error(out)
	}

	out = stringutils.Strip(" ab c ")
	if out != "ab c" {
		t.Error(out)
	}
}

func TestToCharArray(t *testing.T) {

	_1 := []string{" ", "a", "b", " ", "c", " "}
	out := stringutils.ToCharArray(" ab c ")
	if len(out) != len(_1) {
		t.Error(len(out), len(_1))
	}

	for i, v := range _1 {
		if v != _1[i] {
			t.Error(v, _1[i])
		}
	}

	_2 := []string{" ", "中", "国", "a"}
	out = stringutils.ToCharArray(" 中国a")
	if len(out) != len(_2) {
		t.Error(len(out), len(_2))
	}

	for i, v := range _2 {
		if v != _2[i] {
			t.Error(v, _2[i])
		}
	}
}

func TestRegionMatches(t *testing.T) {

	out := stringutils.RegionMatches(" ab c ", true, 0, " aB c ", 0, len(" ab c "))
	if !out {
		t.Error(out)
	}

	out = stringutils.RegionMatches(" ab c ", true, 0, "aB c ", 0, len(" ab c "))
	if out {
		t.Error(out)
	}
}

func TestEqualsIgnoreCase(t *testing.T) {
	out := stringutils.EqualsIgnoreCase(" ab c ", " aB c ")
	if !out {
		t.Error(out)
	}

	out = stringutils.EqualsIgnoreCase(" ab c ", " aBc ")
	if out {
		t.Error(out)
	}

}

func TestCompare(t *testing.T) {

	out := stringutils.Compare("a", "b")
	if out > 0 {
		t.Error(out)
	}

	out = stringutils.Compare("a", "a")
	if out != 0 {
		t.Error(out)
	}

	// the rune of char '中' is '20013'
	// and char '国' has rune '22269'
	out = stringutils.Compare("国", "中")
	if out < 0 {
		t.Error(out)
	}

	// the rune of char '中' is '20013'
	// and char '国' has rune '22269'
	out = stringutils.Compare("中国国", "中国中")
	if out < 0 {
		t.Error(out)
	}
}

func TestCompareIgnoreCase(t *testing.T) {

	out := stringutils.CompareIgnoreCase("ABC", "abc")
	if out != 0 {
		t.Error(out)
	}

	out = stringutils.CompareIgnoreCase("ABC", "abcb")
	if out >= 0 {
		t.Error(out)
	}

	out = stringutils.CompareIgnoreCase("ABC国", "abc国")
	if out != 0 {
		t.Error(out)
	}
}

func TestEqualsAny(t *testing.T) {

	out := stringutils.EqualsAny("ABC", "abc", "abcd")
	if out {
		t.Error(out)
	}

	out = stringutils.EqualsAnyIgnoreCase("ABC", "abc", "abcd")
	if !out {
		t.Error(out)
	}

	out = stringutils.EqualsAny("ABC国", "abc国", "adfsf")
	if out {
		t.Error(out)
	}
}

func TestIndexOfFromIndex(t *testing.T) {
	out := stringutils.IndexOfFromIndex("abcdabcabc", "abc", 1)
	if out != 4 {
		t.Error(out)
	}

	out = stringutils.IndexOfFromIndex("abcdabcabc", "abce", 0)
	if out != -1 {
		t.Error(out)
	}

	out = stringutils.IndexOfFromIndex("ABC国国国", "国", 2)
	if out != 3 {
		t.Error(out)
	}
}

func TestOrdinalIndexOf(t *testing.T) {

	// stringUtils.OrdinalIndexOf("abababab", "abab", 1) = 0
	// stringUtils.OrdinalIndexOf("abababab", "abab", 2) = 2
	// stringUtils.OrdinalIndexOf("abababab", "abab", 3) = 4
	// stringUtils.OrdinalIndexOf("abababab", "abab", 4) = -1

	out := stringutils.OrdinalIndexOf("abababab", "abab", 1, false)
	if out != 0 {
		t.Error(out)
	}

	out = stringutils.OrdinalIndexOf("abababab", "abab", 2, false)
	if out != 2 {
		t.Error(out)
	}

	out = stringutils.OrdinalIndexOf("abababab", "abab", 3, false)
	if out != 4 {
		t.Error(out)
	}

	out = stringutils.OrdinalIndexOf("abababab", "abab", 3, true)
	if out != 0 {
		t.Error(out)
	}

	out = stringutils.OrdinalIndexOf("中国中国中国中国", "中国中国", 3, true)
	if out != 0 {
		t.Error(out)
	}
}

func TestContains(t *testing.T) {

	out := stringutils.Contains("ABC国国国", "国")
	if out != true {
		t.Error(out)
	}

	out = stringutils.Contains("ABC国国国", "21")
	if out != false {
		t.Error(out)
	}

	out = stringutils.Contains("", "")
	if out != true {
		t.Error(out)
	}

	out = stringutils.Contains("abc", "")
	if out != true {
		t.Error(out)
	}

	out = stringutils.Contains("abc", "a")
	if out != true {
		t.Error(out)
	}

	out = stringutils.Contains("abc", "z")
	if out != false {
		t.Error(out)
	}

}

func TestContainsIgnoreCase(t *testing.T) {

	out := stringutils.ContainsIgnoreCase("ABC国国国", "国")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsIgnoreCase("abc", "")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsIgnoreCase("abc", "a")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsIgnoreCase("abc", "z")
	if out == true {
		t.Error(out)
	}

	out = stringutils.ContainsIgnoreCase("abc", "A")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsIgnoreCase("abc", "Z")
	if out == true {
		t.Error(out)
	}
}

func TestIsWhitespace(t *testing.T) {

	// space
	out := stringutils.IsWhitespace(" ")
	if out != true {
		t.Error(out)
	}

	// enter
	out = stringutils.IsWhitespace(`
`)
	if out != true {
		t.Error(out)
	}

	// tab
	out = stringutils.IsWhitespace("	")
	if out != true {
		t.Error(out)
	}

	// tab
	out = stringutils.IsWhitespace("d")
	if out == true {
		t.Error(out)
	}
}

func TestIndexOfAny(t *testing.T) {

	out := stringutils.IndexOfAny("", "ddd")
	if out != -1 {
		t.Error(out)
	}

	out = stringutils.IndexOfAny("zzabyycdxx", "z", "a")
	if out != 0 {
		t.Error(out)
	}

	out = stringutils.IndexOfAny("zzabyycdxx", "b", "y")
	if out != 3 {
		t.Error(out)
	}

	out = stringutils.IndexOfAny("中国Golang的粉丝", "黑", "丝")
	if out != 10 {
		t.Error(out)
	}
}

func TestContainsAny(t *testing.T) {

	out := stringutils.ContainsAny("", "d")
	if out == true {
		t.Error(out)
	}

	out = stringutils.ContainsAny("zzabyycdxx", "z", "a")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsAny("zzabyycdxx", "b", "y")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsAny("zzabyycdxx", "z", "y")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsAny("aba", "z")
	if out == true {
		t.Error(out)
	}

	out = stringutils.ContainsAny("中国Golang的粉丝", "黑", "丝")
	if out != true {
		t.Error(out)
	}
}

func TestIndexOfAnyBut(t *testing.T) {

	out := stringutils.IndexOfAnyBut("zzabyycdxx", "z", "a")
	if out != 3 {
		t.Error(out)
	}

	out = stringutils.IndexOfAnyBut("aba", "z")
	if out != 0 {
		t.Error(out)
	}

	out = stringutils.IndexOfAnyBut("aba", "a", "b")
	if out != -1 {
		t.Error(out)
	}

	out = stringutils.IndexOfAnyBut("中国Golang的粉丝", "黑", "丝")
	if out != 0 {
		t.Error(out)
	}
}

func TestContainsOnly(t *testing.T) {
	out := stringutils.ContainsOnly("ab", "")
	if out != false {
		t.Error(out)
	}

	out = stringutils.ContainsOnly("abab", "a", "b", "c")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsOnly("ab1", "a", "b", "c")
	if out != false {
		t.Error(out)
	}

	out = stringutils.ContainsOnly("abz", "a", "b", "c")
	if out != false {
		t.Error(out)
	}
}

func TestContainsNone(t *testing.T) {
	out := stringutils.ContainsNone("ab", "")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsNone("abab", "x", "y", "z")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsNone("ab1", "x", "y", "z")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsNone("abz", "x", "y", "z")
	if out != false {
		t.Error(out)
	}
}

func TestLastIndexOfAny(t *testing.T) {
	out := stringutils.LastIndexOfAny("zzabyycdxx", "ab", "cd")
	if out != 6 {
		t.Error(out)
	}

	out = stringutils.LastIndexOfAny("zzabyycdxx", "cd", "ab")
	if out != 6 {
		t.Error(out)
	}

	out = stringutils.LastIndexOfAny("zzabyycdxx", "mn", "op")
	if out != -1 {
		t.Error(out)
	}

	out = stringutils.LastIndexOfAny("zzabyycdxx", "mn", "")
	if out != -1 {
		t.Error(out)
	}
}


