package test

import (
	"github.com/printfcoder/goutils/crypto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPwd(t *testing.T) {
	out, err := cryptoutils.NewPwd("1234567", 10)
	assert.Nil(t, err)
	t.Log(out)
}

func TestIsPwdOK(t *testing.T) {

	testStr := "$2a$10$vJ.4bVx.ZmDXmyR.Ayj//OvC0grffAOYRq4LoQiwT/A3JUqVPrNHS"
	out := cryptoutils.IsPwdOK("1234567", testStr)
	assert.True(t, out)

	testStr = "$2a$14$k.DOsS3fhvKAnrq0FsciZufrAELJKykeGuRH7wsSoQ7DWgJc.b3.2"
	out = cryptoutils.IsPwdOK("123456", testStr)
	assert.False(t, out)

}
