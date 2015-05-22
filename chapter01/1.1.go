package oneone

// Function isUnique returns a boolean value indicating if a string has all unique characters
func IsUnique(s string) bool {
	m := make(map[rune]bool, len(s))
	for _, r := range s {
		if m[r] {
			return false
		}
		m[r] = true
	}
	return true
}
