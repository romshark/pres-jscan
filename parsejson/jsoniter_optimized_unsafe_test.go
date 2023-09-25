package parsejson_test

import (
	"errors"

	jsoniter "github.com/json-iterator/go"
)

func NewJsoniterOptimizedUnsafe() Traverser {
	return JsoniterOptimizedUnsafe{
		Iterator: jsoniter.NewIterator(jsoniter.ConfigDefault),
	}
}

type JsoniterOptimizedUnsafe struct{ *jsoniter.Iterator }

func (itr JsoniterOptimizedUnsafe) Traverse(
	input []byte, onVar func(name []byte, t JSONValueType),
) error {
	itr.ResetBytes(input)
	if itr.WhatIsNext() != jsoniter.ObjectValue {
		return errors.New("expected object")
	}
	itr.ReadObjectCB(func(itr *jsoniter.Iterator, s string) bool {
		onVar(unsafeStrToBytes(s), jsoniterTypeMap[itr.WhatIsNext()])
		itr.Skip()
		return true
	})
	return nil
}
