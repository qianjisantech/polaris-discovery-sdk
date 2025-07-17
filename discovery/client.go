package discovery

import (
	"context"
	"encoding/json"
	"fmt"
	"polaris-discovery/model"
	"polaris-discovery/request"
)

// DiscoveryClient 注册中心Client结构体
type DiscoveryClient struct {
	basePath string
}

// NewDiscoveryClient 注册中心Client
func NewDiscoveryClient(addr string) *DiscoveryClient {
	return &DiscoveryClient{
		basePath: addr + "/polaris/discovery",
	}
}

// Register 注册机器
func (r *DiscoveryClient) Register(body model.RegisterRequest) (*model.RegisterResponse, error) {
	client := request.NewClient(r.basePath, 20)
	res, err := client.PostJSON(context.Background(), "/register", body)
	if err != nil {
		return nil, err
	}
	// 尝试解析为成功响应
	var registerResponse model.RegisterResponse
	if err := json.Unmarshal(res, &registerResponse); err == nil {
		return &registerResponse, nil
	}

	// 无法解析的响应
	return nil, fmt.Errorf("无法解析响应: %s", string(res))
}

// HeatBeat 发送心跳
func (r *DiscoveryClient) HeatBeat(code string, body model.HeatBeatRequest) (*model.HeatBeatResponse, error) {
	client := request.NewClient(r.basePath, 20)
	res, err := client.PostJSON(context.Background(), "/heatbeat/"+code, body)
	if err != nil {
		return nil, err // 修正错误返回
	}
	// 尝试解析为成功响应
	var heatBeatResponse model.HeatBeatResponse
	if err := json.Unmarshal(res, &heatBeatResponse); err == nil {
		return &heatBeatResponse, nil
	}

	// 无法解析的响应
	return nil, fmt.Errorf("无法解析响应: %s", string(res))
}
