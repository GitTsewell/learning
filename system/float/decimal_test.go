package float

import "testing"

func TestDecimalAdd(t *testing.T) {
	f, ok := DecimalAdd(0.1, 0.2)
	t.Log(f, ok)
}
