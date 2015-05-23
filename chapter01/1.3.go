package onethree

// In Golang strings are immutable so in order to "mutate" them we need to create either a
// new byte slice or rune slice. Since a UTF-8 character may occupy more than 1 byte we use
// the rune slice.

// Function RemoveDuplicates will remove duplicates from a string
// Will change the order of string characters
func RemoveDuplicates(s string) string {
	r := []rune(s)
	length := len(s) - 1
	for i := 0; i < length; i++ {
		for j := i + 1; j <= length; j++ {
			if r[i] == r[j] {
				r[j] = r[length]
				r = r[:length]
				length--
				j--
			}
		}
	}
	return string(r)
}

// Function RemoveDuplicatesMap will remove duplicates by moving unique values to new positions
// after the entire slice is traversed then the function will return a resliced string.
func RemoveDuplicatesMap(s string) string {
	r := []rune(s)
	m := map[rune]bool{}
	for _, val := range r {
		if _, seen := m[val]; !seen {
			r[len(m)] = val
			m[val] = true
		}
	}
	return string(r[:len(m)])
}

// Function RemoveDuplicatesNewMap creates an empty rune slice rather than a "cloned" version of
// input string s. This way there is no "mutation" just new unique characters added to the slice.
func RemoveDuplicatesNewMap(s string) string {
	r := []rune{}
	m := map[rune]bool{}
	for _, val := range s {
		if _, seen := m[val]; !seen {
			r = append(r, val)
			m[val] = true
		}
	}
	return string(r)
}
