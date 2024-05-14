package phonelist

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestEmptyListShouldBeConsistent(t *testing.T) {
    assert.True(t, isConsistent([]string{}))
}