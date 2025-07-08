package ink

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"log"
	"math/big"
	"time"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/config"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry"
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/extrinsic"
	"github.com/centrifuge/go-substrate-rpc-client/v4/xxhash"
	"golang.org/x/crypto/blake2b"

	"github.com/wetee-dao/ink.go/pallet/revive"
	"github.com/wetee-dao/ink.go/pallet/system"
	gtypes "github.com/wetee-dao/ink.go/pallet/types"
	"github.com/wetee-dao/ink.go/util"
)

// 区块链链接
// Chain client
type ChainClient struct {
	Api      *gsrpc.SubstrateAPI
	Meta     *types.Metadata
	Runtime  *types.RuntimeVersion
	ErrorMap registry.ErrorRegistry
	Hash     types.Hash
	Debug    bool
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

	errMap, err := InitErrors(meta)
	if err != nil {
		return nil, err
	}

	genesisHash, err := api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return nil, err
	}

	runtime, err := api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return nil, err
	}

	return &ChainClient{api, meta, runtime, errMap, genesisHash, debug}, nil
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

	errMap, err := InitErrors(meta)
	if err != nil {
		return err
	}

	c.ErrorMap = errMap
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
func (c *ChainClient) GetAccount(address SignerType) (*types.AccountInfo, error) {
	key, err := types.CreateStorageKey(c.Meta, "System", "Account", address.Public())
	if err != nil {
		panic(err)
	}
	var accountInfo types.AccountInfo
	_, err = c.Api.RPC.State.GetStorageLatest(key, &accountInfo)
	return &accountInfo, err
}

// 签名并提交交易
// Sign and submit transaction
func (c *ChainClient) SignAndSubmit(signer SignerType, call types.Call, untilFinalized bool) error {
	accountInfo, err := c.GetAccount(signer)
	if err != nil {
		return errors.New("GetAccountInfo error: " + err.Error())
	}

	ext := NewExtrinsic(call)
	err = ext.Sign(signer, c.Meta, extrinsic.WithEra(types.ExtrinsicEra{IsImmortalEra: true}, c.Hash),
		extrinsic.WithNonce(types.NewUCompactFromUInt(uint64(accountInfo.Nonce))),
		extrinsic.WithTip(types.NewUCompactFromUInt(0)),
		extrinsic.WithSpecVersion(c.Runtime.SpecVersion),
		extrinsic.WithTransactionVersion(c.Runtime.TransactionVersion),
		extrinsic.WithGenesisHash(c.Hash),
	)
	if err != nil {
		return err
	}

	sub, err := c.Api.RPC.Author.SubmitAndWatchExtrinsic(ext.Extrinsic)
	if err != nil {
		return errors.New("SubmitAndWatchExtrinsic error: " + err.Error())
	}

	defer sub.Unsubscribe()
	timeout := time.After(30 * time.Second)

	extBytes, err := codec.Encode(ext.Extrinsic)
	if err != nil {
		return errors.New("SubmitAndWatchExtrinsic error: " + err.Error())
	}
	hash := blake2b.Sum256(extBytes)

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				_, success, err := c.checkExtrinsic(hash, status.AsInBlock)
				if err != nil {
					return err
				}

				if success && c.Debug {
					util.LogWithPurple("Extrinsic", "InBlock")
				}

				if success && !untilFinalized {
					return nil
				}
			} else if status.IsFinalized {
				_, success, err := c.checkExtrinsic(hash, status.AsFinalized)
				if err != nil {
					return err
				}
				if success {
					if c.Debug {
						util.LogWithPurple("Extrinsic", "Finalized")
					}
					return nil
				}
			} else if status.IsDropped {
				util.LogWithRed("SubmitAndWatchExtrinsic Dropped")
			} else if status.IsUsurped {
				util.LogWithRed("SubmitAndWatchExtrinsic Usurped")
			}
		case err := <-sub.Err():
			if c.Debug {
				util.LogWithRed("SubmitAndWatchExtrinsic ERROR", err.Error())
			}

			return err
		case <-timeout:
			util.LogWithRed("SubmitAndWatchExtrinsic ERROR: timeout")
			return nil
		}
	}
}

