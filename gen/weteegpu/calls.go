package weteegpu

import (
	types1 "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types "github.com/wetee-dao/go-sdk/gen/types"
)

// App create
// 注册任务
func MakeCreateCall(name0 []byte, image1 []byte, meta2 []byte, port3 []uint32, cpu4 uint32, memory5 uint32, disk6 []types.Disk, gpu7 uint32, level8 byte, deposit9 types1.UCompact) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeteeGpu: true,
		AsWeteeGpuField0: &types.WeteeGpuPalletCall{
			IsCreate:         true,
			AsCreateName0:    name0,
			AsCreateImage1:   image1,
			AsCreateMeta2:    meta2,
			AsCreatePort3:    port3,
			AsCreateCpu4:     cpu4,
			AsCreateMemory5:  memory5,
			AsCreateDisk6:    disk6,
			AsCreateGpu7:     gpu7,
			AsCreateLevel8:   level8,
			AsCreateDeposit9: deposit9,
		},
	}
}

// App update
// 更新任务
func MakeUpdateCall(appId0 uint64, name1 []byte, image2 []byte, port3 []uint32, withRestart4 bool) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeteeGpu: true,
		AsWeteeGpuField0: &types.WeteeGpuPalletCall{
			IsUpdate:             true,
			AsUpdateAppId0:       appId0,
			AsUpdateName1:        name1,
			AsUpdateImage2:       image2,
			AsUpdatePort3:        port3,
			AsUpdateWithRestart4: withRestart4,
		},
	}
}

// App settings
// 任务设置
func MakeSetSettingsCall(appId0 uint64, value1 []types.AppSettingInput, withRestart2 bool) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeteeGpu: true,
		AsWeteeGpuField0: &types.WeteeGpuPalletCall{
			IsSetSettings:             true,
			AsSetSettingsAppId0:       appId0,
			AsSetSettingsValue1:       value1,
			AsSetSettingsWithRestart2: withRestart2,
		},
	}
}

// App charge
// 任务充值
func MakeRechargeCall(id0 uint64, deposit1 types1.U128) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeteeGpu: true,
		AsWeteeGpuField0: &types.WeteeGpuPalletCall{
			IsRecharge:         true,
			AsRechargeId0:      id0,
			AsRechargeDeposit1: deposit1,
		},
	}
}

// App restart
// 更新任务
func MakeRestartCall(appId0 uint64, withRestart1 bool) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeteeGpu: true,
		AsWeteeGpuField0: &types.WeteeGpuPalletCall{
			IsRestart:             true,
			AsRestartAppId0:       appId0,
			AsRestartWithRestart1: withRestart1,
		},
	}
}
