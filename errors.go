package main

import "errors"

var (
	ErrIllegalControlChar = errors.New("illegal control character")
	ErrIllegalEscapeSeq   = errors.New("illegal escape sequence")
	ErrUnexpectedEOF      = errors.New("unexpected EOF")
)
