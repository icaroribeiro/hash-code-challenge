package middlewares

import (
    "strings"
)

func CustomHeaderMatcher(key string) (string, bool) {
    switch strings.ToUpper(key) {
    case "X-USER-ID":
        return key, true
    default:
        return key, false
    }
}
