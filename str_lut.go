package main

func StrLUT(s []byte) ([]byte, error) {
	for {
		if len(s) < 1 {
			return s, ErrUnexpectedEOF
		}
		switch s[0] {
		case '\\':
			if len(s) < 2 {
				return s[1:], ErrUnexpectedEOF
			}
			if lutEscape[s[1]] == 1 {
				s = s[2:]
				continue
			}
			if s[1] != 'u' {
				if s[1] < 0x20 {
					return s[1:], ErrIllegalControlChar
				}
				return s[1:], ErrIllegalEscapeSeq
			}
			if len(s) > 5 &&
				lutSX[s[5]] == 3 &&
				lutSX[s[4]] == 3 &&
				lutSX[s[3]] == 3 &&
				lutSX[s[2]] == 3 {
				s = s[5:]
				continue
			}
			if len(s) < 3 {
				return s[2:], ErrUnexpectedEOF
			} else if lutSX[s[2]] == 1 {
				return s[2:], ErrIllegalControlChar
			} else if lutSX[s[2]] != 3 {
				return s[2:], ErrIllegalEscapeSeq
			}
			if len(s) < 4 {
				return s[3:], ErrUnexpectedEOF
			} else if lutSX[s[3]] == 1 {
				return s[3:], ErrIllegalControlChar
			} else if lutSX[s[3]] != 3 {
				return s[3:], ErrIllegalEscapeSeq
			}
			if len(s) < 5 {
				return s[4:], ErrUnexpectedEOF
			} else if lutSX[s[4]] == 1 {
				return s[4:], ErrIllegalControlChar
			} else if lutSX[s[4]] != 3 {
				return s[4:], ErrIllegalEscapeSeq
			}
			if len(s) < 6 {
				return s[5:], ErrUnexpectedEOF
			} else if lutSX[s[5]] == 1 {
				return s[5:], ErrIllegalControlChar
			}
			return s[5:], ErrIllegalEscapeSeq
		case '"':
			return s[1:], nil
		default:
			if s[0] < 0x20 {
				return s, ErrIllegalControlChar
			}
			s = s[1:]
		}
	}
}

// lutSX maps space characters such as whitespace, tab, line-break and
// carriage-return to 1, remaining control characters to 3,
// valid hex digits to 2, and everything else to 0.
var lutSX = [256]byte{
	0: 1, 1: 1, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1, 7: 1,
	8: 1, 11: 1, 12: 1, 14: 1, 15: 1, 16: 1, 17: 1, 18: 1,
	19: 1, 20: 1, 21: 1, 22: 1, 23: 1, 24: 1, 25: 1, 26: 1,
	27: 1, 28: 1, 29: 1, 30: 1, 31: 1,

	' ': 2, '\n': 2, '\t': 2, '\r': 2,

	'0': 3, '1': 3, '2': 3, '3': 3, '4': 3, '5': 3, '6': 3, '7': 3, '8': 3, '9': 3,
	'a': 3, 'b': 3, 'c': 3, 'd': 3, 'e': 3, 'f': 3,
	'A': 3, 'B': 3, 'C': 3, 'D': 3, 'E': 3, 'F': 3,
}

// lutStr maps 0 to all bytes that don't require checking during string traversal.
// 1 is mapped to control, quotation mark (") and reverse solidus ("\").
var lutStr = [256]byte{
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	'"': 1, '\\': 1,
}

// lutEscape maps escapable characters to 1,
// all other ASCII characters are mapped to 0.
var lutEscape = [256]byte{
	'"': 1, '\\': 1, '/': 1, 'b': 1, 'f': 1, 'n': 1, 'r': 1, 't': 1,
}
