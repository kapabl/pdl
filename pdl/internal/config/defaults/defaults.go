package defaults

import _ "embed"

//go:embed common.pdl.config.json
var commonConfig []byte

func Common() ([]byte, error) {
	return clone(commonConfig), nil
}

func clone(data []byte) []byte {
	if data == nil {
		return nil
	}
	result := make([]byte, len(data))
	copy(result, data)
	return result
}
