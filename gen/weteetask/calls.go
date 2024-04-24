package weteetask

import (
	types1 "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types "github.com/wetee-dao/go-sdk/gen/types"
)

// Task create
// 注册任务
func MakeCreateCall(name0 []byte, image1 []byte, meta2 []byte, port3 []uint32, cpu4 uint32, memory5 uint32, disk6 []types.Disk, level7 byte, deposit8 types1.UCompact) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeteeTask: true,
		AsWeteeTaskField0: &types.WeteeTaskPalletCall{
			IsCreate:         true,
			AsCreateName0:    name0,
			AsCreateImage1:   image1,
			AsCreateMeta2:    meta2,
			AsCreatePort3:    port3,
			AsCreateCpu4:     cpu4,
			AsCreateMemory5:  memory5,
			AsCreateDisk6:    disk6,
			AsCreateLevel7:   level7,
			AsCreateDeposit8: deposit8,
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
func MakeUpdateCall(appId0 uint64, name1 []byte, image2 []byte, port3 []uint32, withRestart4 bool) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeteeTask: true,
		AsWeteeTaskField0: &types.WeteeTaskPalletCall{
			IsUpdate:             true,
			AsUpdateAppId0:       appId0,
			AsUpdateName1:        name1,
			AsUpdateImage2:       image2,
			AsUpdatePort3:        port3,
			AsUpdateWithRestart4: withRestart4,
		},
	}
}

// Task settings
// 任务设置
func MakeSetSettingsCall(appId0 uint64, value1 []types.AppSettingInput, withRestart2 bool) types.RuntimeCall {
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
