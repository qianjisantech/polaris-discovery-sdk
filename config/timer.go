package config

import (
	"context"
	"log"
	"time"
)

type Timer struct {
	ticker *time.Ticker
	done   chan struct{}
}

// NewTimer 创建一个定时器，并返回 *time.Ticker（需调用方保存引用）
func NewTimer(ctx context.Context, interval time.Duration, task func()) *time.Ticker {
	ticker := time.NewTicker(interval)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				log.Println("定时器触发，执行任务...")
				task()
			case <-ctx.Done():
				log.Println("定时器停止：上下文取消")
				return
			}
		}
	}()
	return ticker
}

func (t *Timer) run(ctx context.Context, task func()) {
	defer func() {
		t.ticker.Stop()
		if r := recover(); r != nil {
			log.Printf("timer panic: %v", r)
		}
	}()

	for {
		select {
		case <-t.ticker.C:
			task()
		case <-t.done:
			return
		case <-ctx.Done():
			return
		}
	}
}

func (t *Timer) Stop() {
	if t.ticker != nil {
		t.ticker.Stop()
		log.Println("定时器已停止")
	}
}
