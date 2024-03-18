package client

import (
	"errors"

	gtypes "github.com/wetee-dao/go-sdk/gen/types"
)

func GetAccount(client *ChainClient, workID gtypes.WorkId) ([]byte, error) {
	if workID.Wtype.IsAPP {
		app := &App{
			Client: client,
			Signer: nil,
		}
		return app.GetAccount(workID.Id)
	}
	if workID.Wtype.IsTASK {
		task := &Task{
			Client: client,
			Signer: nil,
		}
		return task.GetAccount(workID.Id)
	}
	if workID.Wtype.IsGPU {
		gpuapp := &GpuApp{
			Client: client,
			Signer: nil,
		}
		return gpuapp.GetAccount(workID.Id)
	}
	return nil, errors.New("unknow work type")
}

func GetVersion(client *ChainClient, workID gtypes.WorkId) (ret uint64, err error) {
	if workID.Wtype.IsAPP {
		app := &App{
			Client: client,
			Signer: nil,
		}
		return app.GetVersionLatest(workID.Id)
	}
	if workID.Wtype.IsTASK {
		task := &Task{
			Client: client,
			Signer: nil,
		}
		return task.GetVersionLatest(workID.Id)
	}
	if workID.Wtype.IsGPU {
		gpuapp := &GpuApp{
			Client: client,
			Signer: nil,
		}
		return gpuapp.GetVersionLatest(workID.Id)
	}
	return 0, nil
}
