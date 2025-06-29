package subnet

import (
	"errors"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	chain "github.com/wetee-dao/ink.go"
	"github.com/wetee-dao/ink.go/util"
)

type Subnet struct {
	ChainClient *chain.ChainClient
	Address     types.H160
}

func (c *Subnet) Client() *chain.ChainClient {
	return c.ChainClient
}
func (c *Subnet) ContractAddress() types.H160 {
	return c.Address
}

func (c *Subnet) DryRunSetCode(
	code_hash types.H256, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "set_code",
			Args:     []any{code_hash},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	if v != nil && v.IsErr {
		return nil, nil, errors.New("Contract Reverted: " + v.E.Error())
	}

	return v, gas, nil
}

func (c *Subnet) CallSetCode(
	code_hash types.H256, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "set_code",
			Args:     []any{code_hash},
		},
	)
	return err
}

func (c *Subnet) QueryBootNodes(
	params chain.DryRunCallParams,
) (*util.Result[[]SecretNode, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[[]SecretNode, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "boot_nodes",
			Args:     []any{},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	if v != nil && v.IsErr {
		return nil, nil, errors.New("Contract Reverted: " + v.E.Error())
	}

	return v, gas, nil
}

func (c *Subnet) DryRunSetBootNodes(
	nodes []uint64, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "set_boot_nodes",
			Args:     []any{nodes},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	if v != nil && v.IsErr {
		return nil, nil, errors.New("Contract Reverted: " + v.E.Error())
	}

	return v, gas, nil
}

func (c *Subnet) CallSetBootNodes(
	nodes []uint64, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "set_boot_nodes",
			Args:     []any{nodes},
		},
	)
	return err
}

