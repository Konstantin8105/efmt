package efmt_test

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/Konstantin8105/compare"
	"github.com/Konstantin8105/efmt"
)

func Test(t *testing.T) {
	value := 1.234567890
	var vs []float64
	for i := -12; i < 12; i++ {
		if i == 0 {
			vs = append(vs, 0.0)
		}
		vs = append(vs, value*math.Pow(10, float64(i)))
		vs = append(vs, -vs[len(vs)-1])
	}
	const exponent = 4
	const size = 6
	var buf bytes.Buffer
	for i, v := range vs {
		s := efmt.Sprint(v)
		fmt.Fprintf(&buf, "%20s %20e %20f\n", s, vs[i], vs[i])
		t.Run(s, func(t *testing.T) {
			index := strings.Index(s, "e")
			if 0 <= index {
				s = s[:index]
			}
			s = strings.ReplaceAll(s, "+", "")
			s = strings.ReplaceAll(s, "-", "")
			s = strings.ReplaceAll(s, ".", "")
			if 6 < len(s) {
				t.Errorf("not valid lenght of number: `%s`", s)
			}
		})
	}
	compare.Test(t, ".fts", buf.Bytes())
}

// cpu: Intel(R) Xeon(R) CPU E3-1240 V2 @ 3.40GHz
// Benchmark/-011-4         	 1393530	       843.0 ns/op	      48 B/op	       4 allocs/op
// Benchmark/+002-4         	 1626250	       733.1 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+011-4         	 1539116	       754.9 ns/op	      40 B/op	       3 allocs/op
// Benchmark/fmt.Sprintf-4  	 4528094	       265.0 ns/op	      32 B/op	       2 allocs/op
//
// Benchmark/-011-4         	 1307503	       913.0 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-010-4         	 1336034	       899.8 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-009-4         	 1329030	       899.5 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-008-4         	 1304426	       909.1 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-007-4         	 1294702	       915.4 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-006-4         	 1330118	       946.5 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-005-4         	 1250877	       970.9 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-004-4         	 1264894	       971.3 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-003-4         	 1255070	       909.2 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-002-4         	 1282087	       936.8 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-001-4         	 2033076	       559.2 ns/op	      24 B/op	       3 allocs/op
// Benchmark/+000-4         	 2140952	       545.7 ns/op	      24 B/op	       3 allocs/op
// Benchmark/+001-4         	 2161412	       542.6 ns/op	      24 B/op	       3 allocs/op
// Benchmark/+002-4         	 2157922	       547.7 ns/op	      24 B/op	       3 allocs/op
// Benchmark/+003-4         	 2227048	       535.1 ns/op	      24 B/op	       3 allocs/op
// Benchmark/+004-4         	 2241903	       523.9 ns/op	      24 B/op	       3 allocs/op
// Benchmark/+005-4         	 2152898	       575.6 ns/op	      32 B/op	       3 allocs/op
// Benchmark/+006-4         	 1317602	       866.2 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+007-4         	 1403604	       839.2 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+008-4         	 1410225	       863.8 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+009-4         	 1286062	       855.4 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+010-4         	 1439583	       825.4 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+011-4         	 1461277	       828.6 ns/op	      40 B/op	       3 allocs/op
// Benchmark/fmt.Sprintf-4  	 4527307	       262.3 ns/op	      32 B/op	       2 allocs/op
//
// cpu: Intel(R) Xeon(R) CPU           X5550  @ 2.67GHz
// Benchmark/-011-16         	 1000000	      1153 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-010-16         	 1000000	      1118 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-009-16         	 1000000	      1097 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-008-16         	 1000000	      1115 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-007-16         	 1000000	      1094 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-006-16         	 1000000	      1069 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-005-16         	 1000000	      1114 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-004-16         	  996427	      1138 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-003-16         	 1000000	      1078 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-002-16         	 1000000	      1080 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-001-16         	 1000000	      1145 ns/op	      48 B/op	       4 allocs/op
// Benchmark/+000-16         	 1559452	       736.5 ns/op	      24 B/op	       3 allocs/op
// Benchmark/+001-16         	 1664338	       724.3 ns/op	      24 B/op	       3 allocs/op
// Benchmark/+002-16         	 1704063	       727.6 ns/op	      24 B/op	       3 allocs/op
// Benchmark/+003-16         	 1000000	      1019 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+004-16         	 1000000	      1006 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+005-16         	 1000000	      1013 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+006-16         	 1000000	      1015 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+007-16         	 1000000	      1037 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+008-16         	 1000000	      1077 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+009-16         	 1164051	      1022 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+010-16         	 1000000	      1026 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+011-16         	 1000000	      1016 ns/op	      40 B/op	       3 allocs/op
// Benchmark/fmt.Sprintf-16  	 3964371	       302.5 ns/op	      32 B/op	       2 allocs/op
//
// Benchmark/-011-16         	  918274	      1185 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-010-16         	 1000000	      1167 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-009-16         	 1000000	      1184 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-008-16         	 1000000	      1171 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-007-16         	 1000000	      1144 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-006-16         	 1000000	      1136 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-005-16         	 1000000	      1153 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-004-16         	 1000000	      1126 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-003-16         	 1000000	      1113 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-002-16         	 1000000	      1155 ns/op	      48 B/op	       4 allocs/op
// Benchmark/-001-16         	 1749850	       685.4 ns/op	      24 B/op	       3 allocs/op
// Benchmark/+000-16         	 1737892	       709.0 ns/op	      24 B/op	       3 allocs/op
// Benchmark/+001-16         	 1652661	       691.5 ns/op	      24 B/op	       3 allocs/op
// Benchmark/+002-16         	 1781551	       697.7 ns/op	      24 B/op	       3 allocs/op
// Benchmark/+003-16         	 1755388	       671.0 ns/op	      24 B/op	       3 allocs/op
// Benchmark/+004-16         	 1830135	       648.3 ns/op	      24 B/op	       3 allocs/op
// Benchmark/+005-16         	 1000000	      1043 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+006-16         	 1000000	      1119 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+007-16         	 1000000	      1115 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+008-16         	 1000000	      1104 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+009-16         	 1000000	      1148 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+010-16         	 1000000	      1059 ns/op	      40 B/op	       3 allocs/op
// Benchmark/+011-16         	  973624	      1082 ns/op	      40 B/op	       3 allocs/op
// Benchmark/fmt.Sprintf-16  	 3304615	       325.6 ns/op	      32 B/op	       2 allocs/op
func Benchmark(b *testing.B) {
	for exp := -11; exp <= 11; exp++ {
		b.Run(fmt.Sprintf("%+04d", exp), func(b *testing.B) {
			value := 1.234567890 * math.Pow(10, float64(exp))
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = efmt.Sprint(value)
			}
		})
	}
	b.Run("fmt.Sprintf", func(b *testing.B) {
		value := 1.234567890 * math.Pow(10, float64(-11))
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprint(value)
		}
	})
}
