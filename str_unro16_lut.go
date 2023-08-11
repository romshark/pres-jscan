package main

func StrUnrol16LUT(s []byte) ([]byte, error) {
	for {
		for ; len(s) > 15; s = s[16:] {
			if lutStr[s[0]] != 0 {
				goto CHECK_STRING_CHARACTER
			}
			if lutStr[s[1]] != 0 {
				s = s[1:]
				goto CHECK_STRING_CHARACTER
			}
			if lutStr[s[2]] != 0 {
				s = s[2:]
				goto CHECK_STRING_CHARACTER
			}
			if lutStr[s[3]] != 0 {
				s = s[3:]
				goto CHECK_STRING_CHARACTER
			}
			if lutStr[s[4]] != 0 {
				s = s[4:]
				goto CHECK_STRING_CHARACTER
			}
			if lutStr[s[5]] != 0 {
				s = s[5:]
				goto CHECK_STRING_CHARACTER
			}
			if lutStr[s[6]] != 0 {
				s = s[6:]
				goto CHECK_STRING_CHARACTER
			}
			if lutStr[s[7]] != 0 {
				s = s[7:]
				goto CHECK_STRING_CHARACTER
			}
			if lutStr[s[8]] != 0 {
				s = s[8:]
				goto CHECK_STRING_CHARACTER
			}
			if lutStr[s[9]] != 0 {
				s = s[9:]
				goto CHECK_STRING_CHARACTER
			}
			if lutStr[s[10]] != 0 {
				s = s[10:]
				goto CHECK_STRING_CHARACTER
			}
			if lutStr[s[11]] != 0 {
				s = s[11:]
				goto CHECK_STRING_CHARACTER
			}
			if lutStr[s[12]] != 0 {
				s = s[12:]
				goto CHECK_STRING_CHARACTER
			}
			if lutStr[s[13]] != 0 {
				s = s[13:]
				goto CHECK_STRING_CHARACTER
			}
			if lutStr[s[14]] != 0 {
				s = s[14:]
				goto CHECK_STRING_CHARACTER
			}
			if lutStr[s[15]] != 0 {
				s = s[15:]
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
