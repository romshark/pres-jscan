package parsejson_test

import (
	"encoding/json"

	"github.com/bytedance/sonic"
)

func NewBytedanceSonic() Traverser { return BytedanceSonic{API: sonic.ConfigDefault} }

type BytedanceSonic struct{ sonic.API }

func (dec BytedanceSonic) Traverse(
	input []byte, onVar func(name []byte, t JSONValueType),
) error {
	var m map[string]json.RawMessage
	if err := sonic.ConfigDefault.Unmarshal(input, &m); err != nil {
		return err
	}
	for k, v := range m {
		onVar(unsafeStrToBytes(k), encodingJSONOptCharMap[v[0]])
	}
	return nil
}
