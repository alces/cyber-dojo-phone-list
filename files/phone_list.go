package phonelist

import (
    "strings"
)

func isConsistent(p []string) bool {
    if len(p) < 2 {
        return true
    }
    
    c := deduplicate(sortByLen(p))

    for i := 0; i < len(c); i++ {
        c[i] = canonize(c[i])
    }
    
    for i := 0; i < len(c) - 1; i++ {
        for j := i + 1; j < len(c); j++ {            
            if strings.HasPrefix(c[j], c[i]) {
                return false
            }
        }
    }
    
    return true
}
