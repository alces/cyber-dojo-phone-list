package phonelist

import (
    "strings"
)

func isConsistent(p []string) bool {
    if len(p) < 2 {
        return true
    }
    
    c := deduplicate(p)

    for i := 0; i < len(c); i++ {
        c[i] = canonize(c[i])
    }
    
    for i := 0; i < len(c); i++ {
        for j := 0; j < len(c); j++ {
            if j == i {
                continue
            }
            
            if strings.HasPrefix(c[j], c[i]) {
                return false
            }
        }
    }
    
    return true
}