func (c *Subnet) QueryWorkers(
	params chain.DryRunCallParams,
) (*util.Result[[]Tuple_70, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[[]Tuple_70, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "workers",
			Args:     []any{},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	if v != nil && v.IsErr {
		return nil, nil, errors.New("Contract Reverted: " + v.E.Error())
	}

	return v, gas, nil
}

func (c *Subnet) DryRunWorkerRegister(
	name []byte, p2p_id AccountId, ip Ip, port uint32, level byte, params chain.DryRunCallParams,
) (*util.Result[uint64, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[uint64, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "worker_register",
			Args:     []any{name, p2p_id, ip, port, level},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	if v != nil && v.IsErr {
		return nil, nil, errors.New("Contract Reverted: " + v.E.Error())
	}

	return v, gas, nil
}

func (c *Subnet) CallWorkerRegister(
	name []byte, p2p_id AccountId, ip Ip, port uint32, level byte, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "worker_register",
			Args:     []any{name, p2p_id, ip, port, level},
		},
	)
	return err
}

func (c *Subnet) DryRunWorkerMortgage(
	id uint64, cpu uint32, mem uint32, cvm_cpu uint32, cvm_mem uint32, disk uint32, gpu uint32, deposit types.U256, params chain.DryRunCallParams,
) (*util.Result[uint32, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[uint32, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "worker_mortgage",
			Args:     []any{id, cpu, mem, cvm_cpu, cvm_mem, disk, gpu, deposit},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	if v != nil && v.IsErr {
		return nil, nil, errors.New("Contract Reverted: " + v.E.Error())
	}

	return v, gas, nil
}

func (c *Subnet) CallWorkerMortgage(
	id uint64, cpu uint32, mem uint32, cvm_cpu uint32, cvm_mem uint32, disk uint32, gpu uint32, deposit types.U256, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "worker_mortgage",
			Args:     []any{id, cpu, mem, cvm_cpu, cvm_mem, disk, gpu, deposit},
		},
	)
	return err
}

func (c *Subnet) DryRunWorkerUnmortgage(
	id uint64, mortgage_id uint32, params chain.DryRunCallParams,
) (*util.Result[uint32, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[uint32, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "worker_unmortgage",
			Args:     []any{id, mortgage_id},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	if v != nil && v.IsErr {
		return nil, nil, errors.New("Contract Reverted: " + v.E.Error())
	}

	return v, gas, nil
}

func (c *Subnet) CallWorkerUnmortgage(
	id uint64, mortgage_id uint32, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "worker_unmortgage",
			Args:     []any{id, mortgage_id},
		},
	)
	return err
}

func (c *Subnet) DryRunWorkerStop(
	id uint64, params chain.DryRunCallParams,
) (*util.Result[uint64, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[uint64, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "worker_stop",
			Args:     []any{id},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	if v != nil && v.IsErr {
		return nil, nil, errors.New("Contract Reverted: " + v.E.Error())
	}

	return v, gas, nil
}

func (c *Subnet) CallWorkerStop(
	id uint64, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "worker_stop",
			Args:     []any{id},
		},
	)
	return err
}

func (c *Subnet) QuerySecrets(
	params chain.DryRunCallParams,
) (*util.Result[[]Tuple_78, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[[]Tuple_78, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "secrets",
			Args:     []any{},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	if v != nil && v.IsErr {
		return nil, nil, errors.New("Contract Reverted: " + v.E.Error())
	}

	return v, gas, nil
}

func (c *Subnet) DryRunSecretRegister(
	name []byte, validator_id AccountId, p2p_id AccountId, ip Ip, port uint32, params chain.DryRunCallParams,
) (*util.Result[uint64, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[uint64, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "secret_register",
			Args:     []any{name, validator_id, p2p_id, ip, port},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	if v != nil && v.IsErr {
		return nil, nil, errors.New("Contract Reverted: " + v.E.Error())
	}

	return v, gas, nil
}

func (c *Subnet) CallSecretRegister(
	name []byte, validator_id AccountId, p2p_id AccountId, ip Ip, port uint32, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "secret_register",
			Args:     []any{name, validator_id, p2p_id, ip, port},
		},
	)
	return err
}

func (c *Subnet) DryRunSecretDeposit(
	id uint64, deposit types.U256, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "secret_deposit",
			Args:     []any{id, deposit},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	if v != nil && v.IsErr {
		return nil, nil, errors.New("Contract Reverted: " + v.E.Error())
	}

	return v, gas, nil
}

func (c *Subnet) CallSecretDeposit(
	id uint64, deposit types.U256, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "secret_deposit",
			Args:     []any{id, deposit},
		},
	)
	return err
}

func (c *Subnet) DryRunSecretJoin(
	id uint64, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "secret_join",
			Args:     []any{id},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	if v != nil && v.IsErr {
		return nil, nil, errors.New("Contract Reverted: " + v.E.Error())
	}

	return v, gas, nil
}

func (c *Subnet) CallSecretJoin(
	id uint64, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "secret_join",
			Args:     []any{id},
		},
	)
	return err
}

func (c *Subnet) DryRunSecretDelete(
	id uint64, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "secret_delete",
			Args:     []any{id},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	if v != nil && v.IsErr {
		return nil, nil, errors.New("Contract Reverted: " + v.E.Error())
	}

	return v, gas, nil
}

func (c *Subnet) CallSecretDelete(
	id uint64, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "secret_delete",
			Args:     []any{id},
		},
	)
	return err
}

func (c *Subnet) QueryEpoch(
	params chain.DryRunCallParams,
) (*Tuple_80, *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[Tuple_80](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "epoch",
			Args:     []any{},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	return v, gas, nil
}

func (c *Subnet) DryRunNextEpoch(
	new_key [32]byte, sig [64]byte, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "next_epoch",
			Args:     []any{new_key, sig},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	if v != nil && v.IsErr {
		return nil, nil, errors.New("Contract Reverted: " + v.E.Error())
	}

	return v, gas, nil
}

func (c *Subnet) CallNextEpoch(
	new_key [32]byte, sig [64]byte, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "next_epoch",
			Args:     []any{new_key, sig},
		},
	)
	return err
}

func (c *Subnet) DryRunNextEpochWithGov(
	params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "next_epoch_with_gov",
			Args:     []any{},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	if v != nil && v.IsErr {
		return nil, nil, errors.New("Contract Reverted: " + v.E.Error())
	}

	return v, gas, nil
}

func (c *Subnet) CallNextEpochWithGov(
	params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "next_epoch_with_gov",
			Args:     []any{},
		},
	)
	return err
}
