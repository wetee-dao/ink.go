package client

import (
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
)

func NewExtrinsic(c types.Call) Extrinsic {
	return Extrinsic{
		types.Extrinsic{
			Version: types.ExtrinsicVersion4,
			Method:  c,
		},
	}
}

type Extrinsic struct {
	types.Extrinsic
}

func (e *Extrinsic) Sign(signer *Signer, o types.SignatureOptions) error {
	if e.Type() != types.ExtrinsicVersion4 {
		return fmt.Errorf("unsupported extrinsic version: %v (isSigned: %v, type: %v)", e.Version, e.IsSigned(), e.Type())
	}

	mb, err := codec.Encode(e.Method)
	if err != nil {
		return err
	}

	era := o.Era
	if !o.Era.IsMortalEra {
		era = types.ExtrinsicEra{IsImmortalEra: true}
	}

	payload := ExtrinsicPayloadV4{
		ExtrinsicPayloadV3: types.ExtrinsicPayloadV3{
			Method:      mb,
			Era:         era,
			Nonce:       o.Nonce,
			Tip:         o.Tip,
			SpecVersion: o.SpecVersion,
			GenesisHash: o.GenesisHash,
			BlockHash:   o.BlockHash,
		},
		TransactionVersion: o.TransactionVersion,
	}

	signerPubKey, err := types.NewMultiAddressFromAccountID(signer.Public())
	if err != nil {
		return err
	}

	sig, err := payload.Sign(signer)
	if err != nil {
		return err
	}

	var signature types.MultiSignature
	if signer.KeyType == 0 {
		signature = types.MultiSignature{IsSr25519: true, AsSr25519: sig}
	} else if signer.KeyType == 1 {
		signature = types.MultiSignature{IsEd25519: true, AsEd25519: sig}
	}

	extSig := types.ExtrinsicSignatureV4{
		Signer:    signerPubKey,
		Signature: signature,
		Era:       era,
		Nonce:     o.Nonce,
		Tip:       o.Tip,
	}

	e.Signature = extSig

	// mark the extrinsic as signed
	e.Version |= types.ExtrinsicBitSigned

	return nil
}

type ExtrinsicPayloadV4 struct {
	types.ExtrinsicPayloadV3
	TransactionVersion types.U32
}

// Sign the extrinsic payload with the given derivation path
func (e ExtrinsicPayloadV4) Sign(signer *Signer) (types.Signature, error) {
	b, err := codec.Encode(e)
	if err != nil {
		return types.Signature{}, err
	}

	sig, err := signer.Sign(b)
	return types.NewSignature(sig), err
}

func (e ExtrinsicPayloadV4) Encode(encoder scale.Encoder) error {
	err := encoder.Encode(e.Method)
	if err != nil {
		return err
	}

	err = encoder.Encode(e.Era)
	if err != nil {
		return err
	}

	err = encoder.Encode(e.Nonce)
	if err != nil {
		return err
	}

	err = encoder.Encode(e.Tip)
	if err != nil {
		return err
	}

	err = encoder.Encode(e.SpecVersion)
	if err != nil {
		return err
	}

	err = encoder.Encode(e.TransactionVersion)
	if err != nil {
		return err
	}

	err = encoder.Encode(e.GenesisHash)
	if err != nil {
		return err
	}

	err = encoder.Encode(e.BlockHash)
	if err != nil {
		return err
	}

	return nil
}

// Decode does nothing and always returns an error. ExtrinsicPayloadV4 is only used for encoding, not for decoding
func (e *ExtrinsicPayloadV4) Decode(decoder scale.Decoder) error {
	return fmt.Errorf("decoding of ExtrinsicPayloadV4 is not supported")
}
