package controllor

import (
	"sync"
	"sync/atomic"
)

type Task func()

type Pool struct {
	maxWorkers int32          // 最大协程数
	taskQueue  chan Task      // 任务队列
	workers    int32          // 当前协程数
	wg         sync.WaitGroup // 协程等待组
	closed     int32          // 关闭标志（原子操作）
	quit       chan struct{}  // 关闭信号
}

func NewPool(maxWorkers, queueSize int) *Pool {
	return &Pool{
		maxWorkers: int32(maxWorkers),
		taskQueue:  make(chan Task, queueSize),
		quit:       make(chan struct{}),
	}
}

func (p *Pool) Submit(task Task) {
	if atomic.LoadInt32(&p.closed) == 1 {
		return // 已关闭不再接受新任务
	}

	select {
	case p.taskQueue <- task: // 尝试提交到缓冲队列
		p.tryStartWorker()
	case <-p.quit:
		return
	}
}

func (p *Pool) tryStartWorker() {
	for {
		current := atomic.LoadInt32(&p.workers)
		max := atomic.LoadInt32(&p.maxWorkers)

		if current >= max {
			return
		}

		// 原子操作创建协程
		if atomic.CompareAndSwapInt32(&p.workers, current, current+1) {
			p.wg.Add(1)
			go p.worker()
			return
		}
	}
}

func (p *Pool) worker() {
	defer atomic.AddInt32(&p.workers, -1)
	defer p.wg.Done()

	// 处理所有已入队任务（包括缓冲队列）
	for task := range p.taskQueue {
		task()

		// 动态缩容检查
		if atomic.LoadInt32(&p.workers) > atomic.LoadInt32(&p.maxWorkers) {
			return
		}
	}
}

func (p *Pool) SetMaxWorkers(n int) {
	atomic.StoreInt32(&p.maxWorkers, int32(n))
}

func (p *Pool) Close() {
	if !atomic.CompareAndSwapInt32(&p.closed, 0, 1) {
		return // 避免重复关闭
	}

	close(p.taskQueue) // 停止接受新任务，允许处理缓冲任务
	p.wg.Wait()        // 等待所有任务执行完毕
	close(p.quit)      // 正式关闭信号
}
