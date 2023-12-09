package avgstd

import (
	"fmt"
	"math"
)

// Avg holds the values for calculating average
type Avg struct {
	Count uint64
	Mean  float64
}

// Update updates the state with a new value. Essentially adding a new value to be included in the calculations.
func (as *Avg) Update(value float64) {
	if as.Count == math.MaxUint64 {
		panic("update limit exceeded")
	}
	as.Count++
	delta := value - as.Mean
	as.Mean += delta / float64(as.Count)
}

// Avg returns the current mean or average value
func (as *Avg) Avg() float64 {
	return as.Mean
}

// Samples returns the current number of samples included in the calculations.
func (as *Avg) Samples() uint64 {
	return as.Count
}

// Reset resets all internal values.
func (as *Avg) Reset() {
	as.Count = 0
	as.Mean = 0
}

// String returns a string representation of this struct.
func (as Avg) String() string {
	return fmt.Sprintf("{Mean:%v}", as.Mean)
}
