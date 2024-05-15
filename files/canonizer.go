package phonelist

import (
    "regexp"
    "slices"
)

func canonize(mixed string) string {
    re := regexp.MustCompile("[^0-9]+")
    
    return string(re.ReplaceAll([]byte(mixed), []byte("")))
}

func deduplicate(dirty []string) []string {
    clean := make([]string, len(dirty), 0)
    
    for _, v := range dirty {
        if !slices.Contains(clean, v) {
            clean = append(clean, v)
        }
    }
    
    return clean
}