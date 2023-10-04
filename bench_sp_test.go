package main

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

var implementationsSp = []struct {
	Name string
	Fn   func([]byte) []byte
}{
	{Name: "sp_index_________", Fn: SpIndex},
	{Name: "sp_unrol8_switch_", Fn: SpUnrol8Switch},
	{Name: "sp_unrol8_lut____", Fn: SpUnrol8LUT},
	{Name: "sp_unrol16_switch", Fn: SpUnrol16Switch},
	{Name: "sp_unrol16_lut___", Fn: SpUnrol16LUT},
}

const spTestTail32 = `{"object-key1" : "object-value"}`

type TSp struct {
	Input      string
	ExpectTail string
}

var testsSp = map[string]T{
	"eof": {
		Input: ``,
	},
	"no_spaces_tail1": {
		Input:      `{`,
		ExpectTail: `{`,
	},
	"no_spaces_tail32": {
		Input:      spTestTail32,
		ExpectTail: spTestTail32,
	},
	"whitespace1_tail1": {
		Input:      ` {`,
		ExpectTail: `{`,
	},
	"whitespace1_tail32": {
		Input:      "  " + spTestTail32,
		ExpectTail: spTestTail32,
	},
	"tab1_tail1": {
		Input:      "\t{",
		ExpectTail: `{`,
	},
	"linefeed1_tail1": {
		Input:      "\n{",
		ExpectTail: `{`,
	},
	"carriage_return_tail1": {
		Input:      "\r{",
		ExpectTail: `{`,
	},
	"sequence7_tail1": {
		Input:      repeat(" ", 7) + `{`,
		ExpectTail: `{`,
	},
	"sequence7_tail32": {
		Input:      repeat(" ", 7) + spTestTail32,
		ExpectTail: spTestTail32,
	},
	"sequence8_tail1": {
		Input:      "  \r\n\t\t\t\t{",
		ExpectTail: `{`,
	},
	"sequence16_tail1": {
		Input:      "      \r\n\r\n\r\n\t\t\t\t{",
		ExpectTail: `{`,
	},
	"sequence40_tail32": {
		Input:      repeat(" \t", 20) + spTestTail32,
		ExpectTail: spTestTail32,
	},
	"sequence130_tail32": {
		Input:      repeat(" \t", 65) + spTestTail32,
		ExpectTail: spTestTail32,
	},
}

func TestSp(t *testing.T) {
	for _, testName := range keysSorted(testsSp) {
		td := testsSp[testName]
		t.Run(testName, func(t *testing.T) {
			for _, impl := range implementationsSp {
				t.Run(impl.Name, func(t *testing.T) {
					out := impl.Fn([]byte(td.Input))
					require.Equal(t, td.ExpectTail, string(out))
				})
			}
		})
	}
}

// benchmarksSp refers to keys of tests
var benchmarksSp = []string{
	"tab1_tail1",
	"no_spaces_tail32",
	"whitespace1_tail32",
	"sequence40_tail32",
	"sequence130_tail32",
}

func BenchmarkSp(b *testing.B) {
	for _, testName := range benchmarksSp {
		b.Run(testName, func(b *testing.B) {
			require.Contains(b, testsSp, testName)
			td := testsSp[testName]
			input := []byte(td.Input)
			var resTail []byte

			for _, impl := range implementationsSp {
				b.Run(impl.Name, func(b *testing.B) {
					for n := 0; n < b.N; n++ {
						resTail = impl.Fn(input)
					}
				})
			}

			runtime.KeepAlive(resTail)
		})
	}
}
