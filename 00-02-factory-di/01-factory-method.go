package factory

// 不同的对象实现同一个接口 就相当于是工厂方法模式了
type FruitFactoryMethod interface {
	Fruit() string
}

type appleF struct{}

func (*appleF) Fruit() string { return "apple" }

type bananaF struct{}

func (*bananaF) Fruit() string { return "banana" }
