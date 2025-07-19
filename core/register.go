package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/qianjisantech/polaris-discovery-sdk/resource"
	"github.com/qianjisantech/polaris-discovery-sdk/util"
	"time"
)

// Register 注册机器
func (r *DiscoveryClient) register() (*RegisterResponse, error) {
	var resourceUsage resource.ResourceUsage
	getResourceUsage, err := resourceUsage.GetResourceUsage()
	if err != nil {
		return nil, err
	}

	registerRequest := RegisterRequest{
		IP:                    getResourceUsage.IPAddress,
		Hostname:              getResourceUsage.Hostname,
		CPUCores:              int(getResourceUsage.CPU.Cores),
		CPUUsedPercent:        getResourceUsage.CPU.UsedPercent,
		CPUFreePercent:        getResourceUsage.CPU.FreePercent,
		MemoryTotal:           getResourceUsage.Memory.Total,
		MemoryUsed:            getResourceUsage.Memory.Used,
		MemoryFree:            getResourceUsage.Memory.Free,
		MemoryUsedPercent:     getResourceUsage.Memory.UsedPercent,
		NetworkBytesSent:      getResourceUsage.Network.BytesSent,
		NetworkBytesReceive:   getResourceUsage.Network.BytesReceive,
		NetworkPacketsSent:    uint64(getResourceUsage.Network.PacketsSent),
		NetworkPacketsReceive: uint64(getResourceUsage.Network.PacketsReceive),
	}
	if r.Timeout <= 0 {
		r.Timeout = 30 // 默认超时时间30s
	}
	if r.Timeout > 60 {
		r.Timeout = 60 // 默认最大超时时间60s
	}
	client := util.NewHttpClient(time.Duration(r.Timeout))
	res, err := client.PostJSON(context.Background(), string(RegisterUrl), registerRequest)
	if err != nil {
		return nil, fmt.Errorf("无法解析注册响应----------> %s", string(res))
	}

	var registerResponse RegisterResponse
	if err := json.Unmarshal(res, &registerResponse); err != nil {
		return nil, fmt.Errorf("无法解析响应: %s", string(res))
	}

	return &registerResponse, nil
}
