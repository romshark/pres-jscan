package main

func SpUnrol8LUT(s []byte) []byte {
	for ; len(s) > 7; s = s[8:] {
		if lutSpUnrol8[s[0]] != 1 {
			return s
		}
		if lutSpUnrol8[s[1]] != 1 {
			return s[1:]
		}
		if lutSpUnrol8[s[2]] != 1 {
			return s[2:]
		}
		if lutSpUnrol8[s[3]] != 1 {
			return s[3:]
		}
		if lutSpUnrol8[s[4]] != 1 {
			return s[4:]
		}
		if lutSpUnrol8[s[5]] != 1 {
			return s[5:]
		}
		if lutSpUnrol8[s[6]] != 1 {
			return s[6:]
		}
		if lutSpUnrol8[s[7]] != 1 {
			return s[7:]
		}
	}
	for ; len(s) > 0; s = s[1:] {
		if lutSpUnrol8[s[0]] != 1 {
			return s
		}
	}
	return s
}

// lutSpUnrol8 maps space characters such as whitespace, tab, line-break and
// carriage-return to 1 and all other ASCII characters to 0.
var lutSpUnrol8 = [256]byte{' ': 1, '\n': 1, '\t': 1, '\r': 1}
