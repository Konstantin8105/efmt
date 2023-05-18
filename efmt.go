package efmt

import (
	"fmt"
	"math"
)

const (
	exponent = 4
	size     = 6
)

// Sprintf convert `float64` in engineer format
func Sprint(v float64) string {
	exp := math.Log10(math.Abs(v))
	fl := int(math.Floor(exp))
	ost := fl % 3
	fl = fl - ost
	var format string
	if fl < 0 {
		fl -= 3
		format = fmt.Sprintf("%%%d.%df%%+0%dd", size-1-ost, size-4-ost, exponent)
	} else if fl == 0 {
		if ost < 0 {
			fl -= 3
			format = fmt.Sprintf("%%%d.%df%%+0%dd", size-1-ost, size-4-ost, exponent)
		} else {
			format = fmt.Sprintf("%%.%df", size-1-ost)
			return fmt.Sprintf(format, v)
		}
	} else {
		format = fmt.Sprintf("%%.%df%%+0%dd", size-1-ost, exponent)
	}
	v = v / math.Pow(10, float64(fl))
	return fmt.Sprintf(format, v, fl)
}
