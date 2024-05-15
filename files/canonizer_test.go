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

var containsTestResults = []struct{
    list     []string
    element  string
    expected bool
    message  string
}{
    {[]string{}, "123", false, "empty slice shouldn't contain anything"},
    {[]string{"123", "234", "345"}, "234", true, "should be true when the element is found"},
    {[]string{"123", "234", "345"}, "456", false, "should be false when the element isn't found"},
}

var deduplicateTestResults = []struct{
    argument []string
    expected []string
    message  string
}{
    {[]string{"123", "234"}, []string{"123", "234"}, "list with no duplicates should remain the same"},
    {[]string{"123", "123", "234"}, []string{"123", "234"}, "duplicate items should be removed"},
}

var removeEmptyTestResults = []struct{
    argument []string
    expected []string
    message  string
}{
    {[]string{"123", "", "234"}, []string{"123", "234"}, "should remove empty strings"},  
}    
    
func TestCanonizer(t *testing.T) {
    for _, res := range canonizerTestResults {
        assert.Equal(t, res.expected, canonize(res.argument), res.message)
    }
}

func TestContains(t *testing.T) {
    for _, res := range containsTestResults {
        assert.Equal(t, res.expected, contains(res.list, res.element), res.message)
    }
}

func TestDeduplicate(t *testing.T) {
    for _, res := range deduplicateTestResults {
        assert.Equal(t, res.expected, deduplicate(res.argument), res.message)
    }
}

func TestRemoveEmpty(t *testing.T) {
    for _, res := range removeEmptyTestResults {
        assert.Equal(t, res.expected, removeEmpty(res.argument), res.message)
    }
}