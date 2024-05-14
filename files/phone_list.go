package phonelist

import (
    "strings"
)

func isConsistent(p []string) bool {
    if len(p) <= 1 {
        return true
    }
    
    for i := 0; i < len(p) - 1; i++ {
        for j := i + 1; j < len(p); j++ {
            if strings.HasPrefix(p[j], p[i]) {
                return false
            }
        }
    }
    
    return true
}
