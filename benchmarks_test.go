package main

import (
	"errors"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

var implementations = []struct {
	Name string
	Fn   func([]byte) ([]byte, error)
}{
	{Name: "unrolled_lut", Fn: UnrolledLUT},
}

type T struct {
	Input      string
	ExpectTail string
	// CheckErr is nil when no error is expected
	CheckErr func(*testing.T, error)
}

var tests = map[string]T{
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
	for _, testName := range keysSorted(tests) {
		td := tests[testName]
		t.Run(testName, func(t *testing.T) {
			for _, impl := range implementations {
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

func Benchmark(b *testing.B) {
	for _, testName := range benchmarks {
		b.Run(testName, func(b *testing.B) {
			require.Contains(b, tests, testName)
			td := tests[testName]
			input := []byte(td.Input)

			for _, impl := range implementations {
				b.Run(impl.Name, func(b *testing.B) {
					for n := 0; n < b.N; n++ {
						GTAIL, GERR = impl.Fn(input)
					}
				})
			}
		})
	}
}

func errorIs(expect error) func(t *testing.T, err error) {
	return func(t *testing.T, err error) {
		require.Error(t, err)
		require.True(t, errors.Is(err, expect),
			"actual:   %#v;\nexpected: %#v", err, expect)
	}
}

func MustReadFile(path string) []byte {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return b
}

func keysSorted[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
