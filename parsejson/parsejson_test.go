package parsejson_test

import (
	_ "embed"
	"encoding/json"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

type JSONValueType int8

const (
	_ JSONValueType = iota
	JSONValueTypeObject
	JSONValueTypeArray
	JSONValueTypeString
	JSONValueTypeNumber
	JSONValueTypeBoolean
	JSONValueTypeNull
)

// Traverse calls onVariable for every variable encountered in input.
type Traverser interface {
	Traverse(input []byte, onVariable func(name []byte, t JSONValueType)) error
}

var implementations = []struct {
	Name string
	Make func() Traverser
}{
	{Name: "encoding_json_______", Make: NewEncodingJSON},
	{Name: "encoding_json_unsafe", Make: NewEncodingJSONUnsafe},
	{Name: "encoding_json_opt___", Make: NewEncodingJSONOptimized},
	{Name: "bytedance_sonic_____", Make: NewBytedanceSonic},
	{Name: "jsoniter____________", Make: NewJsoniter},
	{Name: "jsoniter_opt________", Make: NewJsoniterOptimized},
	{Name: "jsoniter_opt_unsafe_", Make: NewJsoniterOptimizedUnsafe},
	{Name: "gofaster_jx_________", Make: NewGofasterJX},
	{Name: "jscan_______________", Make: NewJscan},

	// github.com/valyala/fastjson is disqualified due to violation of the JSON specification.
	// Specification: https://datatracker.ietf.org/doc/html/rfc8259
	// Code example: https://go.dev/play/p/CwIgaLHJLJp
	// This is a known issue: https://github.com/valyala/fastjson/issues/88
	{Name: "fastjson____________", Make: NewFastjson},
}

func Test(t *testing.T) {
	type Expectation struct {
		Name string
		Type JSONValueType
	}
	for _, td := range []struct {
		Name   string
		Input  string
		Expect []Expectation
	}{
		{
			Name:   "empty_object",
			Input:  `{}`,
			Expect: nil,
		},
		{
			Name:   "value_object_small",
			Input:  `{"object":{"foo":"bar"}}`,
			Expect: []Expectation{{"object", JSONValueTypeObject}},
		},
		{
			Name:  "value_object",
			Input: fileJSONObject,
			Expect: []Expectation{
				{"object", JSONValueTypeObject},
				{"array2D", JSONValueTypeArray},
				{"string", JSONValueTypeString},
				{"number", JSONValueTypeNumber},
				{"true", JSONValueTypeBoolean},
				{"false", JSONValueTypeBoolean},
				{"null", JSONValueTypeNull},
			},
		},
		{
			Name:   "value_object_empty",
			Input:  `{"object":{}}`,
			Expect: []Expectation{{"object", JSONValueTypeObject}},
		},
		{
			Name:   "value_array",
			Input:  `{"array":["foo",42]}`,
			Expect: []Expectation{{"array", JSONValueTypeArray}},
		},
		{
			Name:   "value_array_empty",
			Input:  `{"array":[]}`,
			Expect: []Expectation{{"array", JSONValueTypeArray}},
		},
		{
			Name:   "value_string",
			Input:  `{"string":"text"}`,
			Expect: []Expectation{{"string", JSONValueTypeString}},
		},
		{
			Name:   "value_string_empty",
			Input:  `{"string":""}`,
			Expect: []Expectation{{"string", JSONValueTypeString}},
		},
		{
			Name:   "value_number",
			Input:  `{"number":42}`,
			Expect: []Expectation{{"number", JSONValueTypeNumber}},
		},
		{
			Name:   "value_number_float",
			Input:  `{"float":3.1415}`,
			Expect: []Expectation{{"float", JSONValueTypeNumber}},
		},
		{
			Name:   "value_boolean_true",
			Input:  `{"bool":true}`,
			Expect: []Expectation{{"bool", JSONValueTypeBoolean}},
		},
		{
			Name:   "value_boolean_false",
			Input:  `{"bool":false}`,
			Expect: []Expectation{{"bool", JSONValueTypeBoolean}},
		},
	} {
		t.Run(td.Name, func(t *testing.T) {
			if !json.Valid([]byte(td.Input)) {
				t.Fatal("invalid json input")
			}
			for _, impl := range implementations {
				itr := impl.Make()
				t.Run(impl.Name, func(t *testing.T) {
					var actual []Expectation
					err := itr.Traverse(
						[]byte(td.Input),
						func(name []byte, t JSONValueType) {
							actual = append(actual, Expectation{
								Name: string(name),
								Type: t,
							})
						},
					)
					require.NoError(t, err)
					require.Len(t, actual, len(td.Expect))
					for _, expected := range td.Expect {
						require.Contains(t, actual, expected)
					}
				})
			}
		})
	}
}

var (
	GN []byte
	GT JSONValueType
)

func BenchmarkObject(b *testing.B) { benchmark(b, fileJSONObject) }

func benchmark(b *testing.B, input string) {
	in := []byte(input)
	if !json.Valid(in) {
		b.Fatal("invalid json input")
	}
	for _, impl := range implementations {
		itr := impl.Make()
		b.Run(impl.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				if err := itr.Traverse(in, func(name []byte, t JSONValueType) {
					GN, GT = name, t
				}); err != nil {
					panic(err)
				}
			}
		})
	}
}

func unsafeStrToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

//go:embed testdata/object.json
var fileJSONObject string
