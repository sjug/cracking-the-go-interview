package onefive

import "strings"

// Function StringsSpaceReplace uses the standard function strings.Replace
func StringsSpaceReplace(s string) string {
	// If n < 0, there is no limit on the number of replacements
	return strings.Replace(s, " ", "%20", -1)
}

// Function SpaceReplace is a butchery of the standard library to try to
// comprehend how the actual replace function works.
func SpaceReplace(s string) string {
	count := strings.Count(s, " ")
	t := make([]byte, len(s)+count*2)
	w := 0
	start := 0
	for i := 0; i < count; i++ {
		j := start
		j += strings.Index(s[start:], " ")
		w += copy(t[w:], s[start:j])
		w += copy(t[w:], "%20")
		start = j + 1
	}
	w += copy(t[w:], s[start:])
	return string(t[0:w])
}