// 检查交易是否成功
// Check whether the transaction is successful
func (c *ChainClient) checkExtrinsic(extHash types.Hash, blockHash types.Hash) ([]gtypes.EventRecord, bool, error) {
	block, err := c.Api.RPC.Chain.GetBlock(blockHash)
	if err != nil {
		return nil, false, err
	}

	events, err := system.GetEvents(c.Api.RPC.State, blockHash)
	if err != nil {
		return nil, false, err
	}

	cevents := make([]gtypes.EventRecord, 0, len(events))
	for _, e := range events {
		extrinsicIndex := e.Phase.AsApplyExtrinsicField0
		ext := block.Block.Extrinsics[extrinsicIndex]
		extBytes, err := hex.DecodeString(ext[2:])
		if err != nil {
			return nil, false, err
		}
		eventExtHash := blake2b.Sum256(extBytes)

		// 添加相关的event
		if eventExtHash != extHash {
			cevents = append(cevents, e)
		}

		// 判断是否是当前交易的消息
		if eventExtHash != extHash || !e.Event.IsSystem {
			continue
		}

		// 判断是否是交易成功的消息
		if e.Event.AsSystemField0.IsExtrinsicSuccess {
			// if c.Debug {
			// 	util.LogWithPurple("Extrinsic", "ExtrinsicSuccess")
			// }
			return cevents, true, nil
		}

		// 获取交易失败的消息
		if e.Event.AsSystemField0.IsExtrinsicFailed {
			errData := e.Event.AsSystemField0.AsExtrinsicFailedDispatchError0
			if c.Debug {
				util.LogWithPurple("Extrinsic", "ExtrinsicFailed")
			}

			var errInfo error
			// 判断是否是区块链模块错误
			if errData.IsModule {
				merr := errData.AsModuleField0
				info, ierr := c.GetErrorInfo(merr.Index, merr.Error)
				if ierr == nil {
					errInfo = errors.New("tx: module error " + info.Name)
				} else {
					errInfo = errors.New("tx: unknown module error ")
				}
			} else {
				b, err := errData.MarshalJSON()
				if err != nil {
					fmt.Println(err)
					return nil, false, err
				}
				errInfo = errors.New(string(b))
			}

			return nil, false, errInfo
		}
	}

	return nil, false, nil
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
func (c *ChainClient) QueryMapKeys(pallet string, method string, fkeys []any) ([]types.StorageChangeSet, error) {
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
func (c *ChainClient) QueryDoubleMapAll(pallet string, method string, keyarg any, at *types.Hash) ([]types.StorageChangeSet, error) {
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
func (c *ChainClient) GetDoubleMapPrefixKey(pallet string, method string, keyarg any) ([]byte, error) {
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
func (c *ChainClient) QueryDoubleMapKeys(pallet string, method string, keyarg any, skeys []any, at *types.Hash) ([]types.StorageChangeSet, error) {
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
func (c *ChainClient) GetDoubleMapPrefixKeys(pallet string, method string, keyarg any, skeys []any) ([]types.StorageKey, error) {
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

// Get hashers of map {{pallet}}.{{method}}
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

// Call runtime api
func (c *ChainClient) CallRuntimeApi(pallet, method string, args []any, result any) error {
	var buffer bytes.Buffer
	var err error
	encoder := scale.NewEncoder(&buffer)

	// Encode the arguments
	for _, arg := range args {
		err = encoder.Encode(arg)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	// Call runtime api
	var rawResult string
	err = c.Api.Client.Call(&rawResult, "state_call", pallet+"_"+method, "0x"+hex.EncodeToString(buffer.Bytes()))
	if err != nil {
		return err
	}

	// Decode the raw result from hex to bytes
	rawResult = rawResult[2:]
	resultBytes, err := hex.DecodeString(rawResult)
	if err != nil {
		log.Fatalf("Failed to decode result: %v", err)
	}

	// Decode the result using scale.Decoder
	return scale.NewDecoder(bytes.NewReader(resultBytes)).Decode(result)
}

// Get balance of h160
func (c *ChainClient) BalanceOfH160(address string) (types.U128, error) {
	balance := types.NewU128(*big.NewInt(0))
	bt, err := util.HexToH160(address)
	if err != nil {
		return types.U128{}, err
	}
	err = c.CallRuntimeApi("ReviveApi", "balance", []any{bt}, &balance)
	if err != nil {
		return types.U128{}, err
	}

	return balance, nil
}

func (c *ChainClient) MapReviveAccount(signer SignerType) error {
	runtimeCall := revive.MakeMapAccountCall()

	call, err := (runtimeCall).AsCall()
	if err != nil {
		return errors.New("(runtimeCall).AsCall() error: " + err.Error())
	}

	return c.SignAndSubmit(signer, call, true)
}

// Get block gas limit
func (c *ChainClient) InkBlockGasLimit(address [32]byte) error {
	balance := types.NewU128(*big.NewInt(0))
	err := c.CallRuntimeApi("ReviveApi", "block_gas_limit", []any{}, &balance)
	return err
}

// Create prefixed key of {{pallet}}.{{method}}
func CreatePrefixedKey(pallet, method string) []byte {
	return append(xxhash.New128([]byte(pallet)).Sum(nil), xxhash.New128([]byte(method)).Sum(nil)...)
}
