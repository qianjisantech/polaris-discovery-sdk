package conf

type PolarisDiscoveryConf struct {
	Addr              string // 监控中心地址
	HeartbeatInterval int    //心跳间隔 s
	Timeout           int    // 超时时间间隔 s
	Retry             struct {
		MaxAttempts int // 操作失败后的最大重试次数（包含首次尝试） 最大重试次数  默认3次
		Backoff     int //控制重试间隔的基数，实际间隔会随尝试次数指数增长（指数退避算法  重试间隔基数（单位通常为ms） 默认500ms
	}
}
