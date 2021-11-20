package businesspkg

import "sync"

type BusinessLazy interface {
	foo()
}

type SingleLazy struct{}

func (p *SingleLazy) foo() {

}

var (
	lock         sync.Mutex
	instanceLazy *SingleLazy
)

// 懒汉式 instance在需要的时候才创建
// 缺点并发度低，性能瓶颈
func GetInstanceLazy() BusinessLazy {
	lock.Lock()
	defer lock.Unlock()
	if instanceLazy == nil {
		instanceLazy = new(SingleLazy)
	}
	return instanceLazy
}
