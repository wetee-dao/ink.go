package client

import (
	"errors"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/extrinsic"
)

// 签名并提交交易
// Sign and submit transaction
func (c *ChainClient) PartialSign(signer PartialSignerType, call types.Call) ([]byte, error) {
	accountInfo, err := c.GetAccount(signer)
	if err != nil {
		return nil, errors.New("GetAccountInfo error: " + err.Error())
	}

	ext := NewExtrinsic(call)
	return ext.PartialSign(signer, c.Meta, extrinsic.WithEra(types.ExtrinsicEra{IsImmortalEra: true}, c.Hash),
		extrinsic.WithNonce(types.NewUCompactFromUInt(uint64(accountInfo.Nonce))),
		extrinsic.WithTip(types.NewUCompactFromUInt(0)),
		extrinsic.WithSpecVersion(c.Runtime.SpecVersion),
		extrinsic.WithTransactionVersion(c.Runtime.TransactionVersion),
		extrinsic.WithGenesisHash(c.Hash),
	)
}

func (e *Extrinsic) PartialSign(signer PartialSignerType, meta *types.Metadata, opts ...extrinsic.SigningOption) ([]byte, error) {
	if e.Type() != extrinsic.Version4 {
		return nil, extrinsic.ErrInvalidVersion.WithMsg("unsupported extrinsic version: %v (isSigned: %v, type: %v)", e.Version, e.IsSigned(), e.Type())
	}

	encodedMethod, err := codec.Encode(e.Method)
	if err != nil {
		return nil, err
	}

	fieldValues := extrinsic.SignedFieldValues{}
	for _, opt := range opts {
		opt(fieldValues)
	}

	payload, err := createPayload(meta, encodedMethod)
	if err != nil {
		return nil, extrinsic.ErrPayloadCreation.Wrap(err)
	}

	if err := payload.MutateSignedFields(fieldValues); err != nil {
		return nil, extrinsic.ErrPayloadMutation.Wrap(err)
	}

	return PayloadPartialSign(signer, payload)
}

func PayloadPartialSign(signer PartialSignerType, p *extrinsic.Payload) (sig []byte, err error) {
	b, err := codec.Encode(p)
	if err != nil {
		return sig, extrinsic.ErrPayloadEncoding.Wrap(err)
	}

	signatureBytes, err := signer.PartialSign(b)
	if err != nil {
		return sig, extrinsic.ErrPayloadSigning.Wrap(err)
	}

	return signatureBytes, nil
}
