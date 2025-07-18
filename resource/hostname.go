package resource

import (
	"fmt"
	"os"
)

// 获取当前主机名
func (r *ResourceUsage) getHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", fmt.Errorf("获取主机名失败: %v", err)
	}
	return hostname, nil
}
