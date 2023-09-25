package parsejson_test

import "github.com/romshark/jscan/v2"

func NewJscan() Traverser {
	return RomsharkJscan{Parser: jscan.NewParser[[]byte](1024)}
}

type RomsharkJscan struct{ *jscan.Parser[[]byte] }

var romsharkjscanTypeMap = [...]JSONValueType{
	0:                     0, // Zero value
	jscan.ValueTypeObject: JSONValueTypeObject,
	jscan.ValueTypeArray:  JSONValueTypeArray,
	jscan.ValueTypeString: JSONValueTypeString,
	jscan.ValueTypeNumber: JSONValueTypeNumber,
	jscan.ValueTypeFalse:  JSONValueTypeBoolean,
	jscan.ValueTypeTrue:   JSONValueTypeBoolean,
	jscan.ValueTypeNull:   JSONValueTypeNull,
}

func (itr RomsharkJscan) Traverse(
	input []byte, onVar func(name []byte, t JSONValueType),
) error {
	err := itr.Parser.Scan(input, func(i *jscan.Iterator[[]byte]) (err bool) {
		switch i.Level() {
		case 0:
			return i.ValueType() != jscan.ValueTypeObject
		case 1:
			k := i.Key()
			onVar(k[1:len(k)-1], romsharkjscanTypeMap[i.ValueType()])
		}
		return false
	})
	if err.IsErr() {
		return err
	}
	return nil
}
