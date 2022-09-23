package shift

import (
	"math/rand"
	"testing"
	"time"
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

func TestShiftEmpty1(t *testing.T) {
	s := ""
	shift := [][2]int{{0, 1}, {1, 2}}
	want := ""
	result := Shift(s, shift)
	if result != want {
		t.Fatalf(`Shift("%q, %v") = %q, want %q`, s, shift, result, want)
	}
}

func TestShiftEmpty2(t *testing.T) {
	s := "0"
	shift := [][2]int{}
	want := "0"
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

func reverseShift(shift [][2]int) [][2]int {
	rev := [][2]int{}
	for i := len(shift) - 1; i >= 0; i-- {
		sh := shift[i]
		var d int
		if sh[0] == 0 {
			d = 1
		} else {
			d = 0
		}
		rev = append(rev, [2]int{d, sh[1]})
	}
	return rev
}

func TestReverseShift(t *testing.T) {
	s := "abcdefg"
	shift := [][2]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}}
	result := Shift(s, shift)
	rev := reverseShift(shift)
	r := Shift(result, rev)
	if r != s {
		t.Fatalf(`Shift(Shift(%q, %v), %v) = %q, want %q`, s, shift, rev, r, s)
	}
}

func FuzzShift(f *testing.F) {
	rand.Seed(time.Now().UnixNano())
	f.Add("abcdefg", uint8(10))
	f.Add("万有引力", uint8(5))
	f.Fuzz(func(t *testing.T, s string, steps uint8) {
		steps %= 20
		shift := [][2]int{}
		for i := uint8(0); i < steps; i++ {
			d := int(rand.Uint32() % 2)
			a := int(rand.Uint32() % 30)
			shift = append(shift, [2]int{d, a})
		}
		result := Shift(s, shift)
		rev := reverseShift(shift)
		r := Shift(result, rev)
		if r != s {
			t.Errorf(`Shift(Shift(%q, %v), %v) = %q, want %q`, s, shift, rev, r, s)
		}
	})
}
