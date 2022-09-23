package shift

import (
	"testing"
)

// Ref: https://blog.csdn.net/keephide/article/details/108968360

func TestShift1(t *testing.T) {
	s := "abc"
	shift := [][2]int{{0, 1}, {1, 2}}
	want := "cab"
	result := Shift(s, shift)
	if result != want {
		t.Fatalf(`Shift("%q, %v") = %q, want %q`, s, shift, result, want)
	}
}

func TestShift2(t *testing.T) {
	s := "abcdefg"
	shift := [][2]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}}
	want := "efgabcd"
	result := Shift(s, shift)
	if result != want {
		t.Fatalf(`Shift("%q, %v") = %q, want %q`, s, shift, result, want)
	}
}

func TestShiftEmpty(t *testing.T) {
	s := ""
	shift := [][2]int{{0, 1}, {1, 2}}
	want := ""
	result := Shift(s, shift)
	if result != want {
		t.Fatalf(`Shift("%q, %v") = %q, want %q`, s, shift, result, want)
	}
}

func TestShiftOverflow(t *testing.T) {
	s := "abc"
	shift := [][2]int{{0, 4}, {1, 8}}
	want := "cab"
	result := Shift(s, shift)
	if result != want {
		t.Fatalf(`Shift("%q, %v") = %q, want %q`, s, shift, result, want)
	}
}
