package avgstd

import (
	"fmt"
	"math"
)

// AvgStdMinMax holds the values for calculating average and variances
type AvgStdMinMax struct {
	Count uint64
	Mean  float64
	S     float64
	Min   float64
	Max   float64
}

// Update updates the state with a new value. Essentially adding a new value to be included in the calculations.
func (as *AvgStdMinMax) Update(value float64) {
	if as.Count == math.MaxInt64 {
		panic("update limit exceeded")
	}
	as.Count++
	delta := value - as.Mean
	as.Mean += delta / float64(as.Count)
	as.S += (value - as.Mean) * delta
	as.Min = min(as.Min, value)
	as.Max = max(as.Max, value)
}

// Avg returns the current mean or average value
func (as *AvgStdMinMax) Avg() float64 {
	return as.Mean
}

// Variance returns the current population variance.
func (as *AvgStdMinMax) Variance() float64 {
	if as.Count == 0 {
		return 0
	}
	return as.S / float64(as.Count)
}

// SampledVariance returns the current sampled variance.
func (as *AvgStdMinMax) SampledVariance() float64 {
	if as.Count < 2 {
		return 0
	}
	return as.S / float64(as.Count-1)
}

// Samples returns the current number of samples included in the calculations.
func (as *AvgStdMinMax) Samples() uint64 {
	return as.Count
}

// Reset resets all internal values.
func (as *AvgStdMinMax) Reset() {
	as.Count = 0
	as.Mean = 0
	as.S = 0
	as.Min = 0
	as.Max = 0
}

// String returns a string representation of this struct.
func (as AvgStdMinMax) String() string {
	return fmt.Sprintf("{Mean:%v Variance:%v Sampled:%v Samples:%v Min:%v Max:%v}", as.Mean, as.Variance(), as.SampledVariance(), as.Samples(), as.Min, as.Max)
}
