package businesspkg

type BusinessInterface interface {
	foo()
}

type Singleton struct{}

func (p *Singleton) foo() {}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

// 饿汉式
// GetInstance 获取实例
func GetInstance() BusinessInterface {
	return singleton
}
