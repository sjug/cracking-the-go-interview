package onethree

import "testing"

var cases = []struct {
	in   string
	want string
}{
	{
		"aabbcdee",
		"abcde",
	},
	{
		"qqwweertyeuuiop",
		"qwertyuiop",
	},
}

/*
func TestRemoveDuplicates(t *testing.T) {
	for _, tc := range cases {
		got := RemoveDuplicates(tc.in)
		if got != tc.want {
			t.Fatalf("RemoveDuplicates(%s) = %s want %s", tc.in, got, tc.want)
		}
	}
	t.Log("Tested", len(cases), "cases.")
}
*/

func TestRemoveDuplicatesMap(t *testing.T) {
	for _, tc := range cases {
		got := RemoveDuplicatesMap(tc.in)
		if got != tc.want {
			t.Fatalf("RemoveDuplicatesMap(%s) = %s want %s", tc.in, got, tc.want)
		}
	}
	t.Log("Tested", len(cases), "cases.")
}

func TestRemoveDuplicatesNewMap(t *testing.T) {
	for _, tc := range cases {
		got := RemoveDuplicatesNewMap(tc.in)
		if got != tc.want {
			t.Fatalf("RemoveDuplicatesNewMap(%s) = %s want %s", tc.in, got, tc.want)
		}
	}
	t.Log("Tested", len(cases), "cases.")
}

func BenchmarkRemoveDuplicatesMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range cases {
			RemoveDuplicatesMap(tc.in)
		}
	}
}

func BenchmarkRemoveDuplicatesNewMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range cases {
			RemoveDuplicatesNewMap(tc.in)
		}
	}
}
