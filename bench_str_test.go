package main

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

var implementationsStr = []struct {
	Name string
	Fn   func([]byte) ([]byte, error)
}{
	{Name: "str_naive______", Fn: StrNaive},
	{Name: "str_lut________", Fn: StrLUT},
	{Name: "str_unrol8_sw__", Fn: StrUnrol8Switch},
	{Name: "str_unrol8_lut_", Fn: StrUnrol8LUT},
	{Name: "str_unrol16_lut", Fn: StrUnrol16LUT},
}

type T struct {
	Input      string
	ExpectTail string
	// CheckErr is nil when no error is expected
	CheckErr func(*testing.T, error)
}

var testsStr = map[string]T{
	"empty_string": {
		Input:      `"}`,
		ExpectTail: `}`,
	},
	"tiny_string": {
		Input:      `something"}`,
		ExpectTail: `}`,
	},
	"escaped_quotation_mark": {
		Input:      `something\""}`,
		ExpectTail: `}`,
	},
	"escaped_reverse_solidus": {
		Input:      `something\\"}`,
		ExpectTail: `}`,
	},
	"escaped_solidus": {
		Input:      `something\/"}`,
		ExpectTail: `}`,
	},
	"escaped_backspace": {
		Input:      `something\b"}`,
		ExpectTail: `}`,
	},
	"escaped_form_feed": {
		Input:      `something\f"}`,
		ExpectTail: `}`,
	},
	"escaped_line_break": {
		Input:      `something\n"}`,
		ExpectTail: `}`,
	},
	"escaped_carriage_return": {
		Input:      `something\r"}`,
		ExpectTail: `}`,
	},
	"escaped_tab": {
		Input:      `something\t"}`,
		ExpectTail: `}`,
	},
	"escaped_rune": {
		Input:      `something\uaaaa"}`,
		ExpectTail: `}`,
	},
	"testdata_wikipedia_regex": {
		Input: string(MustReadFile("testdata/wikipedia_regex.json"))[1:],
	},

	// Errors
	"err_illegal_control_char": {
		Input:      "followed by zero byte" + string(byte(0x00)) + `"`,
		ExpectTail: string(byte(0x00)) + `"`,
		CheckErr:   errorIs(ErrIllegalControlChar),
	},
	"err_illegal_control_char_in_escape_sequence_at_0": {
		Input:      "followed by zero byte \\" + string(byte(0x00)) + `"`,
		ExpectTail: string(byte(0x00)) + `"`,
		CheckErr:   errorIs(ErrIllegalControlChar),
	},
	"err_illegal_control_char_in_escape_sequence_at_1": {
		Input:      "followed by zero byte \\u" + string(byte(0x00)) + `000"`,
		ExpectTail: string(byte(0x00)) + `000"`,
		CheckErr:   errorIs(ErrIllegalControlChar),
	},
	"err_illegal_control_char_in_escape_sequence_at_2": {
		Input:      "followed by zero byte \\u0" + string(byte(0x00)) + `00"`,
		ExpectTail: string(byte(0x00)) + `00"`,
		CheckErr:   errorIs(ErrIllegalControlChar),
	},
	"err_illegal_control_char_in_escape_sequence_at_3": {
		Input:      "followed by zero byte \\u00" + string(byte(0x00)) + `0"`,
		ExpectTail: string(byte(0x00)) + `0"`,
		CheckErr:   errorIs(ErrIllegalControlChar),
	},
	"err_illegal_control_char_in_escape_sequence_at_4": {
		Input:      "followed by zero byte \\u000" + string(byte(0x00)) + `"`,
		ExpectTail: string(byte(0x00)) + `"`,
		CheckErr:   errorIs(ErrIllegalControlChar),
	},
	"err_illegal_escape_sequence_at_0": {
		Input:      "illegal escape sequence: \\x0000\"",
		ExpectTail: `x0000"`,
		CheckErr:   errorIs(ErrIllegalEscapeSeq),
	},
	"err_illegal_escape_sequence_at_1": {
		Input:      `illegal escape sequence: \u?000"`,
		ExpectTail: `?000"`,
		CheckErr:   errorIs(ErrIllegalEscapeSeq),
	},
	"err_illegal_escape_sequence_at_2": {
		Input:      `illegal escape sequence: \u0?00"`,
		ExpectTail: `?00"`,
		CheckErr:   errorIs(ErrIllegalEscapeSeq),
	},
	"err_illegal_escape_sequence_at_3": {
		Input:      `illegal escape sequence: \u00?0"`,
		ExpectTail: `?0"`,
		CheckErr:   errorIs(ErrIllegalEscapeSeq),
	},
	"err_illegal_escape_sequence_at_4": {
		Input:      `illegal escape sequence: \u000?"`,
		ExpectTail: `?"`,
		CheckErr:   errorIs(ErrIllegalEscapeSeq),
	},
	"err_unexpected_EOF_empty_input": {
		Input:      "",
		ExpectTail: "",
		CheckErr:   errorIs(ErrUnexpectedEOF),
	},
	"err_unexpected_EOF_missing_closing_quotes": {
		Input:      "no closing quotes here ->",
		ExpectTail: "",
		CheckErr:   errorIs(ErrUnexpectedEOF),
	},
	"err_unexpected_EOF_escape_sequence_at_0": {
		Input:    `illegal escape sequence: \`,
		CheckErr: errorIs(ErrUnexpectedEOF),
	},
	"err_unexpected_EOF_escape_sequence_at_1": {
		Input:    `illegal escape sequence: \u`,
		CheckErr: errorIs(ErrUnexpectedEOF),
	},
	"err_unexpected_EOF_escape_sequence_at_2": {
		Input:    `illegal escape sequence: \uf`,
		CheckErr: errorIs(ErrUnexpectedEOF),
	},
	"err_unexpected_EOF_escape_sequence_at_3": {
		Input:    `illegal escape sequence: \uff`,
		CheckErr: errorIs(ErrUnexpectedEOF),
	},
	"err_unexpected_EOF_escape_sequence_at_4": {
		Input:    `illegal escape sequence: \ufff`,
		CheckErr: errorIs(ErrUnexpectedEOF),
	},
}

func Test(t *testing.T) {
	for _, testName := range keysSorted(testsStr) {
		td := testsStr[testName]
		t.Run(testName, func(t *testing.T) {
			for _, impl := range implementationsStr {
				t.Run(impl.Name, func(t *testing.T) {
					out, err := impl.Fn([]byte(td.Input))
					if td.CheckErr != nil {
						td.CheckErr(t, err)
					} else {
						require.NoError(t, err)
					}
					require.Equal(t, td.ExpectTail, string(out))
				})
			}
		})
	}
}

// benchmarks refers to keys of tests
var benchmarks = []string{
	"testdata_wikipedia_regex",
	"tiny_string",
}

var (
	GTAIL []byte
	GERR  error
)

func BenchmarkStr(b *testing.B) {
	for _, testName := range benchmarks {
		b.Run(testName, func(b *testing.B) {
			require.Contains(b, testsStr, testName)
			td := testsStr[testName]
			input := []byte(td.Input)

			var resTail []byte
			var resErr error

			for _, impl := range implementationsStr {
				b.Run(impl.Name, func(b *testing.B) {
					for n := 0; n < b.N; n++ {
						resTail, resErr = impl.Fn(input)
					}
				})
			}

			runtime.KeepAlive(resTail)
			runtime.KeepAlive(resErr)
		})
	}
}
