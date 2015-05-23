package onefour

import "testing"

var cases = []struct {
	in1  string
	in2  string
	want bool
}{
	{
		"george",
		"eggore",
		true,
	},
	{
		"truk",
		"kurt",
		true,
	},
	{
		"elvis",
		"lives",
		true,
	},
	{
		"silent",
		"listen",
		true,
	},
	{
		"frog",
		"dirt",
		false,
	},
}

func TestAnagram(t *testing.T) {
	for _, tc := range cases {
		got := Anagram(tc.in1, tc.in2)
		if got != tc.want {
			t.Fatalf("Anagram(%s, %s) = %t want %t", tc.in1, tc.in2, got, tc.want)
		}
	}
	t.Log("Tested", len(cases), "cases.")
}

func TestAnagramReflect(t *testing.T) {
	for _, tc := range cases {
		got := AnagramReflect(tc.in1, tc.in2)
		if got != tc.want {
			t.Fatalf("AnagramReflect(%s, %s) = %t want %t", tc.in1, tc.in2, got, tc.want)
		}
	}
	t.Log("Tested", len(cases), "cases.")
}

func TestAnagramSort(t *testing.T) {
	for _, tc := range cases {
		got := AnagramSort(tc.in1, tc.in2)
		if got != tc.want {
			t.Fatalf("AnagramSort(%s, %s) = %t want %t", tc.in1, tc.in2, got, tc.want)
		}
	}
	t.Log("Tested", len(cases), "cases.")
}

func BenchmarkAnagram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range cases {
			Anagram(tc.in1, tc.in2)
		}
	}
}

func BenchmarkAnagramReflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range cases {
			AnagramReflect(tc.in1, tc.in2)
		}
	}
}

func BenchmarkAnagramSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range cases {
			AnagramSort(tc.in1, tc.in2)
		}
	}
}
