package oneeight

import "testing"

var cases = []struct {
	in1  string
	in2  string
	want bool
}{
	{
		"george",
		"rgegeo",
		true,
	},
	{
		"waterbottle",
		"erbottlewat",
		true,
	},
	{
		"apple",
		"pleap",
		true,
	},
	{
		"apple",
		"ppale",
		false,
	},
	{
		"frog",
		"dirt",
		false,
	},
	{
		"",
		"jeepers",
		false,
	},
	{
		"swagger",
		"api",
		false,
	},
}

func TestIsRotation(t *testing.T) {
	for _, tc := range cases {
		got := IsRotation(tc.in1, tc.in2)
		if got != tc.want {
			t.Fatalf("IsRotation(%s, %s) = %t want %t", tc.in1, tc.in2, got, tc.want)
		}
	}
	t.Log("Tested", len(cases), "cases.")
}

func BenchmarkIsRotation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range cases {
			IsRotation(tc.in1, tc.in2)
		}
	}
}
