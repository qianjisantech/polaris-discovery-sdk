package core

// RegisterRequest 注册请求体
type RegisterRequest struct {
	IP                    string  `json:"ip"`
	Hostname              string  `json:"hostname"`
	CPUCores              int     `json:"cpu_cores"`
	CPUUsedPercent        float64 `json:"cpu_used_percent"`
	CPUFreePercent        float64 `json:"cpu_free_percent"`
	MemoryTotal           uint64  `json:"memory_total"`
	MemoryUsed            uint64  `json:"memory_used"`
	MemoryFree            uint64  `json:"memory_free"`
	MemoryUsedPercent     float64 `json:"memory_used_percent"`
	NetworkBytesSent      uint64  `json:"network_bytes_sent"`
	NetworkBytesReceive   uint64  `json:"network_bytes_receive"`
	NetworkPacketsSent    uint64  `json:"network_packets_sent"`
	NetworkPacketsReceive uint64  `json:"network_packets_receive"`
	Status                string  `json:"status"`
	ExecuteStatus         string  `json:"execute_status"`
}

// RegisterResponse 注册返回体
type RegisterResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Id                 string `json:"id"`
		IdentificationCode string `json:"identification_code"`
	} `json:"data,omitempty"`
}

// HeatBeatRequest 心跳请求体
type HeatBeatRequest struct {
	IP                    string  `json:"ip"`
	Status                string  `json:"status"`
	ExecuteStatus         string  `json:"execute_status"`
	CPUCores              int     `json:"cpu_cores"`
	CPUUsedPercent        float64 `json:"cpu_used_percent"`
	CPUFreePercent        float64 `json:"cpu_free_percent"`
	MemoryTotal           uint64  `json:"memory_total"`
	MemoryUsed            uint64  `json:"memory_used"`
	MemoryFree            uint64  `json:"memory_free"`
	MemoryUsedPercent     float64 `json:"memory_used_percent"`
	NetworkBytesSent      uint64  `json:"network_bytes_sent"`
	NetworkBytesReceive   uint64  `json:"network_bytes_receive"`
	NetworkPacketsSent    uint64  `json:"network_packets_sent"`
	NetworkPacketsReceive uint64  `json:"network_packets_receive"`
	HostName              string  `json:"hostname"`
}

// HeatBeatResponse 心跳请求体
type HeatBeatResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Id    string                 `json:"id"`
		Name  string                 `json:"name"`
		Tasks []HeatBeatResponseTask `json:"tasks,omitempty"`
	} `json:"data,omitempty"`
}

type HeatBeatResponseTask struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Status       string `json:"status"`
	ListenPort   int    `json:"listen_port"`
	CreateTime   string `json:"create_time"`
	CreateBy     string `json:"create_by"`
	CreateByName string `json:"create_by_name"`
	UpdateBy     string `json:"update_by"`
	UpdateByName string `json:"update_by_name"`
	UpdateTime   string `json:"update_time"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	ExecuteTime  string `json:"execute_time"`
}
