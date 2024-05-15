package phonelist

import (
    "regexp"
)

func canonize(mixed string) string {
    re := regexp.MustCompile("[^0-9]+")
    
    return string(re.ReplaceAll([]byte(mixed), []byte("")))
}

func deduplicate(dirty []string) []string {
    return dirty
}