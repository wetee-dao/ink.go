package revive

import (
	types1 "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types "github.com/wetee-dao/ink.go/pallet/types"
)

// A raw EVM transaction, typically dispatched by an Ethereum JSON-RPC server.
//
// # Parameters
//
// * `payload`: The encoded [`crate::evm::TransactionSigned`].
//
// # Note
//
// This call cannot be dispatched directly; attempting to do so will result in a failed
// transaction. It serves as a wrapper for an Ethereum transaction. When submitted, the
// runtime converts it into a [`sp_runtime::generic::CheckedExtrinsic`] by recovering the
// signer and validating the transaction.
func MakeEthTransactCall(payload0 []byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsRevive: true,
		AsReviveField0: &types.PalletRevivePalletCall{
			IsEthTransact:         true,
			AsEthTransactPayload0: payload0,
		},
	}
}

// Makes a call to an account, optionally transferring some balance.
//
// # Parameters
//
//   - `dest`: Address of the contract to call.
//   - `value`: The balance to transfer from the `origin` to `dest`.
//   - `weight_limit`: The weight limit enforced when executing the constructor.
//   - `storage_deposit_limit`: The maximum amount of balance that can be charged from the
//     caller to pay for the storage consumed.
//   - `data`: The input data to pass to the contract.
//
// * If the account is a smart-contract account, the associated code will be
// executed and any value will be transferred.
// * If the account is a regular account, any value will be transferred.
// * If no account exists and the call value is not less than `existential_deposit`,
// a regular account will be created and any value will be transferred.
func MakeCallCall(dest0 [20]byte, value1 types1.UCompact, weightLimit2 types.Weight, storageDepositLimit3 types1.UCompact, data4 []byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsRevive: true,
		AsReviveField0: &types.PalletRevivePalletCall{
			IsCall:                     true,
			AsCallDest0:                dest0,
			AsCallValue1:               value1,
			AsCallWeightLimit2:         weightLimit2,
			AsCallStorageDepositLimit3: storageDepositLimit3,
			AsCallData4:                data4,
		},
	}
}

// Instantiates a contract from a previously deployed vm binary.
//
// This function is identical to [`Self::instantiate_with_code`] but without the
// code deployment step. Instead, the `code_hash` of an on-chain deployed vm binary
// must be supplied.
func MakeInstantiateCall(value0 types1.UCompact, weightLimit1 types.Weight, storageDepositLimit2 types1.UCompact, codeHash3 [32]byte, data4 []byte, salt5 types.OptionTByteArray321) types.RuntimeCall {
	return types.RuntimeCall{
		IsRevive: true,
		AsReviveField0: &types.PalletRevivePalletCall{
			IsInstantiate:                     true,
			AsInstantiateValue0:               value0,
			AsInstantiateWeightLimit1:         weightLimit1,
			AsInstantiateStorageDepositLimit2: storageDepositLimit2,
			AsInstantiateCodeHash3:            codeHash3,
			AsInstantiateData4:                data4,
			AsInstantiateSalt5:                salt5,
		},
	}
}

// Instantiates a new contract from the supplied `code` optionally transferring
// some balance.
//
// This dispatchable has the same effect as calling [`Self::upload_code`] +
// [`Self::instantiate`]. Bundling them together provides efficiency gains. Please
// also check the documentation of [`Self::upload_code`].
//
// # Parameters
//
//   - `value`: The balance to transfer from the `origin` to the newly created contract.
//   - `weight_limit`: The weight limit enforced when executing the constructor.
//   - `storage_deposit_limit`: The maximum amount of balance that can be charged/reserved
//     from the caller to pay for the storage consumed.
//   - `code`: The contract code to deploy in raw bytes.
//   - `data`: The input data to pass to the contract constructor.
//   - `salt`: Used for the address derivation. If `Some` is supplied then `CREATE2`
//     semantics are used. If `None` then `CRATE1` is used.
//
// Instantiation is executed as follows:
//
// - The supplied `code` is deployed, and a `code_hash` is created for that code.
// - If the `code_hash` already exists on the chain the underlying `code` will be shared.
// - The destination address is computed based on the sender, code_hash and the salt.
// - The smart-contract account is created at the computed address.
// - The `value` is transferred to the new account.
// - The `deploy` function is executed in the context of the newly-created account.
func MakeInstantiateWithCodeCall(value0 types1.UCompact, weightLimit1 types.Weight, storageDepositLimit2 types1.UCompact, code3 []byte, data4 []byte, salt5 types.OptionTByteArray321) types.RuntimeCall {
	return types.RuntimeCall{
		IsRevive: true,
		AsReviveField0: &types.PalletRevivePalletCall{
			IsInstantiateWithCode:                     true,
			AsInstantiateWithCodeValue0:               value0,
			AsInstantiateWithCodeWeightLimit1:         weightLimit1,
			AsInstantiateWithCodeStorageDepositLimit2: storageDepositLimit2,
			AsInstantiateWithCodeCode3:                code3,
			AsInstantiateWithCodeData4:                data4,
			AsInstantiateWithCodeSalt5:                salt5,
		},
	}
}

// Same as [`Self::instantiate_with_code`], but intended to be dispatched **only**
// by an EVM transaction through the EVM compatibility layer.
//
// # Parameters
//
//   - `value`: The balance to transfer from the `origin` to the newly created contract.
//   - `weight_limit`: The gas limit used to derive the transaction weight for transaction
//     payment
//   - `eth_gas_limit`: The Ethereum gas limit governing the resource usage of the execution
//   - `code`: The contract code to deploy in raw bytes.
//   - `data`: The input data to pass to the contract constructor.
//   - `transaction_encoded`: The RLP encoding of the signed Ethereum transaction,
//     represented as [crate::evm::TransactionSigned], provided by the Ethereum wallet. This
//     is used for building the Ethereum transaction root.
//   - effective_gas_price: the price of a unit of gas
//   - encoded len: the byte code size of the `eth_transact` extrinsic
//
// Calling this dispatchable ensures that the origin's nonce is bumped only once,
// via the `CheckNonce` transaction extension. In contrast, [`Self::instantiate_with_code`]
// also bumps the nonce after contract instantiation, since it may be invoked multiple
// times within a batch call transaction.
func MakeEthInstantiateWithCodeCall(value0 [4]uint64, weightLimit1 types.Weight, ethGasLimit2 [4]uint64, code3 []byte, data4 []byte, transactionEncoded5 []byte, effectiveGasPrice6 [4]uint64, encodedLen7 uint32) types.RuntimeCall {
	return types.RuntimeCall{
		IsRevive: true,
		AsReviveField0: &types.PalletRevivePalletCall{
			IsEthInstantiateWithCode:                    true,
			AsEthInstantiateWithCodeValue0:              value0,
			AsEthInstantiateWithCodeWeightLimit1:        weightLimit1,
			AsEthInstantiateWithCodeEthGasLimit2:        ethGasLimit2,
			AsEthInstantiateWithCodeCode3:               code3,
			AsEthInstantiateWithCodeData4:               data4,
			AsEthInstantiateWithCodeTransactionEncoded5: transactionEncoded5,
			AsEthInstantiateWithCodeEffectiveGasPrice6:  effectiveGasPrice6,
			AsEthInstantiateWithCodeEncodedLen7:         encodedLen7,
		},
	}
}

// Same as [`Self::call`], but intended to be dispatched **only**
// by an EVM transaction through the EVM compatibility layer.
//
// # Parameters
//
//   - `dest`: The Ethereum address of the account to be called
//   - `value`: The balance to transfer from the `origin` to the newly created contract.
//   - `weight_limit`: The gas limit used to derive the transaction weight for transaction
//     payment
//   - `eth_gas_limit`: The Ethereum gas limit governing the resource usage of the execution
//   - `data`: The input data to pass to the contract constructor.
//   - `transaction_encoded`: The RLP encoding of the signed Ethereum transaction,
//     represented as [crate::evm::TransactionSigned], provided by the Ethereum wallet. This
//     is used for building the Ethereum transaction root.
//   - effective_gas_price: the price of a unit of gas
//   - encoded len: the byte code size of the `eth_transact` extrinsic
func MakeEthCallCall(dest0 [20]byte, value1 [4]uint64, weightLimit2 types.Weight, ethGasLimit3 [4]uint64, data4 []byte, transactionEncoded5 []byte, effectiveGasPrice6 [4]uint64, encodedLen7 uint32) types.RuntimeCall {
	return types.RuntimeCall{
		IsRevive: true,
		AsReviveField0: &types.PalletRevivePalletCall{
			IsEthCall:                    true,
			AsEthCallDest0:               dest0,
			AsEthCallValue1:              value1,
			AsEthCallWeightLimit2:        weightLimit2,
			AsEthCallEthGasLimit3:        ethGasLimit3,
			AsEthCallData4:               data4,
			AsEthCallTransactionEncoded5: transactionEncoded5,
			AsEthCallEffectiveGasPrice6:  effectiveGasPrice6,
			AsEthCallEncodedLen7:         encodedLen7,
		},
	}
}

