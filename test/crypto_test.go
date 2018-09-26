package test

import (
	"github.com/printfcoder/goutils/crypto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPwd(t *testing.T) {
	out, err := cryptoutils.NewPwd("1234567")
	assert.Nil(t, err)
	t.Log(out)
}

func TestIsPwdOK(t *testing.T) {

	testStr := "$2a$14$SiorEqyfbs8GEWg1OyrLE.61aOUA69BX0GVMoYKU7MzjbsWPuntCS"
	out := cryptoutils.IsPwdOK("123456", testStr)
	assert.True(t, out)

	testStr = "$2a$14$k.DOsS3fhvKAnrq0FsciZufrAELJKykeGuRH7wsSoQ7DWgJc.b3.2"
	out = cryptoutils.IsPwdOK("123456", testStr)
	assert.False(t, out)

}
