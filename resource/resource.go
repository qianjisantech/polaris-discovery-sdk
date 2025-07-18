package resource

import (
	"fmt"
)

// ResourceUsage 资源使用情况结构体
type ResourceUsage struct {
	Hostname  string       `json:"hostname"`
	CPU       CPUUsage     `json:"cpu"`
	Memory    MemoryUsage  `json:"memory"`
	Disks     []DiskUsage  `json:"disks"`
	Network   NetworkUsage `json:"network"`
	IPAddress string       `json:"ip_address"`
}

func (r *ResourceUsage) GetResourceUsage() (*ResourceUsage, error) {
	usage := &ResourceUsage{}
	hostname, err := r.getHostname()
	if err != nil {
		return nil, fmt.Errorf("获取hostname信息失败: %v", err)
	}
	usage.Hostname = hostname
	// 获取CPU信息
	cpuInfo, err := r.getCPUUsage()
	if err != nil {
		return nil, fmt.Errorf("获取CPU信息失败: %v", err)
	}
	usage.CPU = cpuInfo

	// 获取内存信息
	memInfo, err := r.getMemoryUsage()
	if err != nil {
		return nil, fmt.Errorf("获取内存信息失败: %v", err)
	}
	usage.Memory = memInfo

	// 获取磁盘信息
	disks, err := r.getDiskUsage()
	if err != nil {
		return nil, fmt.Errorf("获取磁盘信息失败: %v", err)
	}
	usage.Disks = disks

	// 获取网络信息
	netInfo, err := r.getNetworkUsage()
	if err != nil {
		return nil, fmt.Errorf("获取网络信息失败: %v", err)
	}
	usage.Network = netInfo
	addresses, err := r.getIPAddresses()
	if err != nil {
		return nil, fmt.Errorf("获取ip失败: %v", err)
	}
	usage.IPAddress = addresses
	return usage, nil
}
