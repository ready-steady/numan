package global

import (
	"math"

	"github.com/ready-steady/adapt/algorithm/internal"
)

var (
	infinity = math.Inf(1.0)
)

func index(grid Grid, lindices []uint64, ni uint) ([]uint64, []uint) {
	nn := uint(len(lindices)) / ni
	indices, counts := []uint64(nil), make([]uint, nn)
	for i := uint(0); i < nn; i++ {
		newIndices := grid.Index(lindices[:ni])
		indices = append(indices, newIndices...)
		counts[i] = uint(len(newIndices)) / ni
		lindices = lindices[ni:]
	}
	return indices, counts
}

func score(basis Basis, strategy Strategy, target Target, counts []uint, indices []uint64,
	values, surpluses []float64, ni, no uint) {

	for _, count := range counts {
		oi, oo := count*ni, count*no
		location := Location{
			Indices:   indices[:oi],
			Volumes:   internal.Measure(basis, indices[:oo], ni),
			Values:    values[:oo],
			Surpluses: surpluses[:oo],
		}
		strategy.Push(&location, target.Score(&location))
		indices, values, surpluses = indices[oi:], values[oo:], surpluses[oo:]
	}
}
