package weteedsecret

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

// Make a storage key for CodeMrenclave id={{false [205]}}
//
//	代码版本
func MakeCodeMrenclaveStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "WeTEEDsecret", "CodeMrenclave")
}

var CodeMrenclaveResultDefaultBytes, _ = hex.DecodeString("00")

func GetCodeMrenclave(state state.State, bhash types.Hash) (ret []byte, err error) {
	key, err := MakeCodeMrenclaveStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(CodeMrenclaveResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetCodeMrenclaveLatest(state state.State) (ret []byte, err error) {
	key, err := MakeCodeMrenclaveStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(CodeMrenclaveResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for CodeMrsigner id={{false [205]}}
//
//	代码打包签名人
func MakeCodeMrsignerStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "WeTEEDsecret", "CodeMrsigner")
}

var CodeMrsignerResultDefaultBytes, _ = hex.DecodeString("00")

func GetCodeMrsigner(state state.State, bhash types.Hash) (ret []byte, err error) {
	key, err := MakeCodeMrsignerStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(CodeMrsignerResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetCodeMrsignerLatest(state state.State) (ret []byte, err error) {
	key, err := MakeCodeMrsignerStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(CodeMrsignerResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for NextNodeId id={{false [12]}}
//
//	The id of the next node to be created.
//	获取下一个 node id
func MakeNextNodeIdStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "WeTEEDsecret", "NextNodeId")
}

var NextNodeIdResultDefaultBytes, _ = hex.DecodeString("0000000000000000")

func GetNextNodeId(state state.State, bhash types.Hash) (ret uint64, err error) {
	key, err := MakeNextNodeIdStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(NextNodeIdResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetNextNodeIdLatest(state state.State) (ret uint64, err error) {
	key, err := MakeNextNodeIdStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(NextNodeIdResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for Nodes
//
//	dkg 节点列表
func MakeNodesStorageKey(uint640 uint64) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(uint640)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEDsecret", "Nodes", byteArgs...)
}
func GetNodes(state state.State, bhash types.Hash, uint640 uint64) (ret types1.Node, isSome bool, err error) {
	key, err := MakeNodesStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetNodesLatest(state state.State, uint640 uint64) (ret types1.Node, isSome bool, err error) {
	key, err := MakeNodesStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}
