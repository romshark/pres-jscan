package main

func SpUnrol16Switch(s []byte) []byte {
	for ; len(s) > 15; s = s[16:] {
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
		if isNotSpace(s[8]) {
			return s[8:]
		}
		if isNotSpace(s[9]) {
			return s[9:]
		}
		if isNotSpace(s[10]) {
			return s[10:]
		}
		if isNotSpace(s[11]) {
			return s[11:]
		}
		if isNotSpace(s[12]) {
			return s[12:]
		}
		if isNotSpace(s[13]) {
			return s[13:]
		}
		if isNotSpace(s[14]) {
			return s[14:]
		}
		if isNotSpace(s[15]) {
			return s[15:]
		}
	}
	for ; len(s) > 0; s = s[1:] {
		if isNotSpace(s[0]) {
			return s
		}
	}
	return s
}
