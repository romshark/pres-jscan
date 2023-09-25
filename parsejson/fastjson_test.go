package parsejson_test

import "github.com/valyala/fastjson"

func NewFastjson() Traverser {
	return Fastjson{Parser: new(fastjson.Parser)}
}

type Fastjson struct{ *fastjson.Parser }

var fastjsonTypeMap = [...]JSONValueType{
	fastjson.TypeObject: JSONValueTypeObject,
	fastjson.TypeArray:  JSONValueTypeArray,
	fastjson.TypeString: JSONValueTypeString,
	fastjson.TypeNumber: JSONValueTypeNumber,
	fastjson.TypeTrue:   JSONValueTypeBoolean,
	fastjson.TypeFalse:  JSONValueTypeBoolean,
	fastjson.TypeNull:   JSONValueTypeNull,
}

func (dec Fastjson) Traverse(
	input []byte, onVar func(name []byte, t JSONValueType),
) error {
	val, err := dec.Parser.ParseBytes(input)
	if err != nil {
		return err
	}
	o, err := val.Object()
	if err != nil {
		return err
	}
	o.Visit(func(key []byte, v *fastjson.Value) {
		onVar(key, fastjsonTypeMap[v.Type()])
	})
	return nil
}
