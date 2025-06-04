package util

import (
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
)

// New none value of Option
func NewNone[T any]() Option[T] {
	return Option[T]{
		IsNone: true,
	}
}

// New value of Option
func NewSome[T any](v T) Option[T] {
	return Option[T]{
		IsNone: false,
		V:      v,
	}
}

// Option is a type for rust Option
type Option[T any] struct {
	IsNone bool
	V      T
}

func (ty Option[T]) Encode(encoder scale.Encoder) (err error) {
	if ty.IsNone {
		err = encoder.PushByte(0)
		if err != nil {
			return err
		}
		return nil
	}

	err = encoder.PushByte(1)
	if err != nil {
		return err
	}

	return encoder.Encode(ty.V)
}

func (ty *Option[T]) Decode(decoder scale.Decoder) (err error) {
	variant, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch variant {
	case 0:
		ty.IsNone = true
		return
	case 1:
		err = decoder.Decode(&ty.V)
		if err != nil {
			return err
		}
		return
	default:
		return fmt.Errorf("unrecognized enum")
	}
}
