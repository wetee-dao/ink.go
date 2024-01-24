package client

import (
	"github.com/wetee-dao/go-sdk/gen/weteetask"

	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/pkg/errors"
)

// Worker
type Task struct {
	Client *ChainClient
	Signer *signature.KeyringPair
}

func (w *Task) GetAccount(id uint64) ([]byte, error) {
	res, ok, err := weteetask.GetTaskIdAccountsLatest(w.Client.Api.RPC.State, id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("GetAppIdAccountsLatest => not ok")
	}
	return res[:], nil
}
