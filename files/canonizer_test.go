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
    
func TestCanonizer(t *testing.T) {
    for _, res := range canonizerTestResults {
        assert.Equal(t, res.expected, canonize(res.argument), res.message)
    }
}
