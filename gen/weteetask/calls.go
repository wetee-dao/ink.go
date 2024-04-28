package weteetask

import (
	types1 "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types "github.com/wetee-dao/go-sdk/gen/types"
)

// Task create
// 注册任务
func MakeCreateCall(name0 []byte, image1 []byte, meta2 []byte, port3 []types.Service, command4 types.Command, env5 []types.EnvInput, cpu6 uint32, memory7 uint32, disk8 []types.Disk, level9 byte, deposit10 types1.UCompact) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeteeTask: true,
		AsWeteeTaskField0: &types.WeteeTaskPalletCall{
			IsCreate:          true,
			AsCreateName0:     name0,
			AsCreateImage1:    image1,
			AsCreateMeta2:     meta2,
			AsCreatePort3:     port3,
			AsCreateCommand4:  command4,
			AsCreateEnv5:      env5,
			AsCreateCpu6:      cpu6,
			AsCreateMemory7:   memory7,
			AsCreateDisk8:     disk8,
			AsCreateLevel9:    level9,
			AsCreateDeposit10: deposit10,
		},
	}
}

// Rerun task
// 重启任务
func MakeRerunCall(id0 uint64) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeteeTask: true,
		AsWeteeTaskField0: &types.WeteeTaskPalletCall{
			IsRerun:    true,
			AsRerunId0: id0,
		},
	}
}

// Task update
// 更新任务
func MakeUpdateCall(appId0 uint64, newName1 types.OptionTByteSlice, newImage2 types.OptionTByteSlice, newPort3 types.OptionTServiceSlice, newCommand4 types.OptionTCommand, newEnv5 []types.EnvInput, withRestart6 bool) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeteeTask: true,
		AsWeteeTaskField0: &types.WeteeTaskPalletCall{
			IsUpdate:             true,
			AsUpdateAppId0:       appId0,
			AsUpdateNewName1:     newName1,
			AsUpdateNewImage2:    newImage2,
			AsUpdateNewPort3:     newPort3,
			AsUpdateNewCommand4:  newCommand4,
			AsUpdateNewEnv5:      newEnv5,
			AsUpdateWithRestart6: withRestart6,
		},
	}
}

// Task settings
// 任务设置
func MakeSetSettingsCall(appId0 uint64, value1 []types.EnvInput, withRestart2 bool) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeteeTask: true,
		AsWeteeTaskField0: &types.WeteeTaskPalletCall{
			IsSetSettings:             true,
			AsSetSettingsAppId0:       appId0,
			AsSetSettingsValue1:       value1,
			AsSetSettingsWithRestart2: withRestart2,
		},
	}
}

// Task charge
// 任务充值
func MakeChargeCall(id0 uint64, deposit1 types1.U128) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeteeTask: true,
		AsWeteeTaskField0: &types.WeteeTaskPalletCall{
			IsCharge:         true,
			AsChargeId0:      id0,
			AsChargeDeposit1: deposit1,
		},
	}
}
