package parsejson_test

import "encoding/json"

func NewEncodingJSON() GraphQLVariablesTraverser { return EncodingJSON{} }

type EncodingJSON struct{}

func (EncodingJSON) TraverseJSON(
	input []byte, onVar func(name []byte, t JSONValueType),
) error {
	var m map[string]any
	if err := json.Unmarshal(input, &m); err != nil {
		return err
	}
	for k, v := range m {
		tp := JSONValueTypeNull
		switch v.(type) {
		case map[string]any:
			tp = JSONValueTypeObject
		case []any:
			tp = JSONValueTypeArray
		case string:
			tp = JSONValueTypeString
		case int, float64:
			tp = JSONValueTypeNumber
		case bool:
			tp = JSONValueTypeBoolean
		}
		onVar([]byte(k), tp)
	}
	return nil
}
