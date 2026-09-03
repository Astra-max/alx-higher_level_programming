package libs

import (
	"errors"
	"math"
)

func AddNums(a, b int) (int, error) {

	if b > math.MaxInt  {
		return 0, errors.New("Overflow int")
	}
	if a < math.MinInt {
		return 0, errors.New("underflow")
	}
	return a + b, nil
}