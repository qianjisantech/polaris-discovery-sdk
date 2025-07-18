package main

import (
	"log"
	"polaris-discovery/discovery"
	"polaris-discovery/model"
	"time"
)

func main() {
	client := &discovery.DiscoveryClient{Addr: "http://localhost:8080"}

	err := client.Start(
		func(resp *model.RegisterResponse) {
			log.Printf("===============================注册成功回调方法! ID: %s", resp.Data.Id)
		},
		func(err error) {
			log.Printf("===============================注册失败回调方法: %v", err)
		},
		func(resp *model.HeatBeatResponse) {
			log.Printf("===============================心跳成功回调方法 id: %s", resp.Data.Id)
		},
		func(err error) {
			log.Printf("===============================心跳失败回调方法 : %v", err)
		},
	)

	if err != nil {
		log.Fatal(err)
	}
	// 模拟运行一段时间后停止
	time.Sleep(30 * time.Second)
	client.Stop()
}
