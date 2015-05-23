package onefive

import "testing"

var cases = []struct {
	in   string
	want string
}{
	{
		"Hello there!",
		"Hello%20there!",
	},
	{
		"Where is the rain in spain?",
		"Where%20is%20the%20rain%20in%20spain?",
	},
	{
		" ",
		"%20",
	},
}

func TestStringsSpaceReplace(t *testing.T) {
	for _, tc := range cases {
		got := StringsSpaceReplace(tc.in)
		if got != tc.want {
			t.Fatalf("StringsSpaceReplace(%s) = %s want %s", tc.in, got, tc.want)
		}
	}
	t.Log("Tested", len(cases), "cases.")
}

func TestSpaceReplace(t *testing.T) {
	for _, tc := range cases {
		got := SpaceReplace(tc.in)
		if got != tc.want {
			t.Fatalf("SpaceReplace(%s) = %s want %s", tc.in, got, tc.want)
		}
	}
	t.Log("Tested", len(cases), "cases.")
}

func BenchmarkStringsSpaceReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range cases {
			StringsSpaceReplace(tc.in)
		}
	}
}

func BenchmarkSpaceReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range cases {
			SpaceReplace(tc.in)
		}
	}
}
