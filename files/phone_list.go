package phonelist

import (
    "regexp"
    "strings"
)

func isConsistent(p []string) bool {
    c := make([]string, len(p))

    for i := 0; i < len(p); i++ {
        c[i] = canonize(p[i])
    }
    
    for i := 0; i < len(p) - 1; i++ {
        for j := i + 1; j < len(p); j++ {
            if strings.HasPrefix(c[j], c[i]) {
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