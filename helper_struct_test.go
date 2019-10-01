package helper

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestIsEmpty_OK(t *testing.T) {
    assert.Equal(t, true, IsEmpty(""), "Result not equal")
    assert.Equal(t, true, IsEmpty(nil), "Result not equal")
    assert.Equal(t, true, IsEmpty(struct{}{}), "Result not equal")
}

func TestIsEmpty_Error(t *testing.T) {
    assert.NotEqual(t, true, IsEmpty(struct{Id  int }{Id: 11}), "Result not equal")
}