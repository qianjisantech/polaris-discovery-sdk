package resource

import "github.com/shirou/gopsutil/disk"

// DiskUsage 磁盘使用情况
type DiskUsage struct {
	Device      string  `json:"device"`
	MountPoint  string  `json:"mount_point"`
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
	FileSystem  string  `json:"file_system"`
}

// 获取当前磁盘信息
func (r *ResourceUsage) getDiskUsage() ([]DiskUsage, error) {
	partitions, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}

	var disks []DiskUsage
	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue
		}

		disks = append(disks, DiskUsage{
			Device:      p.Device,
			MountPoint:  p.Mountpoint,
			Total:       usage.Total,
			Used:        usage.Used,
			Free:        usage.Free,
			UsedPercent: usage.UsedPercent,
			FileSystem:  p.Fstype,
		})
	}

	return disks, nil
}
