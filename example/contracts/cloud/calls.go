package cloud

import (
	"errors"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	chain "github.com/wetee-dao/ink.go"
	"github.com/wetee-dao/ink.go/util"
)

func InitCloudContract(client *chain.ChainClient, address string) (chain.Ink, error) {
	contractAddress, err := util.HexToH160(address)
	if err != nil {
		return nil, err
	}
	return &Cloud{
		ChainClient: client,
		Address:     contractAddress,
	}, nil
}

type Cloud struct {
	ChainClient *chain.ChainClient
	Address     types.H160
}

func (c *Cloud) Client() *chain.ChainClient {
	return c.ChainClient
}
func (c *Cloud) ContractAddress() types.H160 {
	return c.Address
}

func (c *Cloud) DryRunSetCode(
	code_hash types.H256, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRunInk[util.Result[util.NullTuple, Error]](
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

func (c *Cloud) CallSetCode(
	code_hash types.H256, __ink_params chain.CallParams,
) error {
	_param := chain.DefaultParamWithOrigin(__ink_params.Signer.AccountID())
	_param.PayAmount = __ink_params.PayAmount
	_, gas, err := c.DryRunSetCode(code_hash, _param)
	if err != nil {
		return err
	}
	return chain.CallInk(
		c,
		__ink_params.Signer,
		__ink_params.PayAmount,
		gas.GasRequired,
		gas.StorageDeposit,
		util.InkContractInput{
			Selector: "set_code",
			Args:     []any{code_hash},
		},
	)
}

func (c *Cloud) TxCallOfSetCode(
	code_hash types.H256, __ink_params chain.CallParams,
) (*types.Call, error) {
	_param := chain.DefaultParamWithOrigin(__ink_params.Signer.AccountID())
	_param.PayAmount = __ink_params.PayAmount
	_, gas, err := c.DryRunSetCode(code_hash, _param)
	if err != nil {
		return nil, err
	}
	return chain.TxCall(
		c,
		__ink_params.Signer,
		__ink_params.PayAmount,
		gas.GasRequired,
		gas.StorageDeposit,
		util.InkContractInput{
			Selector: "set_code",
			Args:     []any{code_hash},
		},
	)
}

func (c *Cloud) DryRunPodLen(
	params chain.DryRunCallParams,
) (*util.Result[uint64, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRunInk[util.Result[uint64, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "pod_len",
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

func (c *Cloud) CallPodLen(
	__ink_params chain.CallParams,
) error {
	_param := chain.DefaultParamWithOrigin(__ink_params.Signer.AccountID())
	_param.PayAmount = __ink_params.PayAmount
	_, gas, err := c.DryRunPodLen(_param)
	if err != nil {
		return err
	}
	return chain.CallInk(
		c,
		__ink_params.Signer,
		__ink_params.PayAmount,
		gas.GasRequired,
		gas.StorageDeposit,
		util.InkContractInput{
			Selector: "pod_len",
			Args:     []any{},
		},
	)
}

func (c *Cloud) TxCallOfPodLen(
	__ink_params chain.CallParams,
) (*types.Call, error) {
	_param := chain.DefaultParamWithOrigin(__ink_params.Signer.AccountID())
	_param.PayAmount = __ink_params.PayAmount
	_, gas, err := c.DryRunPodLen(_param)
	if err != nil {
		return nil, err
	}
	return chain.TxCall(
		c,
		__ink_params.Signer,
		__ink_params.PayAmount,
		gas.GasRequired,
		gas.StorageDeposit,
		util.InkContractInput{
			Selector: "pod_len",
			Args:     []any{},
		},
	)
}

func (c *Cloud) DryRunCreateUserPod(
	params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRunInk[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "create_user_pod",
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

func (c *Cloud) CallCreateUserPod(
	__ink_params chain.CallParams,
) error {
	_param := chain.DefaultParamWithOrigin(__ink_params.Signer.AccountID())
	_param.PayAmount = __ink_params.PayAmount
	_, gas, err := c.DryRunCreateUserPod(_param)
	if err != nil {
		return err
	}
	return chain.CallInk(
		c,
		__ink_params.Signer,
		__ink_params.PayAmount,
		gas.GasRequired,
		gas.StorageDeposit,
		util.InkContractInput{
			Selector: "create_user_pod",
			Args:     []any{},
		},
	)
}

func (c *Cloud) TxCallOfCreateUserPod(
	__ink_params chain.CallParams,
) (*types.Call, error) {
	_param := chain.DefaultParamWithOrigin(__ink_params.Signer.AccountID())
	_param.PayAmount = __ink_params.PayAmount
	_, gas, err := c.DryRunCreateUserPod(_param)
	if err != nil {
		return nil, err
	}
	return chain.TxCall(
		c,
		__ink_params.Signer,
		__ink_params.PayAmount,
		gas.GasRequired,
		gas.StorageDeposit,
		util.InkContractInput{
			Selector: "create_user_pod",
			Args:     []any{},
		},
	)
}

func (c *Cloud) QueryUserPods(
	params chain.DryRunCallParams,
) (*[]Tuple_48, *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRunInk[[]Tuple_48](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "user_pods",
			Args:     []any{},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	return v, gas, nil
}
