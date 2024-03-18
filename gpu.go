package client

import (
	"github.com/wetee-dao/go-sdk/gen/types"
	"github.com/wetee-dao/go-sdk/gen/weteegpu"

	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/pkg/errors"
)

// Worker
type GpuApp struct {
	Client *ChainClient
	Signer *signature.KeyringPair
}

func (w *GpuApp) GetApp(publey []byte, id uint64) (*types.GpuApp, error) {
	if len(publey) != 32 {
		return nil, errors.New("publey length error")
	}

	var mss [32]byte
	copy(mss[:], publey)

	res, ok, err := weteegpu.GetGPUAppsLatest(w.Client.Api.RPC.State, mss, id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("GetK8sClusterAccountsLatest => not start")
	}
	return &res, nil
}

func (w *GpuApp) GetAccount(id uint64) ([]byte, error) {
	res, ok, err := weteegpu.GetAppIdAccountsLatest(w.Client.Api.RPC.State, id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("GetAppIdAccountsLatest => not ok")
	}
	return res[:], nil
}

func (w *GpuApp) GetVersionLatest(id uint64) (uint64, error) {
	res, ok, err := weteegpu.GetAppVersionLatest(w.Client.Api.RPC.State, id)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, errors.New("GetAppIdAccountsLatest => not ok")
	}
	return res, nil
}
