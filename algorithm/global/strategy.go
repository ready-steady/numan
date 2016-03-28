package global

import (
	"math"

	"github.com/ready-steady/adapt/algorithm/external"
	"github.com/ready-steady/adapt/algorithm/internal"
)

var (
	infinity = math.Inf(1.0)
)

// Strategy controls the interpolation process.
type Strategy interface {
	// First returns the initial state of the first iteration.
	First() *State

	// Check returns true if the stopping criteria have been satisfied.
	Check(*external.Progress) bool

	// Score assigns a score to an interpolation element.
	Score(*Element) float64

	// Next consumes the result of the current iteration and returns the initial
	// state of the next one.
	Next(*State, *external.Surrogate) *State
}

// BasicStrategy is a basic strategy satisfying the Strategy interface.
type BasicStrategy struct {
	internal.Active

	ni uint
	no uint

	grid Grid

	lmin uint
	lmax uint
	εa   float64
	εr   float64

	k uint

	global []float64
	local  []float64
	lower  []float64
	upper  []float64
}

// NewStrategy creates a basic strategy.
func NewStrategy(inputs, outputs uint, grid Grid, config *Config) *BasicStrategy {
	return &BasicStrategy{
		ni: inputs,
		no: outputs,

		grid: grid,

		lmin: config.MinLevel,
		lmax: config.MaxLevel,
		εa:   config.AbsoluteError,
		εr:   config.RelativeError,
	}
}

func (self *BasicStrategy) First() *State {
	self.Active = *internal.NewActive(self.ni)
	self.k = ^uint(0)
	self.lower = internal.Repeat(infinity, self.no)
	self.upper = internal.Repeat(-infinity, self.no)

	state := &State{}
	state.Lindices = self.Active.Next(self.k)
	state.Indices, state.Counts = internal.Index(self.grid, state.Lindices, self.ni)

	return state
}

func (self *BasicStrategy) Check(_ *external.Progress) bool {
	no := self.no
	nl := uint(len(self.local)) / no
	if nl == 0 {
		return false
	}
	δ := make([]float64, no)
	for i := uint(0); i < no; i++ {
		δ[i] = math.Max(self.εr*(self.upper[i]-self.lower[i]), self.εa)
	}
	for i := range self.Positions {
		if i >= nl {
			continue
		}
		for j := uint(0); j < no; j++ {
			if self.local[i*no+j] > δ[j] {
				return false
			}
		}
	}
	return true
}

func (self *BasicStrategy) Score(element *Element) float64 {
	score := 0.0
	for _, value := range element.Surpluses {
		score += math.Abs(value)
	}
	score /= float64(uint(len(element.Observations)) / self.no)
	return score
}

func (self *BasicStrategy) Next(current *State, _ *external.Surrogate) *State {
	self.consume(current)

	self.Active.Drop(self.k)
	if len(self.Positions) == 0 {
		return nil
	}
	self.k = internal.LocateMax(self.global, self.Positions)
	if self.global[self.k] <= 0.0 {
		return nil
	}

	state := &State{}
	state.Lindices = self.Active.Next(self.k)
	state.Indices, state.Counts = internal.Index(self.grid, state.Lindices, self.ni)

	return state
}

func (self *BasicStrategy) consume(state *State) {
	no, ng, nl := self.no, uint(len(self.global)), uint(len(self.local))
	nn := uint(len(state.Counts))

	levels := internal.Levelize(state.Lindices, self.ni)

	self.global = append(self.global, state.Scores...)
	global := self.global[ng:]

	self.local = append(self.local, make([]float64, nn*no)...)
	local := self.local[nl:]

	for i, o := uint(0), uint(0); i < nn; i++ {
		count := state.Counts[i]
		if levels[i] < uint64(self.lmin) {
			global[i] = infinity
			for j := uint(0); j < no; j++ {
				local[i*no+j] = infinity
			}
		} else if levels[i] >= uint64(self.lmax) {
			global[i] = -infinity
			for j := uint(0); j < no; j++ {
				local[i*no+j] = -infinity
			}
		} else {
			for j, m := uint(0), count*no; j < m; j++ {
				k := i*no + j%no
				local[k] = math.Max(local[k], math.Abs(state.Surpluses[o+j]))
			}
		}
		o += count
	}

	for i, point := range state.Observations {
		j := uint(i) % no
		self.lower[j] = math.Min(self.lower[j], point)
		self.upper[j] = math.Max(self.upper[j], point)
	}
}
