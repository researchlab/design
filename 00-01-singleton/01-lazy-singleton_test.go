package businesspkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstanceLazy(t *testing.T) {
	assert.Equal(t, GetInstanceLazy(), GetInstanceLazy())
}

func BenckmarkGetInstanceLazyParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstanceLazy() != GetInstanceLazy() {
				b.Error("test lazy fail")
			}
		}
	})
}
