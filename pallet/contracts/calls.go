package contracts

import (
	types1 "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types "github.com/wetee-dao/go-sdk/pallet/types"
)

// Deprecated version if [`Self::call`] for use in an in-storage `Call`.
func MakeCallOldWeightCall(dest0 types.MultiAddress, value1 types1.UCompact, gasLimit2 types1.UCompact, storageDepositLimit3 types.OptionTUCompact, data4 []byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsContracts: true,
		AsContractsField0: &types.PalletContractsPalletCall{
			IsCallOldWeight:                     true,
			AsCallOldWeightDest0:                dest0,
			AsCallOldWeightValue1:               value1,
			AsCallOldWeightGasLimit2:            gasLimit2,
			AsCallOldWeightStorageDepositLimit3: storageDepositLimit3,
			AsCallOldWeightData4:                data4,
		},
	}
}

// Deprecated version if [`Self::instantiate_with_code`] for use in an in-storage `Call`.
func MakeInstantiateWithCodeOldWeightCall(value0 types1.UCompact, gasLimit1 types1.UCompact, storageDepositLimit2 types.OptionTUCompact, code3 []byte, data4 []byte, salt5 []byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsContracts: true,
		AsContractsField0: &types.PalletContractsPalletCall{
			IsInstantiateWithCodeOldWeight:                     true,
			AsInstantiateWithCodeOldWeightValue0:               value0,
			AsInstantiateWithCodeOldWeightGasLimit1:            gasLimit1,
			AsInstantiateWithCodeOldWeightStorageDepositLimit2: storageDepositLimit2,
			AsInstantiateWithCodeOldWeightCode3:                code3,
			AsInstantiateWithCodeOldWeightData4:                data4,
			AsInstantiateWithCodeOldWeightSalt5:                salt5,
		},
	}
}

// Deprecated version if [`Self::instantiate`] for use in an in-storage `Call`.
func MakeInstantiateOldWeightCall(value0 types1.UCompact, gasLimit1 types1.UCompact, storageDepositLimit2 types.OptionTUCompact, codeHash3 [32]byte, data4 []byte, salt5 []byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsContracts: true,
		AsContractsField0: &types.PalletContractsPalletCall{
			IsInstantiateOldWeight:                     true,
			AsInstantiateOldWeightValue0:               value0,
			AsInstantiateOldWeightGasLimit1:            gasLimit1,
			AsInstantiateOldWeightStorageDepositLimit2: storageDepositLimit2,
			AsInstantiateOldWeightCodeHash3:            codeHash3,
			AsInstantiateOldWeightData4:                data4,
			AsInstantiateOldWeightSalt5:                salt5,
		},
	}
}

// Upload new `code` without instantiating a contract from it.
//
// If the code does not already exist a deposit is reserved from the caller
// and unreserved only when [`Self::remove_code`] is called. The size of the reserve
// depends on the size of the supplied `code`.
//
// If the code already exists in storage it will still return `Ok` and upgrades
// the in storage version to the current
// [`InstructionWeights::version`](InstructionWeights).
//
//   - `determinism`: If this is set to any other value but [`Determinism::Enforced`] then
//     the only way to use this code is to delegate call into it from an offchain execution.
//     Set to [`Determinism::Enforced`] if in doubt.
//
// # Note
//
// Anyone can instantiate a contract from any uploaded code and thus prevent its removal.
// To avoid this situation a constructor could employ access control so that it can
// only be instantiated by permissioned entities. The same is true when uploading
// through [`Self::instantiate_with_code`].
//
// Use [`Determinism::Relaxed`] exclusively for non-deterministic code. If the uploaded
// code is deterministic, specifying [`Determinism::Relaxed`] will be disregarded and
// result in higher gas costs.
func MakeUploadCodeCall(code0 []byte, storageDepositLimit1 types.OptionTUCompact, determinism2 types.Determinism) types.RuntimeCall {
	return types.RuntimeCall{
		IsContracts: true,
		AsContractsField0: &types.PalletContractsPalletCall{
			IsUploadCode:                     true,
			AsUploadCodeCode0:                code0,
			AsUploadCodeStorageDepositLimit1: storageDepositLimit1,
			AsUploadCodeDeterminism2:         determinism2,
		},
	}
}

