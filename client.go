package client

import (
	"errors"
	"fmt"
	"hash"
	"time"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/config"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/centrifuge/go-substrate-rpc-client/v4/xxhash"
	"golang.org/x/crypto/blake2b"

	"github.com/wetee-dao/go-sdk/core"
	"github.com/wetee-dao/go-sdk/pallet/system"
	gtypes "github.com/wetee-dao/go-sdk/pallet/types"
)

// 区块链链接
// Chain client
type ChainClient struct {
	Api     *gsrpc.SubstrateAPI
	Meta    *types.Metadata
	Runtime *types.RuntimeVersion
	Hash    types.Hash
	Debug   bool
}

// 初始化区块连链接
// Init chain client
func ClientInit(url string, debug bool) (*ChainClient, error) {
	if url == "" {
		url = config.Default().RPCURL
	}
	api, err := gsrpc.NewSubstrateAPI(url)
	if err != nil {
		return nil, err
	}

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		return nil, err
	}
	gtypes.Meta = *meta

	genesisHash, err := api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return nil, err
	}

	runtime, err := api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return nil, err
	}

	return &ChainClient{api, meta, runtime, genesisHash, debug}, nil
}

// 检查 metadata 是否匹配
// 不匹配就更新
func (c *ChainClient) CheckMetadata() error {
	runtime, err := c.Api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return err
	}

	if c.Runtime.SpecVersion == runtime.SpecVersion {
		return nil
	}

	meta, err := c.Api.RPC.State.GetMetadataLatest()
	if err != nil {
		return err
	}

	c.Runtime = runtime
	c.Meta = meta
	gtypes.Meta = *meta
	return nil
}

// 获取区块高度
// Get block number
func (c *ChainClient) GetBlockNumber() (types.BlockNumber, error) {
	hash, err := c.Api.RPC.Chain.GetHeaderLatest()
	if err != nil {
		return 0, err
	}
	return hash.Number, nil
}

// 获取账户信息
// Get account info
func (c *ChainClient) GetAccount(address *core.Signer) (*types.AccountInfo, error) {
	key, err := types.CreateStorageKey(c.Meta, "System", "Account", address.PublicKey)
	if err != nil {
		panic(err)
	}
	var accountInfo types.AccountInfo
	_, err = c.Api.RPC.State.GetStorageLatest(key, &accountInfo)
	return &accountInfo, err
}

// 签名并提交交易
// Sign and submit transaction
func (c *ChainClient) SignAndSubmit(signer *core.Signer, runtimeCall gtypes.RuntimeCall, untilFinalized bool) error {
	accountInfo, err := c.GetAccount(signer)
	if err != nil {
		return errors.New("GetAccountInfo error: " + err.Error())
	}
	call, err := (runtimeCall).AsCall()
	if err != nil {
		return errors.New("(runtimeCall).AsCall() error: " + err.Error())
	}

	ext := core.NewExtrinsic(call)
	o := types.SignatureOptions{
		BlockHash:          c.Hash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        c.Hash,
		Nonce:              types.NewUCompactFromUInt(uint64(accountInfo.Nonce)),
		SpecVersion:        c.Runtime.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: c.Runtime.TransactionVersion,
	}

	err = ext.Sign(signer, o)
	if err != nil {
		return err
	}

	sub, err := c.Api.RPC.Author.SubmitAndWatchExtrinsic(ext.Extrinsic)
	if err != nil {
		return errors.New("SubmitAndWatchExtrinsic error: " + err.Error())
	}

	defer sub.Unsubscribe()
	timeout := time.After(30 * time.Second)
	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				extBytes, err := codec.Encode(ext)
				if err != nil {
					return err
				}

				// 计算交易的hash
				hash := blake2b.Sum256(extBytes)
				events, err := c.checkExtrinsic(hash, status.AsInBlock)
				if err != nil {
					return err
				}

				// 如果不需要等待交易确认，直接返回
				if events != nil && !untilFinalized {
					return nil
				}
			}
			if status.IsFinalized {
				return nil
			}
		case err := <-sub.Err():
			if c.Debug {
				LogWithRed("SubmitAndWatchExtrinsic ERROR", err.Error())
			}

			return err
		case <-timeout:
			fmt.Println("Extrinsic: timeout")
			return nil
		}
	}
}

