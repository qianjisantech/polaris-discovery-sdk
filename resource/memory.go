package resource

import "github.com/shirou/gopsutil/mem"

// MemoryUsage 内存使用情况
type MemoryUsage struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
}

// 获取当前内存使用情况
func (r *ResourceUsage) getMemoryUsage() (MemoryUsage, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return MemoryUsage{}, err
	}

	return MemoryUsage{
		Total:       v.Total,
		Used:        v.Used,
		Free:        v.Free,
		UsedPercent: v.UsedPercent,
	}, nil
}
