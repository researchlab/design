package factory

import "testing"

func TestFruit(t *testing.T) {
	tests := []struct {
		name string
		args ProductInterface
		want string
	}{
		{
			name: "WuhanApple",
			args: (WuhanFruitFactory{}).ChooseApple(),
			want: "WuhanApple",
		},
		{
			name: "WuhanBanana",
			args: (WuhanFruitFactory{}).ChooseBanana(),
			want: "WuhanBanana",
		},
		{
			name: "HainanApple",
			args: (HainanFruitFactory{}).ChooseApple(),
			want: "HainanApple",
		},
		{
			name: "HainanBanana",
			args: (HainanFruitFactory{}).ChooseBanana(),
			want: "HainanBanana",
		},
	}

	for _, tt := range tests {
		if got := tt.args.Fruit(); got != tt.want {
			t.Errorf("Fruit = %v, want %v", got, tt.want)
		}
	}
}
