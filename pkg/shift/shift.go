package shift

func Shift(s string, shift [][2]int) string {
	if s == "" || len(shift) == 0 {
		return ""
	}
	amount := 0
	for _, sh := range shift {
		if sh[0] == 0 {
			amount -= sh[1]
		} else {
			amount += sh[1]
		}
	}
	if amount < 0 {
		s = leftShift(s, -amount)
	} else if amount > 0 {
		s = rightShift(s, amount)
	}
	return s
}

func leftShift(s string, amount int) string {
	i := amount % len(s)
	return s[i:len(s)] + s[0:i]
}

func rightShift(s string, amount int) string {
	i := amount % len(s)
	return s[len(s)-i:len(s)] + s[0:len(s)-i]
}
