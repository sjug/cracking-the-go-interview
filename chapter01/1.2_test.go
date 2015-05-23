package onetwo

import "testing"

var cases = []struct {
	in   string
	want string
}{
	{
		"asdfghjkl",
		"lkjhgfdsa",
	},
	{
		"qwertyuiop",
		"poiuytrewq",
	},
	{
		"1234567890",
		"0987654321",
	},
	{
		"The quick brown fox jumped over the frog.",
		".gorf eht revo depmuj xof nworb kciuq ehT",
	},
}

func TestIsUnique(t *testing.T) {
	for _, tc := range cases {
		got := Reverse(tc.in)
		if got != tc.want {
			t.Fatalf("Reverse(%s) = %s want %s", tc.in, got, tc.want)
		}
	}
	t.Log("Tested", len(cases), "cases.")
}

func BenchmarkIsUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range cases {
			Reverse(tc.in)
		}
	}
}
