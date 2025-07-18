package resource

import "github.com/shirou/gopsutil/net"

// NetworkUsage 网络使用情况
type NetworkUsage struct {
	BytesSent      uint64 `json:"bytes_sent"`
	BytesReceive   uint64 `json:"bytes_receive"`
	PacketsSent    int32  `json:"packets_sent"`
	PacketsReceive int32  `json:"packets_receive"`
}

// 获取当前网络使用情况
func (r *ResourceUsage) getNetworkUsage() (NetworkUsage, error) {
	counters, err := net.IOCounters(true)
	if err != nil {
		return NetworkUsage{}, err
	}

	var totalSent, totalRecv, totalPacketsSent, totalPacketsRecv uint64
	for _, counter := range counters {
		totalSent += counter.BytesSent
		totalRecv += counter.BytesRecv
		totalPacketsSent += counter.PacketsSent
		totalPacketsRecv += counter.PacketsRecv
	}

	return NetworkUsage{
		BytesSent:      totalSent,
		BytesReceive:   totalRecv,
		PacketsSent:    int32(totalPacketsSent),
		PacketsReceive: int32(totalPacketsRecv),
	}, nil
}
