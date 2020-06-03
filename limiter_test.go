package keyrate_test

import (
	"testing"
	"time"

	"github.com/collectai/keyrate"
)

func TestIntLimiter(t *testing.T) {
	t.Parallel()

	l := keyrate.NewIntLimiter(1, 1)

	a := 1
	b := 7

	cases := []struct {
		in   func() bool
		want bool
	}{
		{
			func() bool { return l.Allow(a) },
			true,
		},
		{
			func() bool { return l.Allow(b) },
			true,
		},
		{
			func() bool { return l.Allow(a) },
			false,
		},
		{
			func() bool { time.Sleep(time.Second); return l.Allow(a) },
			true,
		},
		{
			func() bool { return l.Allow(b) },
			true,
		},
	}

	for i, c := range cases {
		out := c.in()
		if out != c.want {
			t.Errorf("%d: got %t, want %t", i, out, c.want)
		}
	}
}

func TestStringLimiter(t *testing.T) {
	t.Parallel()

	l := keyrate.NewStringLimiter(1, 1)

	a := "a"
	b := "b"

	cases := []struct {
		in   func() bool
		want bool
	}{
		{
			func() bool { return l.Allow(a) },
			true,
		},
		{
			func() bool { return l.Allow(b) },
			true,
		},
		{
			func() bool { return l.Allow(a) },
			false,
		},
		{
			func() bool { time.Sleep(time.Second); return l.Allow(a) },
			true,
		},
		{
			func() bool { return l.Allow(b) },
			true,
		},
	}

	for i, c := range cases {
		out := c.in()
		if out != c.want {
			t.Errorf("%d: got %t, want %t", i, out, c.want)
		}
	}
}

func TestIntLimiterN(t *testing.T) {
	t.Parallel()

	l := keyrate.NewIntLimiter(2, 3)

	a := 6
	b := 13

	cases := []struct {
		in   func() bool
		want bool
	}{
		{
			func() bool { return l.AllowN(a, 2) },
			true,
		},
		{
			func() bool { return l.Allow(b) },
			true,
		},
		{
			func() bool { return l.AllowN(a, 2) },
			false,
		},
		{
			func() bool { return l.AllowN(b, 2) },
			true,
		},
		{
			func() bool { time.Sleep(time.Second); return l.AllowN(a, 3) },
			true,
		},
	}

	for i, c := range cases {
		out := c.in()
		if out != c.want {
			t.Errorf("%d: got %t, want %t", i, out, c.want)
		}
	}
}

func TestStringLimiterN(t *testing.T) {
	t.Parallel()

	l := keyrate.NewStringLimiter(2, 3)

	a := "a"
	b := "b"

	cases := []struct {
		in   func() bool
		want bool
	}{
		{
			func() bool { return l.AllowN(a, 2) },
			true,
		},
		{
			func() bool { return l.Allow(b) },
			true,
		},
		{
			func() bool { return l.AllowN(a, 2) },
			false,
		},
		{
			func() bool { return l.AllowN(b, 2) },
			true,
		},
		{
			func() bool { time.Sleep(time.Second); return l.AllowN(a, 3) },
			true,
		},
	}

	for i, c := range cases {
		out := c.in()
		if out != c.want {
			t.Errorf("%d: got %t, want %t", i, out, c.want)
		}
	}
}
