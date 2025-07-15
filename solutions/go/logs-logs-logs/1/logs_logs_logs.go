package logs

import (
    s "strings"
)
var runeToApp = map[rune]string {
    '❗': "recommendation",
    '🔍': "search",
    '☀': "weather",
}
const logRunes = "☀🔍❗"
// Application identifies the application emitting the given log.
func Application(log string) string {
	idx := s.IndexAny(log, logRunes)
	if idx != -1{
        return runeToApp[[]rune(log)[idx]]
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return s.Replace(log, string(oldRune), string(newRune), -1)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
