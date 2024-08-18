package module

import (
	"errors"

	chain "github.com/wetee-dao/go-sdk"
	gtypes "github.com/wetee-dao/go-sdk/pallet/types"
)

func GetAccount(client *chain.ChainClient, workID gtypes.WorkId) ([]byte, error) {
	if workID.Wtype.IsAPP {
		app := &App{
			Client: client,
			Signer: nil,
		}
		return app.GetAccount(workID.Id)
	} else if workID.Wtype.IsTASK {
		task := &Task{
			Client: client,
			Signer: nil,
		}
		return task.GetAccount(workID.Id)
	} else if workID.Wtype.IsGPU {
		gpuapp := &GpuApp{
			Client: client,
			Signer: nil,
		}
		return gpuapp.GetAccount(workID.Id)
	}
	return nil, errors.New("unknow work type")
}

func GetVersion(client *chain.ChainClient, workID gtypes.WorkId) (ret uint64, err error) {
	if workID.Wtype.IsAPP {
		app := &App{
			Client: client,
			Signer: nil,
		}
		return app.GetVersionLatest(workID.Id)
	} else if workID.Wtype.IsTASK {
		task := &Task{
			Client: client,
			Signer: nil,
		}
		return task.GetVersionLatest(workID.Id)
	} else if workID.Wtype.IsGPU {
		gpuapp := &GpuApp{
			Client: client,
			Signer: nil,
		}
		return gpuapp.GetVersionLatest(workID.Id)
	}

	return 0, errors.New("unknow work type")
}

func GetSecretEnv(client *chain.ChainClient, workID gtypes.WorkId) (ret []byte, err error) {
	if workID.Wtype.IsAPP {
		app := &App{
			Client: client,
			Signer: nil,
		}
		return app.GetSecretEnv(workID.Id)
	} else if workID.Wtype.IsTASK {
		task := &Task{
			Client: client,
			Signer: nil,
		}
		return task.GetSecretEnv(workID.Id)
	} else if workID.Wtype.IsGPU {
		gpuapp := &GpuApp{
			Client: client,
			Signer: nil,
		}
		return gpuapp.GetSecretEnv(workID.Id)
	}

	return nil, errors.New("unknow work type")
}
