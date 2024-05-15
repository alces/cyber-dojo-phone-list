package phonelist

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"
)

var phoneListTestResults = []struct {
    list     []string
    expected bool
    message  string
}{
    {[]string{}, true, "empty list should be consistent"},
    {[]string{"911"}, true, "one-element list should be consistent"},
    {[]string{"911", "91101", "913"}, false, "a list with literal matches should be inconsistent"},
    {[]string{"911", "912", "913"}, true, "a list with no matches should be consistent"},
    {[]string{"911", "91-101", "913"}, false, "a list with matches after canonization should be inconsistent"},
    {[]string{"91101", "911", "913"}, false, "an incosistent list in different order should be inconsistent"},
}

func TestPhoneList(t *testing.T) {
    for _, res := range phoneListTestResults {
        assert.Equal(t, res.expected, isConsistent(res.list), fmt.Sprintf("%#v - %s", res.list, res.message))
    }
}