// 检查交易是否成功
// Check whether the transaction is successful
func (c *ChainClient) checkExtrinsic(extHash types.Hash, blockHash types.Hash) ([]gtypes.EventRecord, error) {
	block, err := c.Api.RPC.Chain.GetBlock(blockHash)
	if err != nil {
		return nil, err
	}

	events, err := system.GetEvents(c.Api.RPC.State, blockHash)
	if err != nil {
		return nil, err
	}

	cevents := make([]gtypes.EventRecord, 0, len(events))
	for _, e := range events {
		extrinsicIndex := e.Phase.AsApplyExtrinsicField0
		ext := block.Block.Extrinsics[extrinsicIndex]
		extBytes, err := codec.Encode(ext)
		if err != nil {
			return nil, err
		}

		// 添加相关的event
		if blake2b.Sum256(extBytes) != extHash {
			cevents = append(cevents, e)
		}

		// 判断是否是当前交易的消息
		if blake2b.Sum256(extBytes) != extHash || !e.Event.IsSystem {
			continue
		}
		if e.Event.AsSystemField0.IsExtrinsicSuccess {
			if c.Debug {
				LogWithRed("Extrinsic", "ExtrinsicSuccess")
			}
			return cevents, nil
		}
		if e.Event.AsSystemField0.IsExtrinsicFailed {
			errData := e.Event.AsSystemField0.AsExtrinsicFailedDispatchError0
			if c.Debug {
				LogWithRed("Extrinsic", "ExtrinsicFailed")
			}

			b, err := errData.MarshalJSON()
			if err != nil {
				fmt.Println(err)
				return nil, err
			}

			return nil, errors.New(string(b))
		}
	}

	return nil, nil
}

// 查询 map 所有数据
// query map data list of map
func (c *ChainClient) QueryMapAll(pallet string, method string) ([]types.StorageChangeSet, error) {
	key := CreatePrefixedKey(pallet, method)

	keys, err := c.Api.RPC.State.GetKeysLatest(key)
	if err != nil {
		return []types.StorageChangeSet{}, err
	}

	set, err := c.Api.RPC.State.QueryStorageAtLatest(keys)
	if err != nil {
		return []types.StorageChangeSet{}, err
	}

	return set, nil
}

// 查询 map 所有数据
// query map data list of map4
func (c *ChainClient) QueryMapKeys(pallet string, method string, fkeys []interface{}) ([]types.StorageChangeSet, error) {
	key := CreatePrefixedKey(pallet, method)

	hashers, err := c.GetHashers(pallet, method)
	if err != nil {
		return nil, err
	}

	keys := make([]types.StorageKey, 0, len(fkeys))
	for i, sk := range fkeys {
		arg2, err := codec.Encode(sk)
		if err != nil {
			return nil, err
		}
		_, err = hashers[0].Write(arg2)
		if err != nil {
			return nil, fmt.Errorf("unable to hash args[%d]: %s Error: %v", 1, arg2, err)
		}
		keys[i] = types.StorageKey(append(key, hashers[1].Sum(nil)...))
	}

	set, err := c.Api.RPC.State.QueryStorageAtLatest(keys)
	if err != nil {
		return []types.StorageChangeSet{}, err
	}

	return set, nil
}

// 查询 double map 第一个 key 的所有数据
// query double map data list of double map
func (c *ChainClient) QueryDoubleMapAll(pallet string, method string, keyarg interface{}, at *types.Hash) ([]types.StorageChangeSet, error) {
	key, err := c.GetDoubleMapPrefixKey(pallet, method, keyarg)
	if err != nil {
		return nil, err
	}

	// query key
	var keys []types.StorageKey
	if at == nil {
		keys, err = c.Api.RPC.State.GetKeysLatest(key)
	} else {
		keys, err = c.Api.RPC.State.GetKeys(key, *at)
	}

	if err != nil {
		return nil, err
	}

	// get all data
	var set []types.StorageChangeSet
	if at == nil {
		set, err = c.Api.RPC.State.QueryStorageAtLatest(keys)
	} else {
		set, err = c.Api.RPC.State.QueryStorageAt(keys, *at)
	}
	if err != nil {
		return nil, err
	}

	return set, nil
}

