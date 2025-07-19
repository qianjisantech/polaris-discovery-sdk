package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/qianjisantech/polaris-discovery-sdk/resource"
	"github.com/qianjisantech/polaris-discovery-sdk/util"
	"log"
	"time"
)

// HeatBeat 发送心跳
func (r *DiscoveryClient) heatBeat(code string) (*HeatBeatResponse, error) {
	log.Printf("-----开始发送心跳-------")
	var resourceUsage resource.ResourceUsage
	getResourceUsage, err := resourceUsage.GetResourceUsage()
	if err != nil {
		return nil, err
	}

	heatBeatRequest := HeatBeatRequest{
		IP:                    getResourceUsage.IPAddress,
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
		Status:                "busy",
		HostName:              getResourceUsage.Hostname,
	}
	if r.Timeout <= 0 {
		r.Timeout = 30 // 默认超时时间30s
	}
	if r.Timeout > 60 {
		r.Timeout = 60 // 默认最大超时时间60s
	}
	client := util.NewHttpClient(time.Duration(r.Timeout))
	heatbeatUrl := r.Addr + string(HeatBeatUrl)
	res, err := client.PostJSON(context.Background(), heatbeatUrl+code, heatBeatRequest)
	if err != nil {
		return nil, err
	}

	var heatBeatResponse HeatBeatResponse
	if err := json.Unmarshal(res, &heatBeatResponse); err != nil {
		return nil, fmt.Errorf("无法解析心跳响应----------> %s", string(res))
	}

	return &heatBeatResponse, nil
}
