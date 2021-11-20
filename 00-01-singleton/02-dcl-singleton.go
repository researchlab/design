package businesspkg

import "sync"

type BusinessDoubleCheckedLocking interface {
	foo()
}

type SingletonDCL struct{}

func (p *SingletonDCL) foo() {

}

var (
	lockDCL     sync.Mutex
	instanceDCL *SingletonDCL
)

// doublecheckedlocking
// 是指 第一次判断不加锁
// 第二次加锁保证线程安全
// 一旦对象建立后，获取对象就不需要加锁了
// doublechecked 指两次检测instanceDCL 是否为nil
func GetInstanceDCL() BusinessDoubleCheckedLocking {
	if instanceDCL == nil {
		lockDCL.Lock()
		defer lockDCL.Unlock()
		if instanceDCL == nil {
			instanceDCL = new(SingletonDCL)
		}
	}
	return instanceDCL
}
