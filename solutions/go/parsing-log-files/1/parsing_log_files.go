package parsinglogfiles

import (
    re "regexp"
)

var validRegex = re.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
func IsValidLine(text string) bool {
	return validRegex.MatchString(text)
}

var splitRegex = re.MustCompile(`<[-~*=]*>`)
func SplitLogLine(text string) []string {
	return splitRegex.Split(text, -1)
}

var quotedPasswordRegex = re.MustCompile(`(?i)('|").*password.*('|")`)
func CountQuotedPasswords(lines []string) int {
    count := 0
    for _, line := range(lines) {
		if quotedPasswordRegex.FindString(line) != "" {
            count += 1
        }
    }
    return count
}

var endOfLineRegex = re.MustCompile(`end-of-line\d+`)
func RemoveEndOfLineText(text string) string {
	return string(endOfLineRegex.ReplaceAll([]byte(text), []byte{}))
}


var userRegex = re.MustCompile(`User\s+(\w+)`)

func TagWithUserName(lines []string) []string {
	result := make([]string, len(lines))
	for idx, line := range lines {
		matches := userRegex.FindStringSubmatch(line)
		if matches == nil || len(matches) < 2 {
			result[idx] = line
		} else {
			username := matches[1]
			result[idx] = "[USR] " + username + " " + line
		}
	}
	return result
}