package pm1

func IsContain(r rune, runes []rune) bool {
	for _, rr := range runes {
		if rr == r {
			return true
		}
	}
	return false
}

func IsWithInRange(s string, ran []rune) bool {
	for _, r := range []rune(s) {
		if !IsContain(r, ran) {
			return false
		}
	}
	return true
}
