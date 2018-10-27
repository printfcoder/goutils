package test

import (
	"github.com/printfcoder/goutils/intutils"
	"testing"
)

func TestRandInt(t *testing.T) {

	out := intutils.RandInt(1000000000, 9999999999)
	if out < 0 {
		t.Error(out)
	}

	t.Log(out)

	out = intutils.RandInt(1000000000, 9999999999)
	if out < 0 {
		t.Error(out)
	}

	t.Log(out)
}

func TestGetNowDate(t *testing.T) {
	out := intutils.GetNowDate()
	if out < 0 {
		t.Error(out)
	}
	t.Log(out)
}

func TestRandIntWithDateAsPrefix(t *testing.T) {
	out := intutils.RandIntWithDateAsPrefix(1000000000, 9999999999)
	if out < 0 {
		t.Error(out)
	}
	t.Log(out)
}
