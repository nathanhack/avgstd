package avgstd

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAvgStd(t *testing.T) {
	tests := []struct {
		inputs         []int
		mean           float64
		variance       float64
		sampleVariance float64
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2, 2.5},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			a := AvgStd{}

			for _, v := range test.inputs {
				a.Update(float64(v))
			}
			fmt.Println(a)
			if a.Avg() != test.mean {
				t.Fatalf("expected %v but found %v", test.mean, a.Avg())
			}
			if a.Variance() != test.variance {
				t.Fatalf("expected %v but found %v", test.variance, a.Variance())
			}
			if a.SampledVariance() != test.sampleVariance {
				t.Fatalf("expected %v but found %v", test.sampleVariance, a.SampledVariance())
			}
			if a.Samples() != len(test.inputs) {
				t.Fatalf("expected %v but found %v", len(test.inputs), a.Samples())
			}
		})
	}
}
