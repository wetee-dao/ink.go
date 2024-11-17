package gpu

import types "github.com/wetee-dao/go-sdk/pallet/types"

// App create
// 注册任务
func MakeCreateCall(name0 []byte, image1 []byte, signer2 []byte, signature3 []byte, meta4 []byte, port5 []types.Service, command6 types.Command, env7 []types.EnvInput, secretEnv8 types.OptionTByteSlice, cpu9 uint32, memory10 uint32, disk11 []types.Disk, gpu12 uint32, sideContainer13 []types.Container, level14 byte, teeVersion15 types.TEEVersion) types.RuntimeCall {
	return types.RuntimeCall{
		IsGpu: true,
		AsGpuField0: &types.WeteeGpuPalletCall{
			IsCreate:                true,
			AsCreateName0:           name0,
			AsCreateImage1:          image1,
			AsCreateSigner2:         signer2,
			AsCreateSignature3:      signature3,
			AsCreateMeta4:           meta4,
			AsCreatePort5:           port5,
			AsCreateCommand6:        command6,
			AsCreateEnv7:            env7,
			AsCreateSecretEnv8:      secretEnv8,
			AsCreateCpu9:            cpu9,
			AsCreateMemory10:        memory10,
			AsCreateDisk11:          disk11,
			AsCreateGpu12:           gpu12,
			AsCreateSideContainer13: sideContainer13,
			AsCreateLevel14:         level14,
			AsCreateTeeVersion15:    teeVersion15,
		},
	}
}

// App update
// 更新任务
func MakeUpdateCall(appId0 uint64, newName1 types.OptionTByteSlice, newImage2 types.OptionTByteSlice, newSigner3 types.OptionTByteSlice, newSignature4 types.OptionTByteSlice, newPort5 types.OptionTServiceSlice, newCommand6 types.OptionTCommand, newEnv7 []types.EnvInput, secretEnv8 types.OptionTByteSlice, withRestart9 bool) types.RuntimeCall {
	return types.RuntimeCall{
		IsGpu: true,
		AsGpuField0: &types.WeteeGpuPalletCall{
			IsUpdate:              true,
			AsUpdateAppId0:        appId0,
			AsUpdateNewName1:      newName1,
			AsUpdateNewImage2:     newImage2,
			AsUpdateNewSigner3:    newSigner3,
			AsUpdateNewSignature4: newSignature4,
			AsUpdateNewPort5:      newPort5,
			AsUpdateNewCommand6:   newCommand6,
			AsUpdateNewEnv7:       newEnv7,
			AsUpdateSecretEnv8:    secretEnv8,
			AsUpdateWithRestart9:  withRestart9,
		},
	}
}

// App restart
// 更新任务
func MakeRestartCall(appId0 uint64) types.RuntimeCall {
	return types.RuntimeCall{
		IsGpu: true,
		AsGpuField0: &types.WeteeGpuPalletCall{
			IsRestart:       true,
			AsRestartAppId0: appId0,
		},
	}
}

// update price
// 更新价格
func MakeUpdatePriceCall(level0 byte, price1 types.Price2) types.RuntimeCall {
	return types.RuntimeCall{
		IsGpu: true,
		AsGpuField0: &types.WeteeGpuPalletCall{
			IsUpdatePrice:       true,
			AsUpdatePriceLevel0: level0,
			AsUpdatePricePrice1: price1,
		},
	}
}
