package factory

import (
	"reflect"
	"testing"
)

func TestNewFruit(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want FruitFactory
	}{
		{
			name: "banana",
			args: args{t: "banana"},
			want: &banana{},
		},
		{
			name: "apple",
			args: args{t: "apple"},
			want: &apple{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFruit((tt.args.t)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFruit() = %v, want %v", got, tt.want)
			}
		})
	}
}
