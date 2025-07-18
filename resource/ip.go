package resource

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"runtime"
	"strings"
)

// 获取当前ip 地址
func (r *ResourceUsage) getIPAddresses() (string, error) {
	var ipList []string

	interfaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("获取网络接口失败: %v", err)
	}

	for _, iface := range interfaces {
		// 操作系统过滤逻辑
		switch runtime.GOOS {
		case "windows":
			if isWindowsLoopback(iface) || !isWindowsInterfaceUp(iface) {
				continue
			}
		case "linux", "darwin":
			if contains(iface.Flags, "loopback") || !contains(iface.Flags, "up") {
				continue
			}
		}

		// 提取纯IP
		for _, addr := range iface.Addrs {
			ip := strings.Split(addr.Addr, "/")[0]
			if isIPv4(ip) { // 额外过滤IPv4
				ipList = append(ipList, ip)
			}
		}
	}

	switch len(ipList) {
	case 0:
		return "", fmt.Errorf("未找到有效的IPv4地址")
	case 1:
		return ipList[0], nil
	default:
		return strings.Join(ipList, ", "), nil
	}
}

// 检查是否为IPv4
func isIPv4(ip string) bool {
	return strings.Count(ip, ".") == 3
}

// Windows特定判断函数
func isWindowsLoopback(iface net.InterfaceStat) bool {
	// Windows回环接口通常包含"Loopback"描述
	return strings.Contains(strings.ToLower(iface.Name), "loopback") ||
		strings.Contains(strings.ToLower(iface.HardwareAddr), "loopback")
}

func isWindowsInterfaceUp(iface net.InterfaceStat) bool {
	// Windows接口启用状态判断
	return !strings.Contains(strings.ToLower(iface.Name), "disabled") &&
		iface.MTU > 0
}

// Unix-like系统辅助函数
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
