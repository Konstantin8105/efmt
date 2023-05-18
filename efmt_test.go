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

// cpu: Intel(R) Xeon(R) CPU E3-1240 V2 @ 3.40GHz
// Benchmark/-011-4         	 1393530	       843.0 ns/op	      48 B/op	       4 allocs/op
// Benchmark/+002-4         	 1626250	       733.1 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+011-4         	 1539116	       754.9 ns/op	      40 B/op	       3 allocs/op
func Benchmark(b *testing.B) {
	for _, exp := range []int{-11, 2, 11} {
		b.Run(fmt.Sprintf("%+04d", exp), func(b *testing.B) {
			value := 1.234567890 * math.Pow(10, float64(exp))
			b.ResetTimer()

			for n := 0; n < b.N; n++ {
				_ = efmt.Sprint(value)
			}
		})
	}
}
