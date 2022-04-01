package avgstd

import (
	"math/big"
	"strconv"
	"testing"
)

func TestAvgStdBig(t *testing.T) {
	tests := []struct {
		inputs         []int
		mean           *big.Float
		variance       *big.Float
		sampleVariance *big.Float
	}{
		{[]int{1, 2, 3, 4, 5}, big.NewFloat(3), big.NewFloat(2), big.NewFloat(2.5)},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			a := AvgStdBig{}

			for _, v := range test.inputs {
				a.Update(big.NewFloat(float64(v)))
			}

			if a.Avg().Cmp(test.mean) != 0 {
				t.Fatalf("expected %v but found %v", test.mean, a.Avg())
			}
			if a.Variance().Cmp(test.variance) != 0 {
				t.Fatalf("expected %v but found %v", test.variance, a.Variance())
			}
			if a.SampledVariance().Cmp(test.sampleVariance) != 0 {
				t.Fatalf("expected %v but found %v", test.sampleVariance, a.SampledVariance())
			}
			if a.Samples().Uint64() != uint64(len(test.inputs)) {
				t.Fatalf("expected %v but found %v", len(test.inputs), a.Samples())
			}
		})
	}
}
