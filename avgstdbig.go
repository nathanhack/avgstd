package avgstd

import (
	"fmt"
	"math/big"
)

var zero = big.NewInt(0)
var one = big.NewInt(1)

//AvgStdBig holds the values for calculating average and variances
type AvgStdBig struct {
	Count *big.Int
	Mean  *big.Float
	S     *big.Float
}

//Update updates the state with a new value. Essentially adding a new value to be included in the calculations.
func (as *AvgStdBig) Update(value *big.Float) {
	if as.Count == nil || as.Mean == nil || as.S == nil {
		as.Reset()
	}
	as.Count.Add(as.Count, one)

	delta := new(big.Float).Sub(value, as.Mean)
	t := new(big.Float).Quo(delta, new(big.Float).SetInt(as.Count))
	as.Mean.Add(as.Mean, t)

	t = new(big.Float).Sub(value, as.Mean)
	t.Mul(t, delta)
	as.S.Add(as.S, t)
}

//Avg returns the current mean or average value
func (as *AvgStdBig) Avg() *big.Float {
	if as.Mean == nil {
		as.Reset()
	}
	return new(big.Float).Copy(as.Mean)
}

//Variance returns the current population variance.
func (as *AvgStdBig) Variance() *big.Float {
	if as.Count == nil || as.S == nil {
		as.Reset()
	}

	if as.Count.Cmp(zero) == 0 {
		return big.NewFloat(0)
	}

	return new(big.Float).Quo(as.S, new(big.Float).SetInt(as.Count))
}

//SampledVariance returns the current sampled variance.
func (as *AvgStdBig) SampledVariance() *big.Float {
	if as.Count == nil || as.S == nil {
		as.Reset()
	}
	if as.Count.Cmp(big.NewInt(2)) == 0 {
		return big.NewFloat(0)
	}
	t := new(big.Float).SetInt(as.Count)
	t.Sub(t, big.NewFloat(1))
	return new(big.Float).Quo(as.S, t)
}

//Samples returns the current number of samples included in the calculations.
func (as *AvgStdBig) Samples() *big.Int {
	if as.Count == nil {
		as.Reset()
	}
	return new(big.Int).Set(as.Count)
}

//Reset resets all internal values.
func (as *AvgStdBig) Reset() {
	as.Count = big.NewInt(0)
	as.Mean = big.NewFloat(0)
	as.S = big.NewFloat(0)
}

//String returns a string representation of this struct.
func (as AvgStdBig) String() string {
	return fmt.Sprintf("{Mean:%v Variance:%v Sampled:%v Samples:%v}", as.Mean, as.Variance(), as.SampledVariance(), as.Samples())
}
