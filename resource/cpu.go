package resource

import (
	"github.com/shirou/gopsutil/cpu"
	"time"
)

// CPUUsage CPU使用情况
type CPUUsage struct {
	Cores       int32   `json:"cores"`
	UsedPercent float64 `json:"used_percent"`
	FreePercent float64 `json:"free_percent"`
}

// 获取当前cpu使用情况
func (r *ResourceUsage) getCPUUsage() (CPUUsage, error) {
	// 获取CPU核心数
	cores, err := cpu.Counts(true)
	if err != nil {
		return CPUUsage{}, err
	}

	// 获取CPU使用率
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return CPUUsage{}, err
	}

	var usedPercent float64
	if len(percent) > 0 {
		usedPercent = percent[0]
	}

	return CPUUsage{
		Cores:       int32(cores),
		UsedPercent: usedPercent,
		FreePercent: 100 - usedPercent,
	}, nil
}
