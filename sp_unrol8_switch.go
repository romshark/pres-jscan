package main

func SpUnrol8Switch(s []byte) []byte {
	for ; len(s) > 7; s = s[8:] {
		if isNotSpace(s[0]) {
			return s
		}
		if isNotSpace(s[1]) {
			return s[1:]
		}
		if isNotSpace(s[2]) {
			return s[2:]
		}
		if isNotSpace(s[3]) {
			return s[3:]
		}
		if isNotSpace(s[4]) {
			return s[4:]
		}
		if isNotSpace(s[5]) {
			return s[5:]
		}
		if isNotSpace(s[6]) {
			return s[6:]
		}
		if isNotSpace(s[7]) {
			return s[7:]
		}
	}
	for ; len(s) > 0; s = s[1:] {
		if isNotSpace(s[0]) {
			return s
		}
	}
	return s
}

func isNotSpace(b byte) bool {
	switch b {
	case ' ', '\r', '\n', '\t':
		return false
	}
	return true
}
