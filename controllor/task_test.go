package controllor

import (
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	pool := NewPool(3, 100)
	defer pool.Close()

	// 动态调整协程数量
	pool.SetMaxWorkers(5)

	// 提交任务
	for i := 0; i < 10; i++ {
		id := i
		pool.Submit(func() {
			time.Sleep(1 * time.Second)
			println("执行任务", id)
		})
	}
}
