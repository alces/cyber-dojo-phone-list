package phonelist

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestEmptyListShouldBeConsistent(t *testing.T) {
    assert.True(t, isConsistent([]string{}))
}

func TestOneElementListShouldBeConsistent(t *testing.T) {
    assert.True(t, isConsistent([]string{"911"}))
}

func TestListWithNoSimplePrefixMatches(t *testing.T) {
    assert.True(t, isConsistent([]string{"911", "912", "913"}))
}
