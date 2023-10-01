package parsejson_test

import (
	"encoding/json"
)

func NewEncodingJSONOptimized() GraphQLVariablesTraverser { return EncodingJSONOptimized{} }

type EncodingJSONOptimized struct{}

var encodingJSONOptCharMap = [256]JSONValueType{
	'{': JSONValueTypeObject,
	'[': JSONValueTypeArray,
	'"': JSONValueTypeString,
	'0': JSONValueTypeNumber,
	'1': JSONValueTypeNumber,
	'2': JSONValueTypeNumber,
	'3': JSONValueTypeNumber,
	'4': JSONValueTypeNumber,
	'5': JSONValueTypeNumber,
	'6': JSONValueTypeNumber,
	'7': JSONValueTypeNumber,
	'8': JSONValueTypeNumber,
	'9': JSONValueTypeNumber,
	'-': JSONValueTypeNumber,
	't': JSONValueTypeBoolean,
	'f': JSONValueTypeBoolean,
	'n': JSONValueTypeNull,
}

func (EncodingJSONOptimized) TraverseJSON(
	input []byte, onVar func(name []byte, t JSONValueType),
) error {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(input, &m); err != nil {
		return err
	}
	for k, v := range m {
		onVar(unsafeStrToBytes(k), encodingJSONOptCharMap[v[0]])
	}
	return nil
}
