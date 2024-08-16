package weteegpu

import types "github.com/wetee-dao/go-sdk/pallet/types"

// App create
// 注册任务
func MakeCreateCall(name0 []byte, image1 []byte, meta2 []byte, port3 []types.Service, command4 types.Command, env5 []types.EnvInput, secretEnv6 types.OptionTByteSlice, cpu7 uint32, memory8 uint32, disk9 []types.Disk, gpu10 uint32, sideContainer11 []types.Container, level12 byte, teeVersion13 types.TEEVersion) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEGpu: true,
		AsWeTEEGpuField0: &types.WeteeGpuPalletCall{
			IsCreate:                true,
			AsCreateName0:           name0,
			AsCreateImage1:          image1,
			AsCreateMeta2:           meta2,
			AsCreatePort3:           port3,
			AsCreateCommand4:        command4,
			AsCreateEnv5:            env5,
			AsCreateSecretEnv6:      secretEnv6,
			AsCreateCpu7:            cpu7,
			AsCreateMemory8:         memory8,
			AsCreateDisk9:           disk9,
			AsCreateGpu10:           gpu10,
			AsCreateSideContainer11: sideContainer11,
			AsCreateLevel12:         level12,
			AsCreateTeeVersion13:    teeVersion13,
		},
	}
}

// App update
// 更新任务
func MakeUpdateCall(appId0 uint64, newName1 types.OptionTByteSlice, newImage2 types.OptionTByteSlice, newPort3 types.OptionTServiceSlice, newCommand4 types.OptionTCommand, newEnv5 []types.EnvInput, secretEnv6 types.OptionTByteSlice, withRestart7 bool) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEGpu: true,
		AsWeTEEGpuField0: &types.WeteeGpuPalletCall{
			IsUpdate:             true,
			AsUpdateAppId0:       appId0,
			AsUpdateNewName1:     newName1,
			AsUpdateNewImage2:    newImage2,
			AsUpdateNewPort3:     newPort3,
			AsUpdateNewCommand4:  newCommand4,
			AsUpdateNewEnv5:      newEnv5,
			AsUpdateSecretEnv6:   secretEnv6,
			AsUpdateWithRestart7: withRestart7,
		},
	}
}

// App restart
// 更新任务
func MakeRestartCall(appId0 uint64) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEGpu: true,
		AsWeTEEGpuField0: &types.WeteeGpuPalletCall{
			IsRestart:       true,
			AsRestartAppId0: appId0,
		},
	}
}

// update price
// 更新价格
func MakeUpdatePriceCall(level0 byte, price1 types.Price2) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEGpu: true,
		AsWeTEEGpuField0: &types.WeteeGpuPalletCall{
			IsUpdatePrice:       true,
			AsUpdatePriceLevel0: level0,
			AsUpdatePricePrice1: price1,
		},
	}
}
