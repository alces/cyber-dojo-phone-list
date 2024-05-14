package phonelist

import (
    "regexp"
    "strings"
)

func isConsistent(p []string) bool {
    for i := 0; i < len(p) - 1; i++ {
        for j := i + 1; j < len(p); j++ {
            if strings.HasPrefix(canonize(p[j]), canonize(p[i])) {
                return false
            }
        }
    }
    
    return true
}

func canonize(mixed string) string {
    re := regexp.MustCompile("[^0-9]+")
    
    return string(re.ReplaceAll([]byte(mixed), []byte("")))
}