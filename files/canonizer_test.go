package phonelist

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var canonizerTestResults = []struct{
    argument string
    expected string
    message  string
}{
    {"123", "123", "numeric string shouldn't change"},
    {"(123) 45-67", "1234567", "all non-numeric characters should be removed"},
}

var duplicateTestResults = []struct{
    argument []string
    expected []string
    message  string
}{
    {[]string{"123", "234"}, []string{"123", "234"}, "list with no duplicates should remain the same"},
    {[]string{"123", "123", "234"}, []string{"123", "234"}, "duplicate items should be removed"},
}

func TestCanonizer(t *testing.T) {
    for _, res := range canonizerTestResults {
        assert.Equal(t, res.expected, canonize(res.argument), res.message)
    }
}

func TestDeduplicate(t *testing.T) {
    for _, res := range duplicateTestResults {
        assert.Equal(t, res.expected, deduplicate(res.argument), res.message)
    }
}