package core

import (
	"context"
	"fmt"
	"github.com/qianjisantech/polaris-discovery-sdk/config"
	"log"
	"math"
	"sync"
	"time"
)

// DiscoveryClient 注册中心Client结构体
type DiscoveryClient struct {
	Addr              string
	ticker            *time.Ticker // 保存定时器引用
	done              chan struct{}
	HeartbeatInterval int        //心跳间隔 s
	stopOnce          sync.Once  // 确保Stop只执行一次
	stopped           bool       // 标记是否已停止
	Timeout           int        // 超时时间间隔 s
	mu                sync.Mutex // 保护并发访问
	Retry             Retry
}

type Retry struct {
	MaxAttempts int // 操作失败后的最大重试次数（包含首次尝试） 最大重试次数  默认3次
	Backoff     int //控制重试间隔的基数，实际间隔会随尝试次数指数增长（指数退避算法  重试间隔基数（单位通常为ms） 默认500ms
}

// NewDiscoveryClient 注册中心Client
func (r *DiscoveryClient) newDiscoveryClient() *DiscoveryClient {
	// 复制当前配置
	client := &DiscoveryClient{
		Addr:              r.Addr,
		done:              make(chan struct{}),
		HeartbeatInterval: r.HeartbeatInterval,
		Timeout:           r.Timeout,
		Retry:             r.Retry,
	}

	// 设置默认重试参数
	if client.Retry.MaxAttempts <= 0 {
		client.Retry.MaxAttempts = 3
	}
	if client.Retry.MaxAttempts > 100 {
		client.Retry.MaxAttempts = 100
	}
	if client.Retry.Backoff <= 0 {
		client.Retry.Backoff = 500 //默认重试间隔500ms
	}
	if client.Retry.Backoff > 500*2000 {
		client.Retry.Backoff = 500 * 2000 //最大重试间隔500 * 2000ms
	}

	return client
}

// registerWithRetry 带重试机制的注册方法
func (r *DiscoveryClient) registerWithRetry() (*RegisterResponse, error) {
	var lastErr error
	var resp *RegisterResponse

	for attempt := 0; attempt < r.Retry.MaxAttempts; attempt++ {
		select {
		case <-r.done:
			return nil, fmt.Errorf("客户端在注册完成前已停止")
		default:
			currentAttempt := attempt + 1
			log.Printf("开始第 %d/%d 次注册尝试", currentAttempt, r.Retry.MaxAttempts)

			resp, lastErr = r.register()

			if lastErr != nil {
				log.Printf("注册调用失败 (第 %d/%d 次尝试), 错误: %v",
					currentAttempt, r.Retry.MaxAttempts, lastErr)
			} else if resp == nil {
				lastErr = fmt.Errorf("收到空响应")
			} else if !resp.Success {
				lastErr = fmt.Errorf("服务返回失败: %s", resp.Message)
			} else {
				return resp, nil
			}

			if currentAttempt < r.Retry.MaxAttempts {
				backoff := time.Duration(math.Pow(2, float64(attempt))) *
					time.Duration(r.Retry.Backoff) * time.Millisecond
				log.Printf("等待 %v 后重试...", backoff)

				select {
				case <-time.After(backoff):
					continue
				case <-r.done:
					return nil, fmt.Errorf("客户端在注册重试过程中停止")
				}
			}
		}
	}

	// 异步触发停止，避免死锁
	go r.Stop()

	return resp, fmt.Errorf("注册失败，已达到最大重试次数 %d，最后错误: %v",
		r.Retry.MaxAttempts, resp.Message)
}

func (r *DiscoveryClient) Start(
	onRegisterSuccess func(*RegisterResponse),
	onRegisterError func(error),
	onHeartbeatSuccess func(*HeatBeatResponse),
	onHeartbeatError func(error),
) error {
	r.mu.Lock()
	if r.stopped {
		r.mu.Unlock()
		return fmt.Errorf("客户端已停止")
	}
	r.mu.Unlock()
	client := r.newDiscoveryClient()
	// 注册过程不需要持有锁
	registerResp, err := client.registerWithRetry()
	if err != nil {
		if onRegisterError != nil {
			onRegisterError(err)
		}
		return fmt.Errorf("注册过程失败: %w", err)
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	// 设置心跳间隔
	if r.HeartbeatInterval <= 0 {
		r.HeartbeatInterval = 5
	}
	if r.HeartbeatInterval > 60*5 {
		r.HeartbeatInterval = 60 * 5
	}

	if registerResp.Success {
		if onRegisterSuccess != nil {
			onRegisterSuccess(registerResp)
		}
		log.Printf("注册成功----->%v", registerResp)

		if registerResp.Data.IdentificationCode != "" {
			r.ticker = config.NewTimer(context.Background(), time.Duration(r.HeartbeatInterval)*time.Second, func() {
				select {
				case <-r.done:
					return // 已停止，不再执行心跳
				default:
					log.Printf("-----开始执行心跳任务...")
					resp, err := r.heatBeat(registerResp.Data.IdentificationCode)
					if err != nil {
						log.Printf("心跳失败----->%v", err)
						if onHeartbeatError != nil {
							onHeartbeatError(err)
						}
						return
					}
					if resp.Success {
						log.Printf("心跳成功-----> %v", resp.Message)
						if onHeartbeatSuccess != nil {
							onHeartbeatSuccess(resp)
						}
					} else {
						log.Printf("心跳异常----->  %v", resp.Message)
						if onHeartbeatError != nil {
							onHeartbeatError(fmt.Errorf(resp.Message))
						}
					}
				}
			})
		}
	} else {
		err := fmt.Errorf("服务注册异常: %s", registerResp.Message)
		if onRegisterError != nil {
			onRegisterError(err)
		}
		return err
	}
	return nil
}

// Stop 方法改进版
func (r *DiscoveryClient) Stop() {
	r.stopOnce.Do(func() {
		// 尝试获取锁，最多等待100ms
		ok := make(chan struct{})
		go func() {
			r.mu.Lock()
			close(ok)
		}()

		select {
		case <-ok:
			defer r.mu.Unlock()
			if r.stopped {
				return
			}
			r.stopped = true
			if r.ticker != nil {
				r.ticker.Stop()
				r.ticker = nil
			}
			close(r.done)
			log.Println("客户端已成功停止")
		case <-time.After(100 * time.Millisecond):
			log.Println("警告：停止超时，可能已在停止过程中")
		}
	})
}

// IsStopped 检查客户端是否已停止
func (r *DiscoveryClient) IsStopped() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.stopped
}
