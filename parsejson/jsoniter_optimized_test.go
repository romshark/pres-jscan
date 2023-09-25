package parsejson_test

import (
	"errors"

	jsoniter "github.com/json-iterator/go"
)

func NewJsoniterOptimized() Traverser {
	return JsoniterOptimized{
		Iterator: jsoniter.NewIterator(jsoniter.ConfigFastest),
	}
}

type JsoniterOptimized struct{ *jsoniter.Iterator }

func (itr JsoniterOptimized) Traverse(
	input []byte, onVar func(name []byte, t JSONValueType),
) error {
	itr.ResetBytes(input)
	if itr.WhatIsNext() != jsoniter.ObjectValue {
		return errors.New("expected object")
	}
	itr.ReadObjectCB(func(itr *jsoniter.Iterator, s string) bool {
		onVar([]byte(s), jsoniterTypeMap[itr.WhatIsNext()])
		itr.Skip()
		return true
	})
	return nil
}
