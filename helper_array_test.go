package helper

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var neddle      = 10
var lostNeddle  = 12
var haystack    = []int{1, 5, 10, 15}

var neddleStr      = "foo"
var lostNeddleStr  = "bar"
var haystackStr    = []string{"foo", "lorem", "ipsum"}

func TestInArray_OK(t *testing.T) {
    rs, idx := InArray(neddle, haystack)
    assert.Equal(t, true, rs, "Not InArray")
    assert.Equal(t, 2, idx, "Not True Index")
}

func TestInArray_Error(t *testing.T) {
    rs, idx := InArray(lostNeddle, haystack)
    assert.NotEqual(t, true, rs, "Not InArray")
    assert.NotEqual(t, 2, idx, "Not True Index")
}

func TestInArrayNoIndex_OK(t *testing.T) {
    rs := InArrayNoIndex(neddle, haystack)
    assert.Equal(t, true, rs, "Not InArray")
}

func TestInArrayNoIndex_Error(t *testing.T) {
    rs := InArrayNoIndex(lostNeddle, haystack)
    assert.NotEqual(t, true, rs, "Not InArray")
}

func TestInArrayContains_OK(t *testing.T) {
    rs := InArrayContains(neddleStr, haystackStr)
    assert.Equal(t, true, rs, "Not InArray")
}

func TestInArrayContains_Error(t *testing.T) {
    rs := InArrayContains(lostNeddleStr, haystackStr)
    assert.NotEqual(t, true, rs, "Not InArray")
}