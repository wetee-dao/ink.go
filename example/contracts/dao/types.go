package dao

import (
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/wetee-dao/ink.go/util"
)

type Call struct { // Composite
	Contract     util.Option[types.H160]
	Selector     [4]byte
	Input        []byte
	Amount       types.U256
	RefTimeLimit uint64
	AllowReentry bool
}
type PropStatus struct { // Enum
	Pending    *bool   // 0
	Ongoing    *bool   // 1
	Confirming *bool   // 2
	Approved   *uint32 // 3
	Rejected   *uint32 // 4
	Canceled   *bool   // 5
}

func (ty PropStatus) Encode(encoder scale.Encoder) (err error) {
	if ty.Pending != nil {
		err = encoder.PushByte(0)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.Ongoing != nil {
		err = encoder.PushByte(1)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.Confirming != nil {
		err = encoder.PushByte(2)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.Approved != nil {
		err = encoder.PushByte(3)
		if err != nil {
			return err
		}
		err = encoder.Encode(*ty.Approved)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.Rejected != nil {
		err = encoder.PushByte(4)
		if err != nil {
			return err
		}
		err = encoder.Encode(*ty.Rejected)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.Canceled != nil {
		err = encoder.PushByte(5)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("unrecognized enum")
}

func (ty *PropStatus) Decode(decoder scale.Decoder) (err error) {
	variant, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch variant {
	case 0:
		t := true
		ty.Pending = &t
		return
	case 1:
		t := true
		ty.Ongoing = &t
		return
	case 2:
		t := true
		ty.Confirming = &t
		return
	case 3:
		err = decoder.Decode(ty.Approved)
		if err != nil {
			return err
		}
		return
	case 4:
		err = decoder.Decode(ty.Rejected)
		if err != nil {
			return err
		}
		return
	case 5:
		t := true
		ty.Canceled = &t
		return
	default:
		return fmt.Errorf("unrecognized enum")
	}
}

type Track struct { // Composite
	Name               []byte
	PreparePeriod      uint32
	DecisionDeposit    types.U256
	MaxDeciding        uint32
	ConfirmPeriod      uint32
	DecisionPeriod     uint32
	MinEnactmentPeriod uint32
	MaxBalance         types.U256
	MinApproval        Curve
	MinSupport         Curve
}
type Curve struct { // Enum
	LinearDecreasing *struct { // 0
		Begin  uint32
		End    uint32
		Length uint32
	}
	SteppedDecreasing *struct { // 1
		Begin  uint32
		End    uint32
		Step   uint32
		Period uint32
	}
	Reciprocal *struct { // 2
		Factor  uint32
		XScale  uint32
		XOffset int64
		YOffset int64
	}
}

func (ty Curve) Encode(encoder scale.Encoder) (err error) {
	if ty.LinearDecreasing != nil {
		err = encoder.PushByte(0)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.LinearDecreasing.Begin)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.LinearDecreasing.End)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.LinearDecreasing.Length)
		if err != nil {
			return err
		}

		return nil
	}

	if ty.SteppedDecreasing != nil {
		err = encoder.PushByte(1)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.SteppedDecreasing.Begin)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.SteppedDecreasing.End)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.SteppedDecreasing.Step)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.SteppedDecreasing.Period)
		if err != nil {
			return err
		}

		return nil
	}

	if ty.Reciprocal != nil {
		err = encoder.PushByte(2)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.Reciprocal.Factor)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.Reciprocal.XScale)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.Reciprocal.XOffset)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.Reciprocal.YOffset)
		if err != nil {
			return err
		}

		return nil
	}
	return fmt.Errorf("unrecognized enum")
}

func (ty *Curve) Decode(decoder scale.Decoder) (err error) {
	variant, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch variant {
	case 0:
		ty.LinearDecreasing = &struct {
			Begin  uint32
			End    uint32
			Length uint32
		}{}

		err = decoder.Decode(&ty.LinearDecreasing.Begin)
		if err != nil {
			return err
		}

		err = decoder.Decode(&ty.LinearDecreasing.End)
		if err != nil {
			return err
		}

		err = decoder.Decode(&ty.LinearDecreasing.Length)
		if err != nil {
			return err
		}

		return
	case 1:
		ty.SteppedDecreasing = &struct {
			Begin  uint32
			End    uint32
			Step   uint32
			Period uint32
		}{}

		err = decoder.Decode(&ty.SteppedDecreasing.Begin)
		if err != nil {
			return err
		}

		err = decoder.Decode(&ty.SteppedDecreasing.End)
		if err != nil {
			return err
		}

		err = decoder.Decode(&ty.SteppedDecreasing.Step)
		if err != nil {
			return err
		}

		err = decoder.Decode(&ty.SteppedDecreasing.Period)
		if err != nil {
			return err
		}

		return
	case 2:
		ty.Reciprocal = &struct {
			Factor  uint32
			XScale  uint32
			XOffset int64
			YOffset int64
		}{}

		err = decoder.Decode(&ty.Reciprocal.Factor)
		if err != nil {
			return err
		}

		err = decoder.Decode(&ty.Reciprocal.XScale)
		if err != nil {
			return err
		}

		err = decoder.Decode(&ty.Reciprocal.XOffset)
		if err != nil {
			return err
		}

		err = decoder.Decode(&ty.Reciprocal.YOffset)
		if err != nil {
			return err
		}

		return
	default:
		return fmt.Errorf("unrecognized enum")
	}
}

type VoteInfo struct { // Composite
	Pledge      types.U256
	Opinion     Opinion
	VoteWeight  types.U256
	UnlockBlock uint32
	CallId      uint32
	Calller     types.H160
	VoteBlock   uint32
	Deleted     bool
}
type Opinion struct { // Enum
	YES *bool // 0
	NO  *bool // 1
}

func (ty Opinion) Encode(encoder scale.Encoder) (err error) {
	if ty.YES != nil {
		err = encoder.PushByte(0)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.NO != nil {
		err = encoder.PushByte(1)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("unrecognized enum")
}

func (ty *Opinion) Decode(decoder scale.Decoder) (err error) {
	variant, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch variant {
	case 0:
		t := true
		ty.YES = &t
		return
	case 1:
		t := true
		ty.NO = &t
		return
	default:
		return fmt.Errorf("unrecognized enum")
	}
}

type Error struct { // Enum
	MemberExisted         *bool // 0
	MemberNotExisted      *bool // 1
	MemberBalanceNotZero  *bool // 2
	PublicJoinNotAllowed  *bool // 3
	LowBalance            *bool // 4
	CallFailed            *bool // 5
	InvalidDeposit        *bool // 6
	TransferFailed        *bool // 7
	MustCallByGov         *bool // 8
	PropNotOngoing        *bool // 9
	PropNotEnd            *bool // 10
	InvalidProposal       *bool // 11
	InvalidProposalStatus *bool // 12
	InvalidProposalCaller *bool // 13
	InvalidDepositTime    *bool // 14
	InvalidVoteTime       *bool // 15
	InvalidVoteStatus     *bool // 16
	InvalidVoteUser       *bool // 17
	ProposalInDecision    *bool // 18
	VoteAlreadyUnlocked   *bool // 19
	InvalidVoteUnlockTime *bool // 20
	ProposalNotConfirmed  *bool // 21
	NoTrack               *bool // 22
	MaxBalanceOverflow    *bool // 23
	TransferDisable       *bool // 24
	InvalidVote           *bool // 25
	SetCodeFailed         *bool // 26
	SpendNotFound         *bool // 27
	SpendAlreadyExecuted  *bool // 28
	SpendTransferError    *bool // 29
}

func (ty Error) Encode(encoder scale.Encoder) (err error) {
	if ty.MemberExisted != nil {
		err = encoder.PushByte(0)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.MemberNotExisted != nil {
		err = encoder.PushByte(1)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.MemberBalanceNotZero != nil {
		err = encoder.PushByte(2)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.PublicJoinNotAllowed != nil {
		err = encoder.PushByte(3)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.LowBalance != nil {
		err = encoder.PushByte(4)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.CallFailed != nil {
		err = encoder.PushByte(5)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.InvalidDeposit != nil {
		err = encoder.PushByte(6)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.TransferFailed != nil {
		err = encoder.PushByte(7)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.MustCallByGov != nil {
		err = encoder.PushByte(8)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.PropNotOngoing != nil {
		err = encoder.PushByte(9)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.PropNotEnd != nil {
		err = encoder.PushByte(10)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.InvalidProposal != nil {
		err = encoder.PushByte(11)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.InvalidProposalStatus != nil {
		err = encoder.PushByte(12)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.InvalidProposalCaller != nil {
		err = encoder.PushByte(13)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.InvalidDepositTime != nil {
		err = encoder.PushByte(14)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.InvalidVoteTime != nil {
		err = encoder.PushByte(15)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.InvalidVoteStatus != nil {
		err = encoder.PushByte(16)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.InvalidVoteUser != nil {
		err = encoder.PushByte(17)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.ProposalInDecision != nil {
		err = encoder.PushByte(18)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.VoteAlreadyUnlocked != nil {
		err = encoder.PushByte(19)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.InvalidVoteUnlockTime != nil {
		err = encoder.PushByte(20)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.ProposalNotConfirmed != nil {
		err = encoder.PushByte(21)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.NoTrack != nil {
		err = encoder.PushByte(22)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.MaxBalanceOverflow != nil {
		err = encoder.PushByte(23)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.TransferDisable != nil {
		err = encoder.PushByte(24)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.InvalidVote != nil {
		err = encoder.PushByte(25)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.SetCodeFailed != nil {
		err = encoder.PushByte(26)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.SpendNotFound != nil {
		err = encoder.PushByte(27)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.SpendAlreadyExecuted != nil {
		err = encoder.PushByte(28)
		if err != nil {
			return err
		}
		return nil
	}

	if ty.SpendTransferError != nil {
		err = encoder.PushByte(29)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("unrecognized enum")
}

func (ty *Error) Decode(decoder scale.Decoder) (err error) {
	variant, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch variant {
	case 0:
		t := true
		ty.MemberExisted = &t
		return
	case 1:
		t := true
		ty.MemberNotExisted = &t
		return
	case 2:
		t := true
		ty.MemberBalanceNotZero = &t
		return
	case 3:
		t := true
		ty.PublicJoinNotAllowed = &t
		return
	case 4:
		t := true
		ty.LowBalance = &t
		return
	case 5:
		t := true
		ty.CallFailed = &t
		return
	case 6:
		t := true
		ty.InvalidDeposit = &t
		return
	case 7:
		t := true
		ty.TransferFailed = &t
		return
	case 8:
		t := true
		ty.MustCallByGov = &t
		return
	case 9:
		t := true
		ty.PropNotOngoing = &t
		return
	case 10:
		t := true
		ty.PropNotEnd = &t
		return
	case 11:
		t := true
		ty.InvalidProposal = &t
		return
	case 12:
		t := true
		ty.InvalidProposalStatus = &t
		return
	case 13:
		t := true
		ty.InvalidProposalCaller = &t
		return
	case 14:
		t := true
		ty.InvalidDepositTime = &t
		return
	case 15:
		t := true
		ty.InvalidVoteTime = &t
		return
	case 16:
		t := true
		ty.InvalidVoteStatus = &t
		return
	case 17:
		t := true
		ty.InvalidVoteUser = &t
		return
	case 18:
		t := true
		ty.ProposalInDecision = &t
		return
	case 19:
		t := true
		ty.VoteAlreadyUnlocked = &t
		return
	case 20:
		t := true
		ty.InvalidVoteUnlockTime = &t
		return
	case 21:
		t := true
		ty.ProposalNotConfirmed = &t
		return
	case 22:
		t := true
		ty.NoTrack = &t
		return
	case 23:
		t := true
		ty.MaxBalanceOverflow = &t
		return
	case 24:
		t := true
		ty.TransferDisable = &t
		return
	case 25:
		t := true
		ty.InvalidVote = &t
		return
	case 26:
		t := true
		ty.SetCodeFailed = &t
		return
	case 27:
		t := true
		ty.SpendNotFound = &t
		return
	case 28:
		t := true
		ty.SpendAlreadyExecuted = &t
		return
	case 29:
		t := true
		ty.SpendTransferError = &t
		return
	default:
		return fmt.Errorf("unrecognized enum")
	}
}
func (ty *Error) Error() string {
	if ty.MemberExisted != nil {
		return "MemberExisted"
	}

	if ty.MemberNotExisted != nil {
		return "MemberNotExisted"
	}

	if ty.MemberBalanceNotZero != nil {
		return "MemberBalanceNotZero"
	}

	if ty.PublicJoinNotAllowed != nil {
		return "PublicJoinNotAllowed"
	}

	if ty.LowBalance != nil {
		return "LowBalance"
	}

	if ty.CallFailed != nil {
		return "CallFailed"
	}

	if ty.InvalidDeposit != nil {
		return "InvalidDeposit"
	}

	if ty.TransferFailed != nil {
		return "TransferFailed"
	}

	if ty.MustCallByGov != nil {
		return "MustCallByGov"
	}

	if ty.PropNotOngoing != nil {
		return "PropNotOngoing"
	}

	if ty.PropNotEnd != nil {
		return "PropNotEnd"
	}

	if ty.InvalidProposal != nil {
		return "InvalidProposal"
	}

	if ty.InvalidProposalStatus != nil {
		return "InvalidProposalStatus"
	}

	if ty.InvalidProposalCaller != nil {
		return "InvalidProposalCaller"
	}

	if ty.InvalidDepositTime != nil {
		return "InvalidDepositTime"
	}

	if ty.InvalidVoteTime != nil {
		return "InvalidVoteTime"
	}

	if ty.InvalidVoteStatus != nil {
		return "InvalidVoteStatus"
	}

	if ty.InvalidVoteUser != nil {
		return "InvalidVoteUser"
	}

	if ty.ProposalInDecision != nil {
		return "ProposalInDecision"
	}

	if ty.VoteAlreadyUnlocked != nil {
		return "VoteAlreadyUnlocked"
	}

	if ty.InvalidVoteUnlockTime != nil {
		return "InvalidVoteUnlockTime"
	}

	if ty.ProposalNotConfirmed != nil {
		return "ProposalNotConfirmed"
	}

	if ty.NoTrack != nil {
		return "NoTrack"
	}

	if ty.MaxBalanceOverflow != nil {
		return "MaxBalanceOverflow"
	}

	if ty.TransferDisable != nil {
		return "TransferDisable"
	}

	if ty.InvalidVote != nil {
		return "InvalidVote"
	}

	if ty.SetCodeFailed != nil {
		return "SetCodeFailed"
	}

	if ty.SpendNotFound != nil {
		return "SpendNotFound"
	}

	if ty.SpendAlreadyExecuted != nil {
		return "SpendAlreadyExecuted"
	}

	if ty.SpendTransferError != nil {
		return "SpendTransferError"
	}
	return "Unknown"
}

type Tuple_98 struct { // Tuple
	F0 types.U256
	F1 types.U256
}
type CurveArg struct { // Enum
	LinearDecreasing *struct { // 0
		Begin  uint32
		End    uint32
		Length uint32
	}
	SteppedDecreasing *struct { // 1
		Begin  uint32
		End    uint32
		Step   uint32
		Period uint32
	}
	Reciprocal *struct { // 2
		XOffsetPercent Percent
		XScaleArg      uint32
		Begin          uint32
		End            uint32
	}
}

func (ty CurveArg) Encode(encoder scale.Encoder) (err error) {
	if ty.LinearDecreasing != nil {
		err = encoder.PushByte(0)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.LinearDecreasing.Begin)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.LinearDecreasing.End)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.LinearDecreasing.Length)
		if err != nil {
			return err
		}

		return nil
	}

	if ty.SteppedDecreasing != nil {
		err = encoder.PushByte(1)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.SteppedDecreasing.Begin)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.SteppedDecreasing.End)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.SteppedDecreasing.Step)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.SteppedDecreasing.Period)
		if err != nil {
			return err
		}

		return nil
	}

	if ty.Reciprocal != nil {
		err = encoder.PushByte(2)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.Reciprocal.XOffsetPercent)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.Reciprocal.XScaleArg)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.Reciprocal.Begin)
		if err != nil {
			return err
		}

		err = encoder.Encode(ty.Reciprocal.End)
		if err != nil {
			return err
		}

		return nil
	}
	return fmt.Errorf("unrecognized enum")
}

func (ty *CurveArg) Decode(decoder scale.Decoder) (err error) {
	variant, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch variant {
	case 0:
		ty.LinearDecreasing = &struct {
			Begin  uint32
			End    uint32
			Length uint32
		}{}

		err = decoder.Decode(&ty.LinearDecreasing.Begin)
		if err != nil {
			return err
		}

		err = decoder.Decode(&ty.LinearDecreasing.End)
		if err != nil {
			return err
		}

		err = decoder.Decode(&ty.LinearDecreasing.Length)
		if err != nil {
			return err
		}

		return
	case 1:
		ty.SteppedDecreasing = &struct {
			Begin  uint32
			End    uint32
			Step   uint32
			Period uint32
		}{}

		err = decoder.Decode(&ty.SteppedDecreasing.Begin)
		if err != nil {
			return err
		}

		err = decoder.Decode(&ty.SteppedDecreasing.End)
		if err != nil {
			return err
		}

		err = decoder.Decode(&ty.SteppedDecreasing.Step)
		if err != nil {
			return err
		}

		err = decoder.Decode(&ty.SteppedDecreasing.Period)
		if err != nil {
			return err
		}

		return
	case 2:
		ty.Reciprocal = &struct {
			XOffsetPercent Percent
			XScaleArg      uint32
			Begin          uint32
			End            uint32
		}{}

		err = decoder.Decode(&ty.Reciprocal.XOffsetPercent)
		if err != nil {
			return err
		}

		err = decoder.Decode(&ty.Reciprocal.XScaleArg)
		if err != nil {
			return err
		}

		err = decoder.Decode(&ty.Reciprocal.Begin)
		if err != nil {
			return err
		}

		err = decoder.Decode(&ty.Reciprocal.End)
		if err != nil {
			return err
		}

		return
	default:
		return fmt.Errorf("unrecognized enum")
	}
}

type Percent struct { // Composite
	V uint32
}
