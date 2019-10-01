package helper

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// We cannot test str result because of random
func TestStringRandomWithCharset_OK(t *testing.T) {
    str := StringRandomWithCharset(10, "abcdef")
    assert.Equal(t, 10, len(str), "Length not equal")
}

func TestStringRandomWithCharset_Error(t *testing.T) {
    str := StringRandomWithCharset(10, "abcdef")
    assert.NotEqual(t, 11, len(str), "Length not equal")
}

// We cannot test str result because of random
func TestStringRandom_OK(t *testing.T) {
    str := StringRandom(10)
    assert.Equal(t, 10, len(str), "Length not equal")
}

func TestStringRandom_Error(t *testing.T) {
    str := StringRandom(10)
    assert.NotEqual(t, 11, len(str), "Length not equal")
}

func TestStringIsJson_OK(t *testing.T) {
    rs := StringIsJSON("{\"id\" : 10}")
    assert.Equal(t, true, rs, "Not JSON")
}

func TestStringIsJson_Error(t *testing.T) {
    rs := StringIsJSON("{id : 10}")
    assert.NotEqual(t, true, rs, "Not JSON")
}


func TestFormatToIdrCurrency_OK(t *testing.T) {
    rs := FormatToIdrCurrency(100000)
    assert.Equal(t, "100.000", rs, "Wrong currency")
}

func TestFormatToIdrCurrency_Error(t *testing.T) {
    rs := FormatToIdrCurrency(1000000)
    assert.NotEqual(t, "1000000", rs, "Wrong currency")
}


func TestWhiteSpaceRemove_OK(t *testing.T) {
    rs := WhiteSpaceRemove(" lorem ipsum ")
    assert.Equal(t, "loremipsum", rs, "Wrong format")
}

func TestWhiteSpaceRemove_Error(t *testing.T) {
    rs := WhiteSpaceRemove(" lorem ipsum ")
    assert.NotEqual(t, " lorem ipsum ", rs, "Wrong format")
}

