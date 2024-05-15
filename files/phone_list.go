package phonelist

import (
    "strings"
)

func isConsistent(phoneList []string) bool {
    if len(phoneList) < 2 {
        return true
    }
    
    sorted := sortByLen(deduplicate(phoneList))

    for i := 0; i < len(sorted); i++ {
        sorted[i] = canonize(sorted[i])
    }
    
    cleaned := removeEmpty(sorted)
 
    for i := 0; i < len(cleaned) - 1; i++ {
        for j := i + 1; j < len(cleaned); j++ {            
            if strings.HasPrefix(cleaned[j], cleaned[i]) {
                return false
            }
        }
    }
    
    return true
}
