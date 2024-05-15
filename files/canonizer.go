package phonelist

import (
    "regexp"
)

func canonize(mixed string) string {
    re := regexp.MustCompile("[^0-9]+")
    
    return string(re.ReplaceAll([]byte(mixed), []byte("")))
}

func deduplicate(dirty []string) []string {
    clean := make([]string, 0, len(dirty))
    
    for _, v := range dirty {
        if !contains(clean, v) {
            clean = append(clean, v)
        }
    }
    
    return clean
}

func contains(slice []string, element string) bool {
    for _, v := range slice {
        if v == element {
            return true
        }
    }
    
    return false
}