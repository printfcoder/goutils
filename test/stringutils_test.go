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
