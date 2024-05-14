package phonelist

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCanonizeNumericString(t *testing.T) {
    assert.Equals(t, "123", canonize("123"))
}