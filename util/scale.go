package util

import (
	"bytes"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
)

func Encode(value any) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	err := scale.NewEncoder(buf).Encode(value)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
