package main

func StrNaive(s []byte) ([]byte, error) {
	for len(s) > 0 {
		switch s[0] {
		case '\\':
			if len(s) < 2 {
				return s[1:], ErrUnexpectedEOF
			}
			switch s[1] {
			case '"', '\\', '/', 'b', 'f', 'n', 'r', 't':
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
				isHex(s[5]) &&
				isHex(s[4]) &&
				isHex(s[3]) &&
				isHex(s[2]) {
				s = s[5:]
				continue
			}
			if len(s) < 3 {
				return s[2:], ErrUnexpectedEOF
			} else if s[2] < 0x20 {
				return s[2:], ErrIllegalControlChar
			} else if !isHex(s[2]) {
				return s[2:], ErrIllegalEscapeSeq
			}
			if len(s) < 4 {
				return s[3:], ErrUnexpectedEOF
			} else if s[3] < 0x20 {
				return s[3:], ErrIllegalControlChar
			} else if !isHex(s[3]) {
				return s[3:], ErrIllegalEscapeSeq
			}
			if len(s) < 5 {
				return s[4:], ErrUnexpectedEOF
			} else if s[4] < 0x20 {
				return s[4:], ErrIllegalControlChar
			} else if !isHex(s[4]) {
				return s[4:], ErrIllegalEscapeSeq
			}
			if len(s) < 6 {
				return s[5:], ErrUnexpectedEOF
			} else if s[5] < 0x20 {
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
	return s, ErrUnexpectedEOF
}

func isHex(b byte) bool {
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'A', 'B', 'C', 'D', 'E', 'F':
		return true
	}
	return false
}
