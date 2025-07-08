package ink

import (
	"crypto"
	goEd25519 "crypto/ed25519"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/vedhavyas/go-subkey/v2"
	"github.com/vedhavyas/go-subkey/v2/ed25519"
	"github.com/vedhavyas/go-subkey/v2/sr25519"
	"github.com/wetee-dao/ink.go/util"
	"golang.org/x/crypto/blake2b"
)

type SignerType interface {
	Public() []byte
	AccountID() types.AccountID
	Sign([]byte) ([]byte, error)
	Verify([]byte, []byte) bool
	SignType() uint8
}

type PartialSignerType interface {
	Public() []byte
	AccountID() types.AccountID
	Sign([]byte) ([]byte, error)
	Verify([]byte, []byte) bool
	SignType() uint8
	PartialSign(msg []byte) ([]byte, error)
}

// Sr25519 Signer or Ed25519 Signer
type Signer struct {
	subkey.KeyPair
	// Address is an SS58 address
	Address string
	// PublicKey
	PublicKey []byte
	// key type
	KeyType uint8
}

func (e *Signer) SignType() uint8 {
	return e.KeyType
}

func (e *Signer) Sign(msg []byte) ([]byte, error) {
	if len(msg) > 256 {
		h := blake2b.Sum256(msg)
		msg = h[:]
	}

	return e.KeyPair.Sign(msg)
}

func (e *Signer) Verify(msg []byte, signature []byte) bool {
	if len(msg) > 256 {
		h := blake2b.Sum256(msg)
		msg = h[:]
	}

	return e.KeyPair.Verify(msg, signature)
}

func (e *Signer) H160Address() types.H160 {
	h160, _ := util.H160FromPublicKey(e.PublicKey)
	return h160
}

func (e *Signer) AccountID() types.AccountID {
	var bt [32]byte
	copy(bt[:], e.PublicKey)

	return types.AccountID(bt)
}

func NewSr25519Pair() (Signer, error) {
	kyr, err := sr25519.Scheme{}.Generate()
	if err != nil {
		return Signer{}, err
	}

	ss58Address := kyr.SS58Address(42)
	return Signer{
		KeyPair:   kyr,
		Address:   ss58Address,
		PublicKey: kyr.Public(),
	}, nil
}

// Sr25519 PairFromSecret generates a sr25519 key pair from a seed or phrase
func Sr25519PairFromSecret(seedOrPhrase string, network uint16) (Signer, error) {
	scheme := sr25519.Scheme{}
	kyr, err := subkey.DeriveKeyPair(scheme, seedOrPhrase)
	if err != nil {
		return Signer{}, err
	}

	ss58Address := kyr.SS58Address(network)
	return Signer{
		KeyPair:   kyr,
		Address:   ss58Address,
		PublicKey: kyr.Public(),
	}, nil
}

func NewEd25519Pair() (Signer, error) {
	kyr, err := ed25519.Scheme{}.Generate()
	if err != nil {
		return Signer{}, err
	}

	ss58Address := kyr.SS58Address(42)
	return Signer{
		KeyType:   1,
		KeyPair:   kyr,
		Address:   ss58Address,
		PublicKey: kyr.Public(),
	}, nil
}

// Ed25519 PairFromSecret generates a ed25519 key pair from a seed or phrase
func Ed25519PairFromSecret(seedOrPhrase string, network uint16) (Signer, error) {
	scheme := ed25519.Scheme{}
	kyr, err := subkey.DeriveKeyPair(scheme, seedOrPhrase)
	if err != nil {
		return Signer{}, err
	}

	ss58Address := kyr.SS58Address(network)
	return Signer{
		KeyType:   1,
		KeyPair:   kyr,
		Address:   ss58Address,
		PublicKey: kyr.Public(),
	}, nil
}

// Ed25519 PairFromPk generates a ed25519 key pair from golang ed25519.PrivateKey
func Ed25519PairFromPk(pk goEd25519.PrivateKey, network uint16) (Signer, error) {
	pub := pk.Public().(goEd25519.PublicKey)
	kyr := Ed25519Signer{
		pub:    &pub,
		secret: &pk,
	}

	ss58Address := kyr.SS58Address(network)
	return Signer{
		KeyType:   1,
		KeyPair:   &kyr,
		Address:   ss58Address,
		PublicKey: pub,
	}, nil
}

type Ed25519Signer struct {
	pub    *goEd25519.PublicKey
	secret *goEd25519.PrivateKey
}

// Public returns the pub key in bytes.
func (e *Ed25519Signer) Public() []byte {
	return *e.pub
}

// Seed returns the seed for this key
func (kr *Ed25519Signer) Seed() []byte {
	return kr.secret.Seed()
}

// AccountID returns the accountID for this key
func (e *Ed25519Signer) AccountID() []byte {
	return e.Public()
}

// SS58Address returns the Base58 public key with checksum and network identifier.
func (e *Ed25519Signer) SS58Address(network uint16) string {
	return subkey.SS58Encode(e.AccountID(), network)
}

func (e *Ed25519Signer) Sign(msg []byte) ([]byte, error) {
	return e.secret.Sign(nil, msg, crypto.Hash(0))
}

func (e *Ed25519Signer) Verify(msg []byte, signature []byte) bool {
	return goEd25519.Verify(*e.pub, msg, signature)
}
