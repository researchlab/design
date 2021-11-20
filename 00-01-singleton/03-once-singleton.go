package businesspkg

import "sync"

type BusinessOnce interface {
	foo()
}

type SingletonOnce struct{}

func (p *SingletonOnce) foo() {}

var (
	instanceOnce *SingletonOnce
	once         sync.Once
)

// 使用golang 提供的sync.Once 单例模式 更简洁
func GetInstanceOnce() BusinessOnce {
	once.Do(func() {
		instanceOnce = new(SingletonOnce)
	})
	return instanceOnce
}
