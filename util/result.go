package util

import (
	"encoding/hex"
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
)

type H160 [20]byte

func (h *H160) String() string {
	return "0x" + hex.EncodeToString(h[:])
}

type Result[T, Err any] struct {
	IsErr bool
	E     Err
	V     T
}

func (r Result[T, Err]) Encode(encoder scale.Encoder) (err error) {
	if r.IsErr {
		err = encoder.PushByte(1)
		if err != nil {
			return err
		}
		err = encoder.Encode(r.E)
		if err != nil {
			return err
		}
		return nil
	}

	err = encoder.PushByte(0)
	if err != nil {
		return err
	}
	err = encoder.Encode(r.V)
	if err != nil {
		return err
	}
	return nil
}

func (r *Result[T, Err]) Decode(decoder scale.Decoder) (err error) {
	variant, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch variant {
	case 0:
		err = decoder.Decode(&r.V)
		if err != nil {
			return err
		}
		return
	case 1:
		r.IsErr = true
		var tmp Err
		err = decoder.Decode(&tmp)
		if err != nil {
			return err
		}
		r.E = tmp
		return
	default:
		return fmt.Errorf("Unrecognized variant")
	}
}
