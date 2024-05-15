package phonelist

import (
    "sort"
)

type ByLen []string

func (ByLen b) Len() int {
    return len(b)
}

func (ByLen b) Swap(i, j int) {
    b[i], b[j] = b[j], b[i]
}

func (ByLen b) Less(i, j int) bool {
    return len(b[i]) < len(b[j])
}

func sortByLen(a []string) []string {
    return sort.Sort(ByLen(a))
}