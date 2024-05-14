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

func TestListWithSimplePrefixMatches(t *testing.T) {
    sample := []string{"911", "91101", "913"}
    assert.False(t, isConsistent(sample), sample)
}

func TestListWithNoSimplePrefixMatches(t *testing.T) {
    sample := []string{"911", "912", "913"}
    assert.True(t, isConsistent(sample), sample)
}
