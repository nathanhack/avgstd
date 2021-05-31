package avgstd

import (
	"fmt"
)

//AvgStd holds the values for calculating average and variances
type AvgStd struct {
	Count int
	Mean  float64
	S     float64
}

//Update updates the state with a new value. Essentially adding a new value to be included in the calculations.
func (as *AvgStd) Update(value float64) {
	as.Count++
	delta := value - as.Mean
	as.Mean += delta / float64(as.Count)
	as.S += (value - as.Mean) * delta
}

//Avg returns the current mean or average value
func (as *AvgStd) Avg() float64 {
	return as.Mean
}

//Variance returns the current population variance.
func (as *AvgStd) Variance() float64 {
	if as.Count == 0 {
		return 0
	}
	return as.S / float64(as.Count)
}

//SampledVariance returns the current sampled variance.
func (as *AvgStd) SampledVariance() float64 {
	if as.Count < 2 {
		return 0
	}
	return as.S / float64(as.Count-1)
}

//Samples returns the current number of samples included in the calculations.
func (as *AvgStd) Samples() int {
	return as.Count
}

//Reset resets all internal values.
func (as *AvgStd) Reset() {
	as.Count = 0
	as.Mean = 0
	as.S = 0
}

//String returns a string representation of this struct.
func (as AvgStd) String() string {
	return fmt.Sprintf("{Mean:%v Variance:%v Sampled:%v Samples:%v}", as.Mean, as.Variance(), as.SampledVariance(), as.Samples())
}
