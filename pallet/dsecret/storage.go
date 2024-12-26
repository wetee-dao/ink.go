package dsecret

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/go-sdk/pallet/types"
)

// Make a storage key for CodeSignature id={{false [14]}}
//
//	DKG 代码版本
func MakeCodeSignatureStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "DSecret", "CodeSignature")
}

var CodeSignatureResultDefaultBytes, _ = hex.DecodeString("00")

func GetCodeSignature(state state.State, bhash types.Hash) (ret []byte, err error) {
	key, err := MakeCodeSignatureStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(CodeSignatureResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetCodeSignatureLatest(state state.State) (ret []byte, err error) {
	key, err := MakeCodeSignatureStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(CodeSignatureResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for CodeSigner id={{false [14]}}
//
//	DKG 代码打包签名人
func MakeCodeSignerStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "DSecret", "CodeSigner")
}

var CodeSignerResultDefaultBytes, _ = hex.DecodeString("00")

func GetCodeSigner(state state.State, bhash types.Hash) (ret []byte, err error) {
	key, err := MakeCodeSignerStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(CodeSignerResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetCodeSignerLatest(state state.State) (ret []byte, err error) {
	key, err := MakeCodeSignerStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(CodeSignerResultDefaultBytes, &ret)
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
	return types.CreateStorageKey(&types1.Meta, "DSecret", "NextNodeId")
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
	return types.CreateStorageKey(&types1.Meta, "DSecret", "Nodes", byteArgs...)
}
func GetNodes(state state.State, bhash types.Hash, uint640 uint64) (ret [32]byte, isSome bool, err error) {
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
func GetNodesLatest(state state.State, uint640 uint64) (ret [32]byte, isSome bool, err error) {
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

// Make a storage key for NodePubServers
//
//	dkg pub server
//	dkg pub 服务
func MakeNodePubServersStorageKey(uint640 uint64) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(uint640)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "DSecret", "NodePubServers", byteArgs...)
}
func GetNodePubServers(state state.State, bhash types.Hash, uint640 uint64) (ret types1.P2PAddr, isSome bool, err error) {
	key, err := MakeNodePubServersStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetNodePubServersLatest(state state.State, uint640 uint64) (ret types1.P2PAddr, isSome bool, err error) {
	key, err := MakeNodePubServersStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}
