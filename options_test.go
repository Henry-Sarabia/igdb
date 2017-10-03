package igdb

import (
	"strconv"
	"testing"
)

func TestNewOptEmpty(t *testing.T) {
	opt := newOpt()

	evl := 0
	avl := len(opt.Values)
	if avl != evl {
		t.Errorf("Expected Values map length %d, got %d", evl, avl)
	}
}

func TestNewOptSingle(t *testing.T) {
	opt := newOpt(OptLimit(20))

	evl := 1
	avl := len(opt.Values)
	if avl != evl {
		t.Errorf("Expected Values map length %d, got %d", evl, avl)
	}
}

func TestNewOptMulti(t *testing.T) {
	opt := newOpt(OptFields("name", "rating"),
		OptFilter("name", Equals, "zelda"),
		OptLimit(5),
		OptOffset(10),
		OptOrder("rating", Descend))

	evl := 5
	avl := len(opt.Values)
	if avl != evl {
		t.Errorf("Expected Values map length %d, got %d", evl, avl)
	}
}

func TestNewOptOverlap(t *testing.T) {
	opt := newOpt(OptFields("name", "rating"),
		OptFilter("name", Equals, "zelda"),
		OptLimit(5),
		OptOffset(10),
		OptOrder("rating", Descend),
		OptFields("id", "popularity"),
		OptFilter("id", NotIn, "1234"),
		OptLimit(25),
		OptOffset(50),
		OptOrder("popularity", Ascend))

	evl := 6
	avl := len(opt.Values)
	if avl != evl {
		t.Errorf("Expecting Values map length %d, got %d", evl, avl)
	}
}

func TestOptOrder(t *testing.T) {
	opt := newOpt()
	optFunc := OptOrder("popularity", Ascend)

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

	eLim := strconv.Itoa(20)
	aLim := opt.Values.Get("limit")

	if aLim != eLim {
		t.Errorf("Expected limit %s, got %s", eLim, aLim)
	}
}

func TestOptOffset(t *testing.T) {
	opt := newOpt()
	optFunc := OptOffset(5)

	optFunc(&opt)

	eOff := strconv.Itoa(5)
	aOff := opt.Values.Get("offset")
	if aOff != eOff {
		t.Errorf("Expected offset %s, got %s", eOff, aOff)
	}
}

func TestOptFields(t *testing.T) {
	opt := newOpt()
	optFunc := OptFields("name", "rating", "popularity")

	optFunc(&opt)

	eFld := "name,rating,popularity"
	aFld := opt.Values.Get("fields")
	if aFld != eFld {
		t.Errorf("Expected fields '%s', got '%s'", eFld, aFld)
	}
}

func TestOptFieldsEmpty(t *testing.T) {
	opt := newOpt()
	optFunc := OptFields()

	optFunc(&opt)

	eFld := ""
	aFld := opt.Values.Get("fields")
	if aFld != eFld {
		t.Errorf("Expected empty fields, got '%s'", aFld)
	}
}

func TestOptFilter(t *testing.T) {
	opt := newOpt()
	optFunc := OptFilter("popularity", LessThanEqual, "50")

	optFunc(&opt)

	eFil := "50"
	aFil := opt.Values.Get("filter[popularity][lte]")
	if aFil != eFil {
		t.Errorf("Expected filter '%s', got '%s'", eFil, aFil)
	}
}

func TestOptSearch(t *testing.T) {
	opt := newOpt()
	optFunc := optSearch("mario party")

	optFunc(&opt)

	eQry := "mario party"
	aQry := opt.Values.Get("search")
	if aQry != eQry {
		t.Errorf("Expected query '%s', got '%s'", eQry, aQry)
	}
}

func TestOptScroll(t *testing.T) {
	opt := newOpt()
	optFunc := OptScroll(3)

	optFunc(&opt)

	ePage := strconv.Itoa(3)
	aPage := opt.Values.Get("scroll")
	if aPage != ePage {
		t.Errorf("Expected page %s, got %s", ePage, aPage)
	}
}
