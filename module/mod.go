package module

import (
	"errors"
	"fmt"

	chain "github.com/wetee-dao/go-sdk"
	gtypes "github.com/wetee-dao/go-sdk/pallet/types"
	"github.com/wetee-dao/go-sdk/pallet/weteeapp"
	"github.com/wetee-dao/go-sdk/pallet/weteedsecret"
	"github.com/wetee-dao/go-sdk/pallet/weteegpu"
	"github.com/wetee-dao/go-sdk/pallet/weteetask"
	"github.com/wetee-dao/go-sdk/pallet/weteeworker"
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

func GetSecretEnv(client *chain.ChainClient, workID gtypes.WorkId) (ret []byte, isSome bool, err error) {
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

	return nil, false, errors.New("unknow work type")
}

// GetWorkCodeSignature 函数根据工作 ID 获取相应的代码签名
func GetWorkCodeSignature(client *chain.ChainClient, workID gtypes.WorkId) (ret []byte, err error) {
	// 判断工作类型是否为 APP
	if workID.Wtype.IsAPP {
		// 调用 weteeapp 获取代码签名的最新数据
		return weteeapp.GetCodeSignatureLatest(client.Api.RPC.State, workID.Id)
	} else if workID.Wtype.IsTASK {
		// 调用 weteetask 获取代码签名的最新数据
		return weteetask.GetCodeSignatureLatest(client.Api.RPC.State, workID.Id)
	} else if workID.Wtype.IsGPU {
		// 调用 weteegpu 获取代码签名的最新数据
		return weteegpu.GetCodeSignatureLatest(client.Api.RPC.State, workID.Id)
	}

	// 如果工作类型未知，返回错误信息
	return nil, errors.New("unknow work type")
}

// 获取全网当前程序的代码版本
// Get CodeSignature/SIgner
func GetDsecretCode(client *chain.ChainClient) ([]byte, []byte, error) {
	// 检查节点代码是否和 wetee 上要求的版本一致
	codeSignature, err := weteedsecret.GetCodeSignatureLatest(client.Api.RPC.State)
	if err != nil {
		fmt.Println("Get code signature error:", err)
		return nil, nil, err
	}
	codeSigner, err := weteedsecret.GetCodeSignerLatest(client.Api.RPC.State)
	if err != nil {
		fmt.Println("Get code signer error:", err)
		return nil, nil, err
	}

	return codeSignature, codeSigner, nil
}

// GetWorkerCode 函数用于获取 weteeworker 的代码签名和签名者
func GetWorkerCode(client *chain.ChainClient) ([]byte, []byte, error) {
	// 获取 weteeworker 的最新代码签名
	codeSignature, err := weteeworker.GetCodeSignatureLatest(client.Api.RPC.State)
	// 处理获取代码签名过程中的错误
	if err != nil {
		fmt.Println("Get code signature error:", err)
		return nil, nil, err
	}

	// 获取 weteeworker 的最新代码签名者
	codeSigner, err := weteeworker.GetCodeSignerLatest(client.Api.RPC.State)
	// 处理获取代码签名者过程中的错误
	if err != nil {
		fmt.Println("Get code signer error:", err)
		return nil, nil, err
	}

	// 返回获取到的代码签名和签名者
	return codeSignature, codeSigner, nil
}

// GetWorkCode 函数用于获取工作代码签名和代码签名者
func GetWorkCode(client *chain.ChainClient, workID gtypes.WorkId) ([]byte, []byte, error) {
	// 获取工作签名
	sig, err := GetWorkSignature(client, workID)
	if err != nil {
		return nil, nil, err
	}

	// 获取工作代码签名者
	signer, err := GetWorkCodeSigner(client, workID)
	if err != nil {
		return nil, nil, err
	}

	// 返回获取到的签名和签名者，以及一个 nil 错误
	return sig, signer, nil
}

// GetWorkSignature 函数根据工作 ID 获取相应的代码签名
func GetWorkSignature(client *chain.ChainClient, workID gtypes.WorkId) ([]byte, error) {
	// 判断工作类型是否为 APP
	if workID.Wtype.IsAPP {
		// 调用 weteeapp 获取代码签名的最新数据
		return weteeapp.GetCodeSignatureLatest(client.Api.RPC.State, workID.Id)
	} else if workID.Wtype.IsTASK {
		// 调用 weteetask 获取代码签名的最新数据
		return weteetask.GetCodeSignatureLatest(client.Api.RPC.State, workID.Id)
	} else if workID.Wtype.IsGPU {
		// 调用 weteegpu 获取代码签名的最新数据
		return weteegpu.GetCodeSignatureLatest(client.Api.RPC.State, workID.Id)
	}

	// 如果工作类型未知，返回错误信息
	return nil, errors.New("unknow work type")
}

// GetWorkCodeSigner 函数根据工作 ID 获取相应的代码签名者
func GetWorkCodeSigner(client *chain.ChainClient, workID gtypes.WorkId) ([]byte, error) {
	// 判断工作类型是否为 APP
	if workID.Wtype.IsAPP {
		// 调用 weteeapp 获取代码签名者的最新数据
		return weteeapp.GetCodeSignerLatest(client.Api.RPC.State, workID.Id)
	} else if workID.Wtype.IsTASK {
		// 调用 weteetask 获取代码签名者的最新数据
		return weteetask.GetCodeSignerLatest(client.Api.RPC.State, workID.Id)
	} else if workID.Wtype.IsGPU {
		// 调用 weteegpu 获取代码签名者的最新数据
		return weteegpu.GetCodeSignerLatest(client.Api.RPC.State, workID.Id)
	}

	// 如果工作类型未知，返回错误信息
	return nil, errors.New("unknow work type")
}
