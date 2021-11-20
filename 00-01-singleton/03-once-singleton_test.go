package businesspkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstanceOnce(t *testing.T) {
	assert.Equal(t, GetInstanceOnce(), GetInstanceOnce())
}

func BenchmarkGetInstanceOnceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstanceOnce() != GetInstanceOnce() {
				b.Error("test once fail")
			}
		}
	})
}
