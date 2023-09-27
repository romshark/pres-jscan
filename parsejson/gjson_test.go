package parsejson_test

import (
	"errors"

	"github.com/tidwall/gjson"
)

func NewGJSON() Traverser { return GJSON{} }

type GJSON struct{}

var gjsonTypeMap = [...]JSONValueType{
	gjson.String: JSONValueTypeString,
	gjson.Number: JSONValueTypeNumber,
	gjson.True:   JSONValueTypeBoolean,
	gjson.False:  JSONValueTypeBoolean,
	gjson.Null:   JSONValueTypeNull,
}

func (GJSON) Traverse(
	input []byte, onVar func(name []byte, t JSONValueType),
) error {
	if !gjson.ValidBytes(input) {
		return errors.New("invalid json")
	}
	r := gjson.ParseBytes(input)
	if !r.IsObject() {
		return errors.New("expected object")
	}
	r.ForEach(func(key, value gjson.Result) bool {
		t := JSONValueTypeNull
		switch {
		case value.IsObject():
			t = JSONValueTypeObject
		case value.IsArray():
			t = JSONValueTypeArray
		default:
			t = gjsonTypeMap[value.Type]
		}
		onVar(unsafeStrToBytes(key.Str), t)
		return true
	})
	return nil
}
