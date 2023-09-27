package parsejson_test

import (
	"errors"

	jsoniter "github.com/json-iterator/go"
)

func NewJsoniter() Traverser {
	return Jsoniter{
		Iterator: jsoniter.NewIterator(jsoniter.ConfigFastest),
	}
}

type Jsoniter struct{ *jsoniter.Iterator }

var jsoniterTypeMap = [...]JSONValueType{
	0:                    0, // Zero value
	jsoniter.ObjectValue: JSONValueTypeObject,
	jsoniter.ArrayValue:  JSONValueTypeArray,
	jsoniter.StringValue: JSONValueTypeString,
	jsoniter.NumberValue: JSONValueTypeNumber,
	jsoniter.BoolValue:   JSONValueTypeBoolean,
	jsoniter.NilValue:    JSONValueTypeNull,
}

func (itr Jsoniter) Traverse(
	input []byte, onVar func(name []byte, t JSONValueType),
) (err error) {
	itr.ResetBytes(input)
	if itr.WhatIsNext() != jsoniter.ObjectValue {
		return errors.New("expected object")
	}
	itr.ReadObjectCB(func(itr *jsoniter.Iterator, s string) bool {
		n := itr.WhatIsNext()
		if n == jsoniter.InvalidValue {
			err = errors.New("invalid value")
			return false
		}
		onVar([]byte(s), jsoniterTypeMap[n])
		itr.Skip()
		return true
	})
	return err
}
