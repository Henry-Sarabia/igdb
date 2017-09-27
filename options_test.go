package igdb

import (
	"strconv"
	"testing"
)

func TestOptOrder(t *testing.T) {
	opt := newOpt()
	optFunc := OptOrder("popularity", Asc)

	optFunc(&opt)

	eOrd := "popularity:asc"
	aOrd := opt.Values.Get("order")
	if aOrd != eOrd {
		t.Errorf("Expected order '%s', got '%s'", eOrd, aOrd)
	}
}

func TestOptLimit(t *testing.T) {
	opt := newOpt()
	optFunc := OptLimit(20)

	optFunc(&opt)

	eLim := 20
	aLim, err := strconv.Atoi(opt.Values.Get("limit"))
	if err != nil {
		t.Error(err)
	}
	if aLim != eLim {
		t.Errorf("Expected limit %d, got %d", eLim, aLim)
	}
}

func TestOptOffset(t *testing.T) {
	opt := newOpt()
	optFunc := OptOffset(5)

	optFunc(&opt)

	eOff := 5
	aOff, err := strconv.Atoi(opt.Values.Get("offset"))
	if err != nil {
		t.Error(err)
	}
	if aOff != eOff {
		t.Errorf("Expected offset %d, got %d", eOff, aOff)
	}
}
