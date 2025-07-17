package model

// RegisterRequest 注册请求体
type RegisterRequest struct {
	IP                    string `json:"ip"`
	CPUCores              int    `json:"cpu_cores"`
	CPUUsedPercent        string `json:"cpu_used_percent"`
	CPUFreePercent        string `json:"cpu_free_percent"`
	MemoryTotal           string `json:"memory_total"`
	MemoryUsed            string `json:"memory_used"`
	MemoryFree            string `json:"memory_free"`
	MemoryUsedPercent     string `json:"memory_used_percent"`
	NetworkBytesSent      string `json:"network_bytes_sent"`
	NetworkBytesReceive   string `json:"network_bytes_receive"`
	NetworkPacketsSent    string `json:"network_packets_sent"`
	NetworkPacketsReceive string `json:"network_packets_receive"`
}

// RegisterResponse 注册返回体
type RegisterResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		IP                    string `json:"ip"`
		CPUCores              int    `json:"cpu_cores"`
		CPUUsedPercent        string `json:"cpu_used_percent"`
		CPUFreePercent        string `json:"cpu_free_percent"`
		MemoryTotal           string `json:"memory_total"`
		MemoryUsed            string `json:"memory_used"`
		MemoryFree            string `json:"memory_free"`
		MemoryUsedPercent     string `json:"memory_used_percent"`
		NetworkBytesSent      string `json:"network_bytes_sent"`
		NetworkBytesReceive   string `json:"network_bytes_receive"`
		NetworkPacketsSent    string `json:"network_packets_sent"`
		NetworkPacketsReceive string `json:"network_packets_receive"`
	} `json:"data,omitempty"`
}

// HeatBeatRequest 心跳请求体
type HeatBeatRequest struct {
	ID                    string `json:"id"`
	IP                    string `json:"ip"`
	Status                string `json:"status"`
	CPUCores              int    `json:"cpu_cores"`
	CPUUsedPercent        string `json:"cpu_used_percent"`
	CPUFreePercent        string `json:"cpu_free_percent"`
	MemoryTotal           string `json:"memory_total"`
	MemoryUsed            string `json:"memory_used"`
	MemoryFree            string `json:"memory_free"`
	MemoryUsedPercent     string `json:"memory_used_percent"`
	NetworkBytesSent      string `json:"network_bytes_sent"`
	NetworkBytesReceive   string `json:"network_bytes_receive"`
	NetworkPacketsSent    string `json:"network_packets_sent"`
	NetworkPacketsReceive string `json:"network_packets_receive"`
}

// HeatBeatResponse 心跳请求体
type HeatBeatResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		IP   string `json:"ip"`
		Name string `json:"name"`
	} `json:"data,omitempty"`
}
