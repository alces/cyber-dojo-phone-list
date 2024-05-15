package phonelist

import (
    "testing"
    
    "github.com/stretchr/testify/assert"
)

var sortTestResults = []struct {
    unsorted []string
    sorted   []string
    message  string
}{
    {[]string{}, []string{}, "should work with an empty slice"},
    {[]string{"123", "4", "56"}, []string{"4", "56", "123"}, "a slice should be sorted by element length"},
}

func TestSortByLen(t *testing.T) {
    for _, res := range sortTestResults {
        assert.Equal(t, res.sorted, sortByLen(res.unsorted), res.message)
    }
}

func TestSortByLenPurity(t *testing.T) {
    unsorted := []string{"123", "4", "56"}
    expected := []string{"123", "4", "56"}

    _ = sortByLen(unsorted)

    assert.Equal(t, expected, unsorted, "the original slice should remain unchanged")
}