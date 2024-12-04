package utility

import "strings"

func SplitAndTrim(input string, separator string) (string, string, bool) {
	before, after, found := strings.Cut(input, separator)

	return strings.TrimSpace(before), strings.TrimSpace(after), found
}
