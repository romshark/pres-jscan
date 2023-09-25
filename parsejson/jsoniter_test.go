package parsejson_test

import (
	"errors"

	jsoniter "github.com/json-iterator/go"
)

func NewJsoniter() Traverser {
	return Jsoniter{
		Iterator: jsoniter.NewIterator(jsoniter.ConfigDefault),
	}
}

var jsoniterTypeMap = [...]JSONValueType{
	0:                    0, // Zero value
	jsoniter.ObjectValue: JSONValueTypeObject,
	jsoniter.ArrayValue:  JSONValueTypeArray,
	jsoniter.StringValue: JSONValueTypeString,
	jsoniter.NumberValue: JSONValueTypeNumber,
	jsoniter.BoolValue:   JSONValueTypeBoolean,
	jsoniter.NilValue:    JSONValueTypeNull,
}

type Jsoniter struct{ *jsoniter.Iterator }

func (itr Jsoniter) Traverse(
	input []byte, onVar func(name []byte, t JSONValueType),
) error {
	itr.ResetBytes(input)
	if itr.WhatIsNext() != jsoniter.ObjectValue {
		return errors.New("expected object")
	}
	for f := itr.ReadObject(); f != ""; f = itr.ReadObject() {
		onVar([]byte(f), jsoniterTypeMap[itr.WhatIsNext()])
		itr.ReadAny()
	}
	return nil
}
