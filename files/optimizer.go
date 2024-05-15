package phonelist

import (
    "sort"
)

type ByLen []string

func (b ByLen) Len() int {
    return len(b)
}

func (b ByLen) Swap(i, j int) {
    b[i], b[j] = b[j], b[i]
}

func (b ByLen) Less(i, j int) bool {
    return len(b[i]) < len(b[j])
}

func sortByLen(a []string) []string {
    sort.Sort(ByLen(a))
    
    return a
}