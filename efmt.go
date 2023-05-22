package efmt

import (
	"fmt"
	"math"
)

const (
	exponent = 3
	size     = 6
)

// Sprintf convert `float64` in engineer format
func Sprint(v float64) string {
	av := math.Abs(v)
	if v == 0.0 {
		return "0.00000"
	}
	switch {
	case 0.1 <= av && av < 1.0:
		format := fmt.Sprintf("%%.%df", size-1)
		return fmt.Sprintf(format, v)
	case 1.0 <= av && av < 10.0:
		format := fmt.Sprintf("%%.%df", size-1)
		return fmt.Sprintf(format, v)
	case 10.0 <= av && av < 100.0:
		format := fmt.Sprintf("%%.%df", size-2)
		return fmt.Sprintf(format, v)
	case 100.0 <= av && av < 1000.0:
		format := fmt.Sprintf("%%.%df", size-3)
		return fmt.Sprintf(format, v)
	case 1000.0 <= av && av < 10000.0:
		format := fmt.Sprintf("%%.%df", size-4)
		return fmt.Sprintf(format, v)
	case 10000.0 <= av && av < 100000.0:
		format := fmt.Sprintf("%%.%df", size-5)
		return fmt.Sprintf(format, v)
	case 100000.0 <= av && av < 1000000.0:
		format := fmt.Sprintf("%%%d.f.", size)
		return fmt.Sprintf(format, v)
	}
	exp := math.Log10(math.Abs(v))
	fl := int(math.Floor(exp))
	ost := fl % 3
	fl = fl - ost
	var format string
	if fl < 0 {
		if ost == 0 {
			format = fmt.Sprintf("%%%d.%dfe%%+0%dd", size-1+ost, size-1-ost, exponent)
		} else {
			fl -= 3
			format = fmt.Sprintf("%%%d.%dfe%%+0%dd", size-1+ost, size-4-ost, exponent)
		}
	} else if fl == 0 {
		if ost < 0 {
			fl -= 3
			format = fmt.Sprintf("%%%d.%dfe%%+0%dd", size-1-ost, size-4-ost, exponent)
		} else {
			format = fmt.Sprintf("%%.%df", size-1-ost)
			return fmt.Sprintf(format, v)
		}
	} else {
		format = fmt.Sprintf("%%.%dfe%%+0%dd", size-1-ost, exponent)
	}
	v = v / math.Pow(10, float64(fl))
	return fmt.Sprintf(format, v, fl)
}
