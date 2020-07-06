package float

import (
	"github.com/shopspring/decimal"
)

func DecimalAdd(f1, f2 float64) (float64, bool) {
	a := decimal.NewFromFloat(f1)
	b := decimal.NewFromFloat(f2)
	return a.Add(b).Float64()
}
