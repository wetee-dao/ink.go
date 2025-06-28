package dao

import (
	"errors"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	chain "github.com/wetee-dao/ink.go"
	"github.com/wetee-dao/ink.go/util"
)

type Dao struct {
	ChainClient *chain.ChainClient
	Address     types.H160
}

func (c *Dao) Client() *chain.ChainClient {
	return c.ChainClient
}
func (c *Dao) ContractAddress() types.H160 {
	return c.Address
}

func (c *Dao) DryRunSetCode(
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

func (c *Dao) CallSetCode(
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

func (c *Dao) QueryMemberList(
	params chain.DryRunCallParams,
) (*[]types.H160, *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[[]types.H160](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Member::list",
			Args:     []any{},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	return v, gas, nil
}

func (c *Dao) QueryMemberGetPublicJoin(
	params chain.DryRunCallParams,
) (*bool, *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[bool](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Member::get_public_join",
			Args:     []any{},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	return v, gas, nil
}

func (c *Dao) DryRunMemberPublicJoin(
	params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Member::public_join",
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

func (c *Dao) CallMemberPublicJoin(
	params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Member::public_join",
			Args:     []any{},
		},
	)
	return err
}

func (c *Dao) DryRunMemberSetPublicJoin(
	public_join bool, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Member::set_public_join",
			Args:     []any{public_join},
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

func (c *Dao) CallMemberSetPublicJoin(
	public_join bool, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Member::set_public_join",
			Args:     []any{public_join},
		},
	)
	return err
}

func (c *Dao) DryRunMemberJoin(
	new_user types.H160, balance types.U256, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Member::join",
			Args:     []any{new_user, balance},
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

func (c *Dao) CallMemberJoin(
	new_user types.H160, balance types.U256, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Member::join",
			Args:     []any{new_user, balance},
		},
	)
	return err
}

func (c *Dao) DryRunMemberLevae(
	params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Member::levae",
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

func (c *Dao) CallMemberLevae(
	params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Member::levae",
			Args:     []any{},
		},
	)
	return err
}

func (c *Dao) DryRunMemberLevaeWithBurn(
	params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Member::levae_with_burn",
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

func (c *Dao) CallMemberLevaeWithBurn(
	params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Member::levae_with_burn",
			Args:     []any{},
		},
	)
	return err
}

func (c *Dao) DryRunMemberDelete(
	user types.H160, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Member::delete",
			Args:     []any{user},
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

func (c *Dao) CallMemberDelete(
	user types.H160, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Member::delete",
			Args:     []any{user},
		},
	)
	return err
}

func (c *Dao) QueryErc20BalanceOf(
	user types.H160, params chain.DryRunCallParams,
) (*Tuple_98, *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[Tuple_98](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Erc20::balance_of",
			Args:     []any{user},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	return v, gas, nil
}

func (c *Dao) DryRunErc20EnableTransfer(
	params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Erc20::enable_transfer",
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

func (c *Dao) CallErc20EnableTransfer(
	params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Erc20::enable_transfer",
			Args:     []any{},
		},
	)
	return err
}

func (c *Dao) DryRunErc20Transfer(
	to types.H160, amount types.U256, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Erc20::transfer",
			Args:     []any{to, amount},
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

func (c *Dao) CallErc20Transfer(
	to types.H160, amount types.U256, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Erc20::transfer",
			Args:     []any{to, amount},
		},
	)
	return err
}

func (c *Dao) DryRunErc20Burn(
	amount types.U256, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Erc20::burn",
			Args:     []any{amount},
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

func (c *Dao) CallErc20Burn(
	amount types.U256, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Erc20::burn",
			Args:     []any{amount},
		},
	)
	return err
}

func (c *Dao) DryRunSudoSudo(
	call Call, params chain.DryRunCallParams,
) (*util.Result[[]byte, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[[]byte, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Sudo::sudo",
			Args:     []any{call},
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

func (c *Dao) CallSudoSudo(
	call Call, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Sudo::sudo",
			Args:     []any{call},
		},
	)
	return err
}

func (c *Dao) DryRunSudoRemoveSudo(
	params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Sudo::remove_sudo",
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

func (c *Dao) CallSudoRemoveSudo(
	params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Sudo::remove_sudo",
			Args:     []any{},
		},
	)
	return err
}

func (c *Dao) DryRunGovSetDefalutTrack(
	id uint16, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::set_defalut_track",
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

func (c *Dao) CallGovSetDefalutTrack(
	id uint16, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::set_defalut_track",
			Args:     []any{id},
		},
	)
	return err
}

func (c *Dao) QueryGovDefalutTrack(
	params chain.DryRunCallParams,
) (*util.Option[uint16], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Option[uint16]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::defalut_track",
			Args:     []any{},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	return v, gas, nil
}

func (c *Dao) QueryGovTrackList(
	page uint16, size uint16, params chain.DryRunCallParams,
) (*[]Track, *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[[]Track](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::track_list",
			Args:     []any{page, size},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	return v, gas, nil
}

func (c *Dao) QueryGovTrack(
	id uint16, params chain.DryRunCallParams,
) (*util.Option[Track], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Option[Track]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::track",
			Args:     []any{id},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	return v, gas, nil
}

func (c *Dao) DryRunGovAddTrack(
	name []byte, prepare_period uint32, decision_deposit types.U256, max_deciding uint32, confirm_period uint32, decision_period uint32, min_enactment_period uint32, max_balance types.U256, min_approval CurveArg, min_support CurveArg, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::add_track",
			Args:     []any{name, prepare_period, decision_deposit, max_deciding, confirm_period, decision_period, min_enactment_period, max_balance, min_approval, min_support},
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

func (c *Dao) CallGovAddTrack(
	name []byte, prepare_period uint32, decision_deposit types.U256, max_deciding uint32, confirm_period uint32, decision_period uint32, min_enactment_period uint32, max_balance types.U256, min_approval CurveArg, min_support CurveArg, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::add_track",
			Args:     []any{name, prepare_period, decision_deposit, max_deciding, confirm_period, decision_period, min_enactment_period, max_balance, min_approval, min_support},
		},
	)
	return err
}

func (c *Dao) DryRunGovEditTrack(
	id uint16, name []byte, prepare_period uint32, decision_deposit types.U256, max_deciding uint32, confirm_period uint32, decision_period uint32, min_enactment_period uint32, max_balance types.U256, min_approval CurveArg, min_support CurveArg, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::edit_track",
			Args:     []any{id, name, prepare_period, decision_deposit, max_deciding, confirm_period, decision_period, min_enactment_period, max_balance, min_approval, min_support},
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

func (c *Dao) CallGovEditTrack(
	id uint16, name []byte, prepare_period uint32, decision_deposit types.U256, max_deciding uint32, confirm_period uint32, decision_period uint32, min_enactment_period uint32, max_balance types.U256, min_approval CurveArg, min_support CurveArg, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::edit_track",
			Args:     []any{id, name, prepare_period, decision_deposit, max_deciding, confirm_period, decision_period, min_enactment_period, max_balance, min_approval, min_support},
		},
	)
	return err
}

func (c *Dao) QueryGovProposals(
	page uint16, size uint16, params chain.DryRunCallParams,
) (*[]Call, *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[[]Call](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::proposals",
			Args:     []any{page, size},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	return v, gas, nil
}

func (c *Dao) QueryGovProposal(
	id uint32, params chain.DryRunCallParams,
) (*util.Option[Call], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Option[Call]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::proposal",
			Args:     []any{id},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	return v, gas, nil
}

func (c *Dao) DryRunGovSubmitProposal(
	call Call, track_id uint16, params chain.DryRunCallParams,
) (*util.Result[uint32, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[uint32, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::submit_proposal",
			Args:     []any{call, track_id},
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

func (c *Dao) CallGovSubmitProposal(
	call Call, track_id uint16, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::submit_proposal",
			Args:     []any{call, track_id},
		},
	)
	return err
}

func (c *Dao) DryRunGovCancelProposal(
	proposal_id uint32, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::cancel_proposal",
			Args:     []any{proposal_id},
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

func (c *Dao) CallGovCancelProposal(
	proposal_id uint32, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::cancel_proposal",
			Args:     []any{proposal_id},
		},
	)
	return err
}

func (c *Dao) DryRunGovDepositProposal(
	proposal_id uint32, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::deposit_proposal",
			Args:     []any{proposal_id},
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

func (c *Dao) CallGovDepositProposal(
	proposal_id uint32, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::deposit_proposal",
			Args:     []any{proposal_id},
		},
	)
	return err
}

func (c *Dao) QueryGovVoteList(
	proposal_id uint32, params chain.DryRunCallParams,
) (*[]VoteInfo, *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[[]VoteInfo](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::vote_list",
			Args:     []any{proposal_id},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	return v, gas, nil
}

func (c *Dao) DryRunGovVote(
	vote_id types.U128, params chain.DryRunCallParams,
) (*util.Option[VoteInfo], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Option[VoteInfo]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::vote",
			Args:     []any{vote_id},
		},
	)
	if err != nil && !errors.Is(err, chain.ErrContractReverted) {
		return nil, nil, err
	}
	return v, gas, nil
}

func (c *Dao) CallGovVote(
	vote_id types.U128, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::vote",
			Args:     []any{vote_id},
		},
	)
	return err
}

func (c *Dao) DryRunGovSubmitVote(
	proposal_id uint32, opinion Opinion, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::submit_vote",
			Args:     []any{proposal_id, opinion},
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

func (c *Dao) CallGovSubmitVote(
	proposal_id uint32, opinion Opinion, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::submit_vote",
			Args:     []any{proposal_id, opinion},
		},
	)
	return err
}

func (c *Dao) DryRunGovCancelVote(
	vote_id types.U128, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::cancel_vote",
			Args:     []any{vote_id},
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

func (c *Dao) CallGovCancelVote(
	vote_id types.U128, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::cancel_vote",
			Args:     []any{vote_id},
		},
	)
	return err
}

func (c *Dao) DryRunGovUnlock(
	vote_id types.U128, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::unlock",
			Args:     []any{vote_id},
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

func (c *Dao) CallGovUnlock(
	vote_id types.U128, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::unlock",
			Args:     []any{vote_id},
		},
	)
	return err
}

func (c *Dao) DryRunGovExecProposal(
	proposal_id uint32, params chain.DryRunCallParams,
) (*util.Result[[]byte, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[[]byte, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::exec_proposal",
			Args:     []any{proposal_id},
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

func (c *Dao) CallGovExecProposal(
	proposal_id uint32, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::exec_proposal",
			Args:     []any{proposal_id},
		},
	)
	return err
}

func (c *Dao) QueryGovProposalStatus(
	proposal_id uint32, params chain.DryRunCallParams,
) (*util.Result[PropStatus, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[PropStatus, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Gov::proposal_status",
			Args:     []any{proposal_id},
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

func (c *Dao) DryRunTreasurySpend(
	track_id uint16, to types.H160, _assert_id uint64, amount types.U256, params chain.DryRunCallParams,
) (*util.Result[uint64, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[uint64, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Treasury::spend",
			Args:     []any{track_id, to, _assert_id, amount},
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

func (c *Dao) CallTreasurySpend(
	track_id uint16, to types.H160, _assert_id uint64, amount types.U256, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Treasury::spend",
			Args:     []any{track_id, to, _assert_id, amount},
		},
	)
	return err
}

func (c *Dao) DryRunTreasuryPayout(
	spend_index uint64, params chain.DryRunCallParams,
) (*util.Result[util.NullTuple, Error], *chain.DryRunReturnGas, error) {
	v, gas, err := chain.DryRun[util.Result[util.NullTuple, Error]](
		c,
		params.Origin,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Treasury::payout",
			Args:     []any{spend_index},
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

func (c *Dao) CallTreasuryPayout(
	spend_index uint64, params chain.CallParams,
) error {
	err := chain.Call(
		c,
		params.Signer,
		params.PayAmount,
		params.GasLimit,
		params.StorageDepositLimit,
		util.InkContractInput{
			Selector: "Treasury::payout",
			Args:     []any{spend_index},
		},
	)
	return err
}
