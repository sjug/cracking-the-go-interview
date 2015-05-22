package oneone

import "testing"

var cases = []struct {
	in   string
	want bool
}{
	{
		"asdfghjkl",
		true,
	},
	{
		"qwertyuiopq",
		false,
	},
}

func TestIsUnique(t *testing.T) {
	for _, tc := range cases {
		got := IsUnique(tc.in)
		if got != tc.want {
			t.Fatalf("IsUnique(%s) = %t want %t")
		}
	}
	t.Log("Tested", len(cases), "cases.")
}

func BenchmarkIsUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range cases {
			IsUnique(tc.in)
		}
	}
}
