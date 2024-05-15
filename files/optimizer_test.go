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
    {[]string{"123", "4", "56"}, []string{"4", "56", "123"}, "a slice should be sorted by element length"},
}

func TestSortByLen(t *testing.T) {
    for _, res := range sortTestResults {
        assert.Equal(t, res.sorted, sortByLen(res.unsorted), res.message)
    }
}