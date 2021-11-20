package factory

import "testing"

func TestFruitFactoryMethod(t *testing.T) {
	tests := []struct {
		name string
		args FruitFactoryMethod
		want string
	}{
		{
			name: "apple",
			args: &appleF{},
			want: "apple",
		},
		{
			name: "banana",
			args: &bananaF{},
			want: "banana",
		},
	}

	for _, tt := range tests {
		if got := tt.args.Fruit(); got != tt.want {
			t.Errorf("FruitFactoryMethod = %v, want %v", got, tt.want)
		}
	}
}