// 查询 double map 第一个 key 前缀
// get double map prefix key of double map {{pallet}}.{{method}}.{{frist key}}
func (c *ChainClient) GetDoubleMapPrefixKey(pallet string, method string, keyarg interface{}) ([]byte, error) {
	arg, err := codec.Encode(keyarg)
	if err != nil {
		return nil, err
	}

	// create key prefix
	key := CreatePrefixedKey(pallet, method)
	hashers, err := c.GetHashers(pallet, method)
	if err != nil {
		return nil, err
	}

	// write key
	_, err = hashers[0].Write(arg)
	if err != nil {
		return nil, fmt.Errorf("unable to hash args[%d]: %s Error: %v", 0, arg, err)
	}

	// append hash to key
	key = append(key, hashers[0].Sum(nil)...)

	return key, nil
}

// 查询 double map 第一个 key 的所有数据
// query double map data list of double map
func (c *ChainClient) QueryDoubleMapKeys(pallet string, method string, keyarg interface{}, skeys []interface{}, at *types.Hash) ([]types.StorageChangeSet, error) {
	keys, err := c.GetDoubleMapPrefixKeys(pallet, method, keyarg, skeys)
	if err != nil {
		return nil, err
	}

	// get all data
	var set []types.StorageChangeSet
	if at == nil {
		set, err = c.Api.RPC.State.QueryStorageAtLatest(keys)
	} else {
		set, err = c.Api.RPC.State.QueryStorageAt(keys, *at)
	}
	if err != nil {
		return nil, err
	}

	return set, nil
}

// 查询 double map 第一个 key 前缀 和 多个 第二个 key
// get double map prefix key of double map {{pallet}}.{{method}}.{{frist key}}
func (c *ChainClient) GetDoubleMapPrefixKeys(pallet string, method string, keyarg interface{}, skeys []interface{}) ([]types.StorageKey, error) {
	arg, err := codec.Encode(keyarg)
	if err != nil {
		return nil, err
	}

	// create key prefix
	key := CreatePrefixedKey(pallet, method)
	hashers, err := c.GetHashers(pallet, method)
	if err != nil {
		return nil, err
	}

	// write key
	_, err = hashers[0].Write(arg)
	if err != nil {
		return nil, fmt.Errorf("unable to hash args[%d]: %s Error: %v", 0, arg, err)
	}

	// append hash to key
	key = append(key, hashers[0].Sum(nil)...)
	keys := make([]types.StorageKey, 0, len(skeys))
	for _, sk := range skeys {
		arg2, err := codec.Encode(sk)
		if err != nil {
			return nil, err
		}
		_, err = hashers[1].Write(arg2)
		if err != nil {
			return nil, fmt.Errorf("unable to hash args[%d]: %s Error: %v", 1, arg2, err)
		}
		keys = append(keys, types.StorageKey(append(key, hashers[1].Sum(nil)...)))
	}

	return keys, nil
}

// get hashers of map {{pallet}}.{{method}}
func (c *ChainClient) GetHashers(pallet, method string) ([]hash.Hash, error) {
	// get entry metadata
	// 获取储存元数据
	entryMeta, err := c.Meta.FindStorageEntryMetadata(pallet, method)
	if err != nil {
		return nil, err
	}

	// check if it's a map
	// 判断是否为map
	if !entryMeta.IsMap() {
		return nil, errors.New(pallet + "." + method + "is not map")
	}

	// get map hashers
	// 获取储存的 hasher 函数
	hashers, err := entryMeta.Hashers()
	if err != nil {
		return nil, err
	}
	return hashers, nil
}

// create prefixed key of {{pallet}}.{{method}}
func CreatePrefixedKey(pallet, method string) []byte {
	return append(xxhash.New128([]byte(pallet)).Sum(nil), xxhash.New128([]byte(method)).Sum(nil)...)
}
