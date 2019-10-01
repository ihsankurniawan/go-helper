package helper

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var img = "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"

func TestBase64ToImg_OK(t *testing.T) {
    _, err := Base64ToImg("/tmp/tempfile", img)
    assert.Nil(t, err)
}

func TestBase64ToImg_Error(t *testing.T) {
    _, err := Base64ToImg("/tmp/tempfile", "wrong base64 encoded")
    assert.NotNil(t, err)
}
