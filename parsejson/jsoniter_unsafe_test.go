package parsejson_test

import (
	"errors"

	jsoniter "github.com/json-iterator/go"
)

func NewJsoniterUnsafe() Traverser {
	return JsoniterUnsafe{
		Iterator: jsoniter.NewIterator(jsoniter.ConfigDefault),
	}
}

type JsoniterUnsafe struct{ *jsoniter.Iterator }

func (itr JsoniterUnsafe) Traverse(
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
		onVar(unsafeStrToBytes(s), jsoniterTypeMap[n])
		itr.Skip()
		return true
	})
	return err
}
