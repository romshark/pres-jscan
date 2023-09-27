package parsejson_test

import (
	"errors"

	"github.com/go-faster/jx"
)

func NewGofasterJX() Traverser {
	return GofasterJX{Decoder: new(jx.Decoder)}
}

type GofasterJX struct{ *jx.Decoder }

var gofasterjxTypeMap = [...]JSONValueType{
	0:         0, // Zero value
	jx.Object: JSONValueTypeObject,
	jx.Array:  JSONValueTypeArray,
	jx.String: JSONValueTypeString,
	jx.Number: JSONValueTypeNumber,
	jx.Bool:   JSONValueTypeBoolean,
	jx.Null:   JSONValueTypeNull,
}

func (dec GofasterJX) Traverse(
	input []byte, onVar func(name []byte, t JSONValueType),
) error {
	dec.ResetBytes(input)
	if dec.Next() != jx.Object {
		return errors.New("expected object")
	}
	return dec.ObjBytes(func(dec *jx.Decoder, key []byte) error {
		onVar(key, gofasterjxTypeMap[dec.Next()])
		return dec.Skip()
	})
}
