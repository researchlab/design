package builder

import "fmt"

const (
	defaultMaxTotal = 10
	defaultMaxIdle  = 9
	defaultMinIdle  = 1
)

// ResourcePoolConfig resource pool
type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

// ResourcePoolConfigBuilder 用于构建 ResourcePoolConfig
type ResourcePoolConfigBuilder struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

// SetMinIdle set minIdle
func (p *ResourcePoolConfigBuilder) SetMinIdle(minIdle int) error {
	if minIdle < 0 {
		return fmt.Errorf("min idle cannot < 0, input: %d", minIdle)
	}
	p.minIdle = minIdle
	return nil
}

// SetMaxIdle set maxIdle
func (p *ResourcePoolConfigBuilder) SetMaxIdle(maxIdle int) error {
	if maxIdle <= 0 {
		return fmt.Errorf("max idle cannot <= 0, input: %d", maxIdle)
	}
	p.maxIdle = maxIdle
	return nil
}

// SetMaxTotal set maxTotal
func (p *ResourcePoolConfigBuilder) SetMaxTotal(maxTotal int) error {
	if maxTotal <= 0 {
		return fmt.Errorf("max total cannot <=0, input: %d", maxTotal)
	}
	p.maxTotal = maxTotal
	return nil
}

// Build build
func (p *ResourcePoolConfigBuilder) Build() (*ResourcePoolConfig, error) {
	if p.name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	if p.minIdle == 0 {
		p.minIdle = defaultMinIdle
	}

	if p.maxIdle == 0 {
		p.maxIdle = defaultMaxIdle
	}

	if p.maxTotal == 0 {
		p.maxTotal = defaultMaxTotal
	}

	if p.maxTotal < p.maxIdle {
		return nil, fmt.Errorf("max total(%d) cannot < max idle (%d)", p.maxTotal, p.maxIdle)
	}

	if p.minIdle > p.maxIdle {
		return nil, fmt.Errorf("max idle(%d) cannot < min idle (%d)", p.maxIdle, p.minIdle)
	}

	return &ResourcePoolConfig{
		name:     p.name,
		maxTotal: p.maxTotal,
		maxIdle:  p.maxIdle,
		minIdle:  p.minIdle,
	}, nil
}
