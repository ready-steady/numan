package polynomial

// Closed is a basis in [0, 1]^n.
type Closed struct {
	nd uint
	np uint
}

// NewClosed creates a basis in [0, 1]^n.
func NewClosed(dimensions uint, order uint) *Closed {
	return &Closed{dimensions, order}
}

// Compute evaluates a basis function.
func (self *Closed) Compute(index []uint64, point []float64) float64 {
	nd, np := self.nd, uint64(self.np)

	value := 1.0
	for i := uint(0); i < nd; i++ {
		level, power := levelMask&index[i], np
		if level < power {
			power = level
		}
		if power == 0 {
			continue // value *= 1.0
		}

		order := index[i] >> levelSize

		x := point[i]
		xi, step := node(level, order)
		delta := x - xi
		if delta < 0.0 {
			delta = -delta
		}
		if delta >= step {
			return 0.0 // value *= 0.0
		}

		if power == 1 {
			value *= 1.0 - delta/step
			continue
		}

		xl, xr := xi-step, xi+step
		value *= (x - xl) / (xi - xl)
		value *= (x - xr) / (xi - xr)

		level, order = parent(level, order)
		for j := uint64(3); j < power; j++ {
			level, order = parent(level, order)
			xj, _ := node(level, order)
			value *= (x - xj) / (xi - xj)
		}
	}

	return value
}

// Integrate computes the integral of a basis function.
func (_ *Closed) Integrate(_ []uint64) float64 {
	return 0.0
}

func node(level, order uint64) (x, step float64) {
	if level == 0 {
		x, step = 0.5, 1.0
	} else {
		step = 1.0 / float64(uint64(2)<<(level-1))
		x = step * float64(order)
	}
	return
}

func parent(level, order uint64) (uint64, uint64) {
	switch level {
	case 0:
		panic("the root does not have a parent")
	case 1:
		level = 0
		order = 0
	case 2:
		level = 1
		order -= 1
	default:
		level -= 1
		if ((order-1)/2)%2 == 0 {
			order = (order + 1) / 2
		} else {
			order = (order - 1) / 2
		}
	}
	return level, order
}
