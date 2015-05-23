package onefour

import (
	"reflect"
	"sort"
)

// Function Anagram creates two maps from the two strings then compares the two maps
func Anagram(a, b string) bool {
	m1 := map[rune]int{}
	m2 := map[rune]int{}
	if len(a) == len(b) {
		for _, valA := range a {
			if _, seen := m1[valA]; !seen {
				m1[valA]++
			}
		}
		for _, valB := range b {
			if _, seen := m2[valB]; !seen {
				m2[valB]++
			}
		}
	} else {
		return false
	}
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if w, seen := m2[k]; !seen || w != v {
			return false
		}
	}
	return true
}

// Function AnagramReflect also creates two maps from the input strings, then compares
// them using the much slower but much simpler "reflect" standard library.
func AnagramReflect(a, b string) bool {
	m1 := map[rune]int{}
	m2 := map[rune]int{}
	if len(a) == len(b) {
		for _, valA := range a {
			if _, seen := m1[valA]; !seen {
				m1[valA]++
			}
		}
		for _, valB := range b {
			if _, seen := m2[valB]; !seen {
				m2[valB]++
			}
		}
	} else {
		return false
	}
	return reflect.DeepEqual(m1, m2)
}

// In order to use the geric golang "sort" function we need to implement the
// sort interface: Len, Less, Swap.
type ByLength []rune

func (s ByLength) Len() int           { return len(s) }
func (s ByLength) Less(i, j int) bool { return s[i] < s[j] }
func (s ByLength) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Function AnagramSort converts input strings to runes, then sorts the runes and does
// a rune by rune comparison. Surprisingly this is the fastest anagram function I have.
func AnagramSort(a, b string) bool {
	r1 := []rune(a)
	r2 := []rune(b)
	if len(r1) != len(r2) {
		return false
	}
	sort.Sort(ByLength(r1))
	sort.Sort(ByLength(r2))
	for i := 0; i < len(r1); i++ {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}
