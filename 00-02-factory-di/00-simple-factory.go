package factory

type FruitFactory interface {
	Fruit() string
}

func NewFruit(t string) FruitFactory {
	switch t {
	case "apple":
		return &apple{}
	case "banana":
		return &banana{}
	}
	return nil
}

type apple struct{}

func (p *apple) Fruit() string { return "apple" }

type banana struct{}

func (p *banana) Fruit() string { return "banana" }
