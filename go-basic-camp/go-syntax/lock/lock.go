package lock

import "sync"

// LockDemo
// 优先使用 RWMutex，优先加读锁
// 常用的并发优化手段，用读写锁来优化读锁

type LockDemo struct {
	lock sync.Mutex
}

func NewLockDemo() *LockDemo {
	return &LockDemo{}
}

func (l *LockDemo) PanicDemo() {
	l.lock.Lock()
	// 在中间 panic 了，无法释放锁
	panic("abc")
	l.lock.Unlock()
}

func (l *LockDemo) DeferDemo() {
	l.lock.Lock()
	defer l.lock.Unlock()
}

func (l LockDemo) NoPointerDemo() {
	l.lock.Lock()
	defer l.lock.Unlock()
}

type LockDemoV1 struct {
	lock *sync.Mutex
}

func NewLockDemoV1() LockDemoV1 {
	return LockDemoV1{
		// 如果不初始化，lock 就是 nil，你一用就 panic
		lock: &sync.Mutex{},
	}
}

func (l LockDemoV1) NoPointerDemo() {
	l.lock.Lock()
	defer l.lock.Unlock()
}
