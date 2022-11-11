package util

import "strings"

func RemoveEmptyLines(text string) string {

	result := make([]string, 0)
	for _, line := range strings.Split(text, "\n") {
		if len(strings.TrimSpace(line)) != 0 {
			result = append(result, line)
		}
	}
	return strings.Join(result, "\n")
}
