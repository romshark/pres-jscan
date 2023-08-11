package main

func SpIndex(s []byte) []byte {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ', '\r', '\n', '\t':
			continue
		}
		// Not a space character
		return s[i:]
	}
	return nil
}
