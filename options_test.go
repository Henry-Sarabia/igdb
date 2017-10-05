package igdb

import (
	"strconv"
	"testing"
)

func TestOptOrder(t *testing.T) {
	opt, err := newOpt()
	if err != nil {
		t.Error(err)
	}
	optFunc := OptOrder("popularity", Ascend)

	optFunc(opt)

	eOrd := "popularity:asc"
	aOrd := opt.Values.Get("order")
	if aOrd != eOrd {
		t.Errorf("Expected order '%s', got '%s'", eOrd, aOrd)
	}
}

func TestOptLimit(t *testing.T) {
	opt, err := newOpt()
	if err != nil {
		t.Error(err)
	}
	optFunc := OptLimit(20)
	// var limitTests = []struct {
	// 	Name string
	// 	Limit int
	// } {
	// 	{"Happy path", 20},
	// 	{"invalid negative limit", }
	// }

	optFunc(opt)

	eLim := strconv.Itoa(20)
	aLim := opt.Values.Get("limit")

	if aLim != eLim {
		t.Errorf("Expected limit %s, got %s", eLim, aLim)
	}
}

func TestOptOffset(t *testing.T) {
	opt, err := newOpt()
	if err != nil {
		t.Error(err)
	}
	optFunc := OptOffset(5)

	optFunc(opt)

	eOff := strconv.Itoa(5)
	aOff := opt.Values.Get("offset")
	if aOff != eOff {
		t.Errorf("Expected offset %s, got %s", eOff, aOff)
	}
}

func TestOptFields(t *testing.T) {
	opt, err := newOpt()
	if err != nil {
		t.Error(err)
	}
	optFunc := OptFields("name", "rating", "popularity")

	optFunc(opt)

	eFld := "name,rating,popularity"
	aFld := opt.Values.Get("fields")
	if aFld != eFld {
		t.Errorf("Expected fields '%s', got '%s'", eFld, aFld)
	}
}

func TestOptFieldsEmpty(t *testing.T) {
	opt, err := newOpt()
	if err != nil {
		t.Error(err)
	}
	optFunc := OptFields()

	optFunc(opt)

	eFld := ""
	aFld := opt.Values.Get("fields")
	if aFld != eFld {
		t.Errorf("Expected empty fields, got '%s'", aFld)
	}
}

func TestOptFilter(t *testing.T) {
	opt, err := newOpt()
	if err != nil {
		t.Error(err)
	}
	optFunc := OptFilter("popularity", LessThanEqual, "50")

	optFunc(opt)

	eFil := "50"
	aFil := opt.Values.Get("filter[popularity][lte]")
	if aFil != eFil {
		t.Errorf("Expected filter '%s', got '%s'", eFil, aFil)
	}
}

func TestOptSearch(t *testing.T) {
	opt, err := newOpt()
	if err != nil {
		t.Error(err)
	}
	optFunc := optSearch("mario party")

	optFunc(opt)

	eQry := "mario party"
	aQry := opt.Values.Get("search")
	if aQry != eQry {
		t.Errorf("Expected query '%s', got '%s'", eQry, aQry)
	}
}

func TestOptScroll(t *testing.T) {
	opt, err := newOpt()
	if err != nil {
		t.Error(err)
	}
	optFunc := OptScroll(3)

	optFunc(opt)

	ePage := strconv.Itoa(3)
	aPage := opt.Values.Get("scroll")
	if aPage != ePage {
		t.Errorf("Expected page %s, got %s", ePage, aPage)
	}
}

func TestNewOptEmpty(t *testing.T) {
	opt, err := newOpt()
	if err != nil {
		t.Error(err)
	}

	evl := 0
	avl := len(opt.Values)
	if avl != evl {
		t.Errorf("Expected Values map length %d, got %d", evl, avl)
	}
}

func TestNewOptSingle(t *testing.T) {
	opt, err := newOpt(OptLimit(20))
	if err != nil {
		t.Error(err)
	}

	evl := 1
	avl := len(opt.Values)
	if avl != evl {
		t.Errorf("Expected Values map length %d, got %d", evl, avl)
	}
}

func TestNewOptMulti(t *testing.T) {
	opt, err := newOpt(OptFields("name", "rating"),
		OptFilter("name", Equals, "zelda"),
		OptLimit(5),
		OptOffset(10),
		OptOrder("rating", Descend))
	if err != nil {
		t.Error(err)
	}

	evl := 5
	avl := len(opt.Values)
	if avl != evl {
		t.Errorf("Expected Values map length %d, got %d", evl, avl)
	}
}
