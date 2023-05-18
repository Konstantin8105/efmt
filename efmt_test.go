package efmt_test

import (
	"bytes"
	"fmt"
	"math"
	"testing"

	"github.com/Konstantin8105/compare"
	"github.com/Konstantin8105/efmt"
)

func Test(t *testing.T) {
	var vs []float64
	for i := -12; i < 12; i++ {
		vs = append(vs, 1.234567890*math.Pow(10, float64(i)))
		vs = append(vs, -vs[len(vs)-1])
	}
	const exponent = 4
	const size = 6
	var buf bytes.Buffer
	for i, v := range vs {
		s := efmt.Sprint(v)
		fmt.Fprintf(&buf, "%20s %20e\n", s, vs[i])
	}
	compare.Test(t, ".fts", buf.Bytes())
}
