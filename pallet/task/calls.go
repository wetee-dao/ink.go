package task

import types "github.com/wetee-dao/go-sdk/pallet/types"

// Task create
// 注册任务
func MakeCreateCall(name0 []byte, image1 []byte, signer2 []byte, signature3 []byte, meta4 []byte, port5 []types.Service, command6 types.Command, env7 []types.EnvInput, secretEnv8 types.OptionTByteSlice, cpu9 uint32, memory10 uint32, disk11 []types.Disk, level12 byte, teeVersion13 types.TEEVersion) types.RuntimeCall {
	return types.RuntimeCall{
		IsTask: true,
		AsTaskField0: &types.WeteeTaskPalletCall{
			IsCreate:             true,
			AsCreateName0:        name0,
			AsCreateImage1:       image1,
			AsCreateSigner2:      signer2,
			AsCreateSignature3:   signature3,
			AsCreateMeta4:        meta4,
			AsCreatePort5:        port5,
			AsCreateCommand6:     command6,
			AsCreateEnv7:         env7,
			AsCreateSecretEnv8:   secretEnv8,
			AsCreateCpu9:         cpu9,
			AsCreateMemory10:     memory10,
			AsCreateDisk11:       disk11,
			AsCreateLevel12:      level12,
			AsCreateTeeVersion13: teeVersion13,
		},
	}
}

// Rerun task
// 重启任务
func MakeRerunCall(id0 uint64) types.RuntimeCall {
	return types.RuntimeCall{
		IsTask: true,
		AsTaskField0: &types.WeteeTaskPalletCall{
			IsRerun:    true,
			AsRerunId0: id0,
		},
	}
}

// Task update
// 更新任务
func MakeUpdateCall(appId0 uint64, newName1 types.OptionTByteSlice, newImage2 types.OptionTByteSlice, newSigner3 types.OptionTByteSlice, newSignature4 types.OptionTByteSlice, newPort5 types.OptionTServiceSlice, newCommand6 types.OptionTCommand, newEnv7 []types.EnvInput, secretEnv8 types.OptionTByteSlice, withRestart9 bool) types.RuntimeCall {
	return types.RuntimeCall{
		IsTask: true,
		AsTaskField0: &types.WeteeTaskPalletCall{
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

// update price
// 更新价格
func MakeUpdatePriceCall(level0 byte, price1 types.Price1) types.RuntimeCall {
	return types.RuntimeCall{
		IsTask: true,
		AsTaskField0: &types.WeteeTaskPalletCall{
			IsUpdatePrice:       true,
			AsUpdatePriceLevel0: level0,
			AsUpdatePricePrice1: price1,
		},
	}
}
