package main

// func SimpleLoop(b []byte) ([]byte, error) {
// 	for len(b) > 0 {
// 		switch b[0] {
// 		case '\\': // Escape sequence
// 			b = b[1:]
// 			if len(b) < 1 {
// 				return b, ErrUnexpectedEOF
// 			}
// 			switch b[0] {
// 			case '"', '\\', '/', 'b', 'f', 'n', 'r', 't':
// 				b = b[1:]
// 				continue
// 			case 'u':
// 				if len(b) < 5 {
// 					return nil, ErrUnexpectedEOF
// 				}
// 				if err := expectHex(b[1]); err != nil {
// 					return b[1:], err
// 				}
// 				if err := expectHex(b[2]); err != nil {
// 					return b[2:], err
// 				}
// 				if err := expectHex(b[3]); err != nil {
// 					return b[3:], err
// 				}
// 				if err := expectHex(b[4]); err != nil {
// 					return b[4:], err
// 				}
// 				b = b[1:]
// 			}
// 			if b[0] < 0x20 {
// 				// Illegal control character
// 				return b, ErrIllegalControlChar
// 			}
// 		case '"': // End of string
// 			return b[1:], nil
// 		default:
// 			if b[0] < 0x20 {
// 				return b, ErrIllegalControlChar
// 			}
// 			b = b[1:]
// 		}
// 	}
// 	return b, ErrUnexpectedEOF
// }

// func expectHex(b byte) error {
// 	switch b {
// 	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
// 		'a', 'b', 'c', 'd', 'e', 'f', 'A', 'B', 'C', 'D', 'E', 'F':
// 		return nil
// 	}
// 	if b < 0x20 {
// 		return ErrIllegalControlChar
// 	}
// 	return ErrIllegalEscapeSeq
// }
