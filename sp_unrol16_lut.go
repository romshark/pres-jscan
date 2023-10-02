package main

func SpUnrol16LUT(s []byte) []byte {
	for ; len(s) > 15; s = s[16:] {
		if lutSpUnrol16[s[0]] != 1 {
			return s
		}
		if lutSpUnrol16[s[1]] != 1 {
			return s[1:]
		}
		if lutSpUnrol16[s[2]] != 1 {
			return s[2:]
		}
		if lutSpUnrol16[s[3]] != 1 {
			return s[3:]
		}
		if lutSpUnrol16[s[4]] != 1 {
			return s[4:]
		}
		if lutSpUnrol16[s[5]] != 1 {
			return s[5:]
		}
		if lutSpUnrol16[s[6]] != 1 {
			return s[6:]
		}
		if lutSpUnrol16[s[7]] != 1 {
			return s[7:]
		}
		if lutSpUnrol16[s[8]] != 1 {
			return s[8:]
		}
		if lutSpUnrol16[s[9]] != 1 {
			return s[9:]
		}
		if lutSpUnrol16[s[10]] != 1 {
			return s[10:]
		}
		if lutSpUnrol16[s[11]] != 1 {
			return s[11:]
		}
		if lutSpUnrol16[s[12]] != 1 {
			return s[12:]
		}
		if lutSpUnrol16[s[13]] != 1 {
			return s[13:]
		}
		if lutSpUnrol16[s[14]] != 1 {
			return s[14:]
		}
		if lutSpUnrol16[s[15]] != 1 {
			return s[15:]
		}
	}
	for ; len(s) > 0; s = s[1:] {
		if lutSpUnrol16[s[0]] != 1 {
			return s
		}
	}
	return s
}

// lutSpUnrol16 maps space characters such as whitespace, tab, line-break and
// carriage-return to 1 and all other ASCII characters to 0.
var lutSpUnrol16 = [256]byte{' ': 1, '\n': 1, '\t': 1, '\r': 1}