// Executes a Substrate runtime call from an Ethereum transaction.
//
// This dispatchable is intended to be called **only** through the EVM compatibility
// layer. The provided call will be dispatched using `RawOrigin::Signed`.
//
// # Parameters
//
// * `origin`: Must be an [`Origin::EthTransaction`] origin.
// * `call`: The Substrate runtime call to execute.
// * `transaction_encoded`: The RLP encoding of the Ethereum transaction,
func MakeEthSubstrateCallCall(call0 types.RuntimeCall, transactionEncoded1 []byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsRevive: true,
		AsReviveField0: &types.PalletRevivePalletCall{
			IsEthSubstrateCall:                    true,
			AsEthSubstrateCallCall0:               &call0,
			AsEthSubstrateCallTransactionEncoded1: transactionEncoded1,
		},
	}
}

// Upload new `code` without instantiating a contract from it.
//
// If the code does not already exist a deposit is reserved from the caller
// The size of the reserve depends on the size of the supplied `code`.
//
// # Note
//
// Anyone can instantiate a contract from any uploaded code and thus prevent its removal.
// To avoid this situation a constructor could employ access control so that it can
// only be instantiated by permissioned entities. The same is true when uploading
// through [`Self::instantiate_with_code`].
//
// If the refcount of the code reaches zero after terminating the last contract that
// references this code, the code will be removed automatically.
func MakeUploadCodeCall(code0 []byte, storageDepositLimit1 types1.UCompact) types.RuntimeCall {
	return types.RuntimeCall{
		IsRevive: true,
		AsReviveField0: &types.PalletRevivePalletCall{
			IsUploadCode:                     true,
			AsUploadCodeCode0:                code0,
			AsUploadCodeStorageDepositLimit1: storageDepositLimit1,
		},
	}
}

// Remove the code stored under `code_hash` and refund the deposit to its owner.
//
// A code can only be removed by its original uploader (its owner) and only if it is
// not used by any contract.
func MakeRemoveCodeCall(codeHash0 [32]byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsRevive: true,
		AsReviveField0: &types.PalletRevivePalletCall{
			IsRemoveCode:          true,
			AsRemoveCodeCodeHash0: codeHash0,
		},
	}
}

// Privileged function that changes the code of an existing contract.
//
// This takes care of updating refcounts and all other necessary operations. Returns
// an error if either the `code_hash` or `dest` do not exist.
//
// # Note
//
// This does **not** change the address of the contract in question. This means
// that the contract address is no longer derived from its code hash after calling
// this dispatchable.
func MakeSetCodeCall(dest0 [20]byte, codeHash1 [32]byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsRevive: true,
		AsReviveField0: &types.PalletRevivePalletCall{
			IsSetCode:          true,
			AsSetCodeDest0:     dest0,
			AsSetCodeCodeHash1: codeHash1,
		},
	}
}

// Register the callers account id so that it can be used in contract interactions.
//
// This will error if the origin is already mapped or is a eth native `Address20`. It will
// take a deposit that can be released by calling [`Self::unmap_account`].
func MakeMapAccountCall() types.RuntimeCall {
	return types.RuntimeCall{
		IsRevive: true,
		AsReviveField0: &types.PalletRevivePalletCall{
			IsMapAccount: true,
		},
	}
}

// Unregister the callers account id in order to free the deposit.
//
// There is no reason to ever call this function other than freeing up the deposit.
// This is only useful when the account should no longer be used.
func MakeUnmapAccountCall() types.RuntimeCall {
	return types.RuntimeCall{
		IsRevive: true,
		AsReviveField0: &types.PalletRevivePalletCall{
			IsUnmapAccount: true,
		},
	}
}

// Dispatch an `call` with the origin set to the callers fallback address.
//
// Every `AccountId32` can control its corresponding fallback account. The fallback account
// is the `AccountId20` with the last 12 bytes set to `0xEE`. This is essentially a
// recovery function in case an `AccountId20` was used without creating a mapping first.
func MakeDispatchAsFallbackAccountCall(call0 types.RuntimeCall) types.RuntimeCall {
	return types.RuntimeCall{
		IsRevive: true,
		AsReviveField0: &types.PalletRevivePalletCall{
			IsDispatchAsFallbackAccount:      true,
			AsDispatchAsFallbackAccountCall0: &call0,
		},
	}
}
