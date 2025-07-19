package main

import (
	"github.com/qianjisantech/polaris-discovery-sdk/core"
	"log"
	"time"
)

func main() {

	client := &core.DiscoveryClient{
		Addr:              "http://localhost:8080",
		HeartbeatInterval: 5,
		Timeout:           30,
		Retry: core.Retry{
			MaxAttempts: 3,
			Backoff:     2000, // 2秒
		},
	}

	err := client.Start(
		func(resp *core.RegisterResponse) {
			log.Printf("注册成功! ID: %s", resp.Data.Id)
		},
		func(err error) {
			log.Printf("注册失败: %v", err)
		},
		func(resp *core.HeatBeatResponse) {
			log.Printf("心跳成功 id: %s", resp.Data.Id)
		},
		func(err error) {
			log.Printf("心跳失败: %v", err)
		},
	)

	if err != nil {
		log.Printf("启动失败: %v", err)
		return
	}

	// 模拟运行一段时间后停止
	time.Sleep(1000 * time.Second)
	client.Stop()
	log.Println("测试完成")
}
