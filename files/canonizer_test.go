package phonelist

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCanonizeNumericString(t *testing.T) {
    assert.Equal(t, "123", canonize("123"))
}

func TestCanonizeMixedString(t *testing.T) {
    assert.Equal(t, "1234567", canonize("(123)45-67"))
}