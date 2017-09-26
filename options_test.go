package igdb

import "testing"

func TestOptOrder(t *testing.T) {
	opt := newOpt()
	optFunc := OptOrder("popularity", Asc)

	optFunc(&opt)

	eOrd := "popularity:asc"
	aOrd := opt.Values.Get("order")
	if aOrd != eOrd {
		t.Errorf("Expected '%s', got '%s'", eOrd, aOrd)
	}
}
