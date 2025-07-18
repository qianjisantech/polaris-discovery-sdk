package discovery

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"polaris-discovery/model"
	"polaris-discovery/request"
	"polaris-discovery/resource"
	"time"
)

// DiscoveryClient 注册中心Client结构体
type DiscoveryClient struct {
	Addr     string
	basePath string
	ticker   *time.Ticker // 保存定时器引用
	done     chan struct{}
}

// NewDiscoveryClient 注册中心Client
func (r *DiscoveryClient) NewDiscoveryClient() *DiscoveryClient {
	return &DiscoveryClient{
		basePath: r.Addr + "/polaris/discovery",
		done:     make(chan struct{}),
	}
}

func (r *DiscoveryClient) Start(
	onRegisterSuccess func(*model.RegisterResponse),
	onRegisterError func(error),
	onHeartbeatSuccess func(*model.HeatBeatResponse),
	onHeartbeatError func(error),
) error {
	client := r.NewDiscoveryClient()
	registerResp, err := client.register()
	if err != nil {
		if onRegisterError != nil {
			onRegisterError(err)
		}
		return err
	}

	if registerResp.Success {
		if onRegisterSuccess != nil {
			onRegisterSuccess(registerResp)
		}
		log.Printf("注册成功: %v", registerResp)
		if registerResp.Data.IdentificationCode != "" {
			// 保存定时器，避免被 GC 回收
			r.ticker = NewTimer(context.Background(), 2*time.Second, func() {
				log.Println("执行心跳任务...")
				resp, err := client.heatBeat(registerResp.Data.IdentificationCode)
				if err != nil {
					log.Printf("心跳失败: %v", err)
					if onHeartbeatError != nil {
						onHeartbeatError(err)
					}
					return
				}
				if resp.Success {
					log.Println("心跳成功")
					if onHeartbeatSuccess != nil {
						onHeartbeatSuccess(resp)
					}
				} else {
					log.Printf("心跳异常: %s", resp.Message)
					if onHeartbeatError != nil {
						onHeartbeatError(fmt.Errorf(resp.Message))
					}
				}
			})
		}
	} else {
		err := fmt.Errorf("服务注册异常: %s", registerResp.Message)
		if onRegisterError != nil {
			onRegisterError(err)
		}
		return err
	}
	return nil
}

// Register 注册机器
func (r *DiscoveryClient) register() (*model.RegisterResponse, error) {
	var resourceUsage resource.ResourceUsage
	getResourceUsage, err := resourceUsage.GetResourceUsage()
	if err != nil {
		return nil, err
	}

	registerRequest := model.RegisterRequest{
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

	client := request.NewClient(r.basePath, 20)
	res, err := client.PostJSON(context.Background(), "/register", registerRequest)
	if err != nil {
		return nil, err
	}

	var registerResponse model.RegisterResponse
	if err := json.Unmarshal(res, &registerResponse); err != nil {
		return nil, fmt.Errorf("无法解析响应: %s", string(res))
	}

	return &registerResponse, nil
}

// HeatBeat 发送心跳
func (r *DiscoveryClient) heatBeat(code string) (*model.HeatBeatResponse, error) {
	log.Printf("发送心跳")
	var resourceUsage resource.ResourceUsage
	getResourceUsage, err := resourceUsage.GetResourceUsage()
	if err != nil {
		return nil, err
	}

	heatBeatRequest := model.HeatBeatRequest{
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

	client := request.NewClient(r.basePath, 20)
	res, err := client.PostJSON(context.Background(), "/heatbeat/"+code, heatBeatRequest)
	if err != nil {
		return nil, err
	}

	var heatBeatResponse model.HeatBeatResponse
	if err := json.Unmarshal(res, &heatBeatResponse); err != nil {
		return nil, fmt.Errorf("无法解析响应: %s", string(res))
	}

	return &heatBeatResponse, nil
}
