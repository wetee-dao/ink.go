package client

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/extrinsic"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/extrinsic/extensions"
)

func NewExtrinsic(c types.Call) Extrinsic {
	return Extrinsic{
		extrinsic.Extrinsic{
			Version: extrinsic.Version4,
			Method:  c,
		},
	}
}

type Extrinsic struct {
	extrinsic.Extrinsic
}

func (e *Extrinsic) Sign(signer SignerType, meta *types.Metadata, opts ...extrinsic.SigningOption) error {
	if e.Type() != extrinsic.Version4 {
		return extrinsic.ErrInvalidVersion.WithMsg("unsupported extrinsic version: %v (isSigned: %v, type: %v)", e.Version, e.IsSigned(), e.Type())
	}

	encodedMethod, err := codec.Encode(e.Method)
	if err != nil {
		return err
	}

	fieldValues := extrinsic.SignedFieldValues{}
	for _, opt := range opts {
		opt(fieldValues)
	}

	payload, err := createPayload(meta, encodedMethod)
	if err != nil {
		return extrinsic.ErrPayloadCreation.Wrap(err)
	}

	if err := payload.MutateSignedFields(fieldValues); err != nil {
		return extrinsic.ErrPayloadMutation.Wrap(err)
	}

	signerPubKey, err := types.NewMultiAddressFromAccountID(signer.Public())
	if err != nil {
		return err
	}

	sig, err := PayloadSign(signer, payload) //payload.Sign(signer)
	if err != nil {
		return extrinsic.ErrPayloadSigning.Wrap(err)
	}

	var signature types.MultiSignature
	switch signer.SignType() {
	case 0:
		signature = types.MultiSignature{IsSr25519: true, AsSr25519: sig}
	case 1:
		signature = types.MultiSignature{IsEd25519: true, AsEd25519: sig}
	}

	extSignature := &extrinsic.Signature{
		Signer:       signerPubKey,
		Signature:    signature,
		SignedFields: payload.SignedFields,
	}

	e.Signature = extSignature

	// mark the extrinsic as signed
	e.Version |= extrinsic.BitSigned

	return nil
}

func createPayload(meta *types.Metadata, encodedCall []byte) (*extrinsic.Payload, error) {
	payload := &extrinsic.Payload{
		EncodedCall: encodedCall,
	}

	for _, signedExtension := range meta.AsMetadataV14.Extrinsic.SignedExtensions {
		signedExtensionType, ok := meta.AsMetadataV14.EfficientLookup[signedExtension.Type.Int64()]

		if !ok {
			return nil, extrinsic.ErrSignedExtensionTypeNotDefined.WithMsg("lookup ID - '%d'", signedExtension.Type.Int64())
		}

		signedExtensionName := extensions.SignedExtensionName(signedExtensionType.Path[len(signedExtensionType.Path)-1])

		// TODO pull requset to go-substrate-rpc-client
		if signedExtensionName == "WeightReclaim" {
			continue
		}

		payloadMutatorFn, ok := extrinsic.PayloadMutatorFns[signedExtensionName]

		if !ok {
			return nil, extrinsic.ErrSignedExtensionTypeNotSupported.WithMsg("signed extension '%s'", signedExtensionName)
		}

		payloadMutatorFn(payload)
	}

	return payload, nil
}

func PayloadSign(signer SignerType, p *extrinsic.Payload) (sig types.SignatureHash, err error) {
	b, err := codec.Encode(p)
	if err != nil {
		return sig, extrinsic.ErrPayloadEncoding.Wrap(err)
	}

	signatureBytes, err := signer.Sign(b)
	if err != nil {
		return sig, extrinsic.ErrPayloadSigning.Wrap(err)
	}

	sig = types.NewSignature(signatureBytes)

	return sig, nil
}
