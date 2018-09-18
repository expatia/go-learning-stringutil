package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"!dlrow ,olleH", "Hello, world"},
		{"", ""},
	}

	for _, c := range cases {
		// got := Reverse(c.in)
		// if got != c.want {
		// 	t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		// }

		if got := Reverse(c.in); got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