// Remove the code stored under `code_hash` and refund the deposit to its owner.
//
// A code can only be removed by its original uploader (its owner) and only if it is
// not used by any contract.
func MakeRemoveCodeCall(codeHash0 [32]byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsContracts: true,
		AsContractsField0: &types.PalletContractsPalletCall{
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
func MakeSetCodeCall(dest0 types.MultiAddress, codeHash1 [32]byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsContracts: true,
		AsContractsField0: &types.PalletContractsPalletCall{
			IsSetCode:          true,
			AsSetCodeDest0:     dest0,
			AsSetCodeCodeHash1: codeHash1,
		},
	}
}

// Makes a call to an account, optionally transferring some balance.
//
// # Parameters
//
//   - `dest`: Address of the contract to call.
//   - `value`: The balance to transfer from the `origin` to `dest`.
//   - `gas_limit`: The gas limit enforced when executing the constructor.
//   - `storage_deposit_limit`: The maximum amount of balance that can be charged from the
//     caller to pay for the storage consumed.
//   - `data`: The input data to pass to the contract.
//
// * If the account is a smart-contract account, the associated code will be
// executed and any value will be transferred.
// * If the account is a regular account, any value will be transferred.
// * If no account exists and the call value is not less than `existential_deposit`,
// a regular account will be created and any value will be transferred.
func MakeCallCall(dest0 types.MultiAddress, value1 types1.UCompact, gasLimit2 types.Weight, storageDepositLimit3 types.OptionTUCompact, data4 []byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsContracts: true,
		AsContractsField0: &types.PalletContractsPalletCall{
			IsCall:                     true,
			AsCallDest0:                dest0,
			AsCallValue1:               value1,
			AsCallGasLimit2:            gasLimit2,
			AsCallStorageDepositLimit3: storageDepositLimit3,
			AsCallData4:                data4,
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
//   - `gas_limit`: The gas limit enforced when executing the constructor.
//   - `storage_deposit_limit`: The maximum amount of balance that can be charged/reserved
//     from the caller to pay for the storage consumed.
//   - `code`: The contract code to deploy in raw bytes.
//   - `data`: The input data to pass to the contract constructor.
//   - `salt`: Used for the address derivation. See [`Pallet::contract_address`].
//
// Instantiation is executed as follows:
//
// - The supplied `code` is deployed, and a `code_hash` is created for that code.
// - If the `code_hash` already exists on the chain the underlying `code` will be shared.
// - The destination address is computed based on the sender, code_hash and the salt.
// - The smart-contract account is created at the computed address.
// - The `value` is transferred to the new account.
// - The `deploy` function is executed in the context of the newly-created account.
func MakeInstantiateWithCodeCall(value0 types1.UCompact, gasLimit1 types.Weight, storageDepositLimit2 types.OptionTUCompact, code3 []byte, data4 []byte, salt5 []byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsContracts: true,
		AsContractsField0: &types.PalletContractsPalletCall{
			IsInstantiateWithCode:                     true,
			AsInstantiateWithCodeValue0:               value0,
			AsInstantiateWithCodeGasLimit1:            gasLimit1,
			AsInstantiateWithCodeStorageDepositLimit2: storageDepositLimit2,
			AsInstantiateWithCodeCode3:                code3,
			AsInstantiateWithCodeData4:                data4,
			AsInstantiateWithCodeSalt5:                salt5,
		},
	}
}

// Instantiates a contract from a previously deployed wasm binary.
//
// This function is identical to [`Self::instantiate_with_code`] but without the
// code deployment step. Instead, the `code_hash` of an on-chain deployed wasm binary
// must be supplied.
func MakeInstantiateCall(value0 types1.UCompact, gasLimit1 types.Weight, storageDepositLimit2 types.OptionTUCompact, codeHash3 [32]byte, data4 []byte, salt5 []byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsContracts: true,
		AsContractsField0: &types.PalletContractsPalletCall{
			IsInstantiate:                     true,
			AsInstantiateValue0:               value0,
			AsInstantiateGasLimit1:            gasLimit1,
			AsInstantiateStorageDepositLimit2: storageDepositLimit2,
			AsInstantiateCodeHash3:            codeHash3,
			AsInstantiateData4:                data4,
			AsInstantiateSalt5:                salt5,
		},
	}
}

// When a migration is in progress, this dispatchable can be used to run migration steps.
// Calls that contribute to advancing the migration have their fees waived, as it's helpful
// for the chain. Note that while the migration is in progress, the pallet will also
// leverage the `on_idle` hooks to run migration steps.
func MakeMigrateCall(weightLimit0 types.Weight) types.RuntimeCall {
	return types.RuntimeCall{
		IsContracts: true,
		AsContractsField0: &types.PalletContractsPalletCall{
			IsMigrate:             true,
			AsMigrateWeightLimit0: weightLimit0,
		},
	}
}
