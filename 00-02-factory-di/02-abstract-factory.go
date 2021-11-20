package factory

// FactoryMethod
type ProductInterface interface {
	Fruit() string
}

type HainanApple struct{}

func (HainanApple) Fruit() string { return "HainanApple" }

type HainanBanana struct{}

func (HainanBanana) Fruit() string { return "HainanBanana" }

type WuhanApple struct{}

func (WuhanApple) Fruit() string { return "WuhanApple" }

type WuhanBanana struct{}

func (WuhanBanana) Fruit() string { return "WuhanBanana" }

// AbstractFactory
type FruitInterface interface {
	ChooseApple() ProductInterface
	ChooseBanana() ProductInterface
}

type WuhanFruitFactory struct{}

func (WuhanFruitFactory) ChooseApple() ProductInterface {
	return WuhanApple{}
}

func (WuhanFruitFactory) ChooseBanana() ProductInterface {
	return WuhanBanana{}
}

type HainanFruitFactory struct{}

func (HainanFruitFactory) ChooseApple() ProductInterface {
	return HainanApple{}
}

func (HainanFruitFactory) ChooseBanana() ProductInterface {
	return HainanBanana{}
}
