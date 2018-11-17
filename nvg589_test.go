package nvg589

import (
	"testing"
)

func TestConnect(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"192.168.1.254", "http://192.168.1.254/cgi-bin/broadbandstatistics.ha"},
	}

	for _, c := range cases {
		got := Connect(c.in)
		if string(got) != c.want {
			t.Errorf("Connect(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
