package messagequeue

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types1 "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types "github.com/wetee-dao/go-sdk/pallet/types"
)

// Make a storage key for BookStateFor
//
//	The index of the first and last (non-empty) pages.
func MakeBookStateForStorageKey(aggregateMessageOrigin0 types.AggregateMessageOrigin) (types1.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(aggregateMessageOrigin0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types1.CreateStorageKey(&types.Meta, "MessageQueue", "BookStateFor", byteArgs...)
}

var BookStateForResultDefaultBytes, _ = hex.DecodeString("0000000000000000000000000000000000000000000000000000000000")

func GetBookStateFor(state state.State, bhash types1.Hash, aggregateMessageOrigin0 types.AggregateMessageOrigin) (ret types.BookState, err error) {
	key, err := MakeBookStateForStorageKey(aggregateMessageOrigin0)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(BookStateForResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetBookStateForLatest(state state.State, aggregateMessageOrigin0 types.AggregateMessageOrigin) (ret types.BookState, err error) {
	key, err := MakeBookStateForStorageKey(aggregateMessageOrigin0)
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(BookStateForResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for ServiceHead id={{false [121]}}
//
//	The origin at which we should begin servicing.
func MakeServiceHeadStorageKey() (types1.StorageKey, error) {
	return types1.CreateStorageKey(&types.Meta, "MessageQueue", "ServiceHead")
}
func GetServiceHead(state state.State, bhash types1.Hash) (ret types.AggregateMessageOrigin, isSome bool, err error) {
	key, err := MakeServiceHeadStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetServiceHeadLatest(state state.State) (ret types.AggregateMessageOrigin, isSome bool, err error) {
	key, err := MakeServiceHeadStorageKey()
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for Pages
//
//	The map of page indices to pages.
func MakePagesStorageKey(tupleOfAggregateMessageOriginUint320 types.AggregateMessageOrigin, tupleOfAggregateMessageOriginUint321 uint32) (types1.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfAggregateMessageOriginUint320)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfAggregateMessageOriginUint321)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types1.CreateStorageKey(&types.Meta, "MessageQueue", "Pages", byteArgs...)
}
func GetPages(state state.State, bhash types1.Hash, tupleOfAggregateMessageOriginUint320 types.AggregateMessageOrigin, tupleOfAggregateMessageOriginUint321 uint32) (ret types.Page, isSome bool, err error) {
	key, err := MakePagesStorageKey(tupleOfAggregateMessageOriginUint320, tupleOfAggregateMessageOriginUint321)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetPagesLatest(state state.State, tupleOfAggregateMessageOriginUint320 types.AggregateMessageOrigin, tupleOfAggregateMessageOriginUint321 uint32) (ret types.Page, isSome bool, err error) {
	key, err := MakePagesStorageKey(tupleOfAggregateMessageOriginUint320, tupleOfAggregateMessageOriginUint321)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}
