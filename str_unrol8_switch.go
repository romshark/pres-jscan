package main

func StrUnrol8Switch(s []byte) ([]byte, error) {
	for {
		for ; len(s) > 7; s = s[8:] {
			if needsCheck(s[0]) {
				goto CHECK_STRING_CHARACTER
			}
			if needsCheck(s[1]) {
				s = s[1:]
				goto CHECK_STRING_CHARACTER
			}
			if needsCheck(s[2]) {
				s = s[2:]
				goto CHECK_STRING_CHARACTER
			}
			if needsCheck(s[3]) {
				s = s[3:]
				goto CHECK_STRING_CHARACTER
			}
			if needsCheck(s[4]) {
				s = s[4:]
				goto CHECK_STRING_CHARACTER
			}
			if needsCheck(s[5]) {
				s = s[5:]
				goto CHECK_STRING_CHARACTER
			}
			if needsCheck(s[6]) {
				s = s[6:]
				goto CHECK_STRING_CHARACTER
			}
			if needsCheck(s[7]) {
				s = s[7:]
				goto CHECK_STRING_CHARACTER
			}
			continue
		}

	CHECK_STRING_CHARACTER:
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

func needsCheck(b byte) bool {
	switch b {
	case 0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7,
		0x8, 0x9, 0xA, 0xB, 0xC, 0xD, 0xE, 0xF,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16,
		0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D,
		0x1E, 0x1F, 0x20,
		'\\', '"':
		return true
	}
	return false
}
