package businesspkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstanceDCL(t *testing.T) {
	assert.Equal(t, GetInstanceDCL(), GetInstanceDCL())
}

func BenchmarkGetInstanceDCLParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstanceDCL() != GetInstanceDCL() {
				b.Error("test dcl fail")
			}
		}
	})
}
