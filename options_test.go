package igdb

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewOpt(t *testing.T) {
	var optTests = []struct {
		Name     string
		OptFuncs []optionFunc
		ExpCount int
		ExpErr   error
	}{
		{"Empty option", []optionFunc{}, 0, nil},
		{"Single option", []optionFunc{OptLimit(4)}, 1, nil},
		{"Multiple options", []optionFunc{OptOffset(10), OptLimit(50), OptFields("id", "rating"), OptFilter("rating", OpLessThan, "40"), OptOrder("rating", AscendingOrder)}, 5, nil},
		{"Multiple filter options", []optionFunc{OptFilter("popularity", OpLessThan, "50"), OptFilter("rating", OpGreaterThan, "50")}, 2, nil},
	}

	for _, ot := range optTests {
		t.Run(ot.Name, func(t *testing.T) {
			opt, err := newOpt(ot.OptFuncs...)
			if !reflect.DeepEqual(err, ot.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", ot.ExpErr, err)
			}

			actCount := len(opt.Values)
			if actCount != ot.ExpCount {
				t.Fatalf("Expected count of %d, got %d", ot.ExpCount, actCount)
			}
		})
	}
}

func TestOptOrder(t *testing.T) {
	var orderTests = []struct {
		Name   string
		Field  string
		Order  order
		ExpOrd string
		ExpErr error
	}{
		{"Non-empty field", "popularity", AscendingOrder, "popularity:asc", nil},
		{"Empty field", "", DescendingOrder, "", ErrEmptyField},
	}

	for _, ot := range orderTests {
		t.Run(ot.Name, func(t *testing.T) {
			opt, err := newOpt()
			if err != nil {
				t.Fatalf(err.Error())
			}
			optFunc := OptOrder(ot.Field, ot.Order)

			err = optFunc(opt)
			if !reflect.DeepEqual(err, ot.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", ot.ExpErr, err)
			}

			actOrd := opt.Values.Get("order")
			if actOrd != ot.ExpOrd {
				t.Fatalf("Expected order '%s', got '%s'", ot.ExpOrd, actOrd)
			}
		})
	}
}

func TestOptLimit(t *testing.T) {
	var limitTests = []struct {
		Name   string
		Limit  int
		ExpLim string
		ExpErr error
	}{
		{"Limit within range", 15, "15", nil},
		{"Zero limit", 0, "", ErrOutOfRange},
		{"Limit below range", -10, "", ErrOutOfRange},
		{"Limit above range", 51, "", ErrOutOfRange},
	}

	for _, lt := range limitTests {
		t.Run(lt.Name, func(t *testing.T) {
			opt, err := newOpt()
			if err != nil {
				t.Fatalf(err.Error())
			}
			optFunc := OptLimit(lt.Limit)

			err = optFunc(opt)
			if !reflect.DeepEqual(err, lt.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", lt.ExpErr, err)
			}

			actLim := opt.Values.Get("limit")
			if actLim != lt.ExpLim {
				t.Fatalf("Expected limit '%s', got '%s'", lt.ExpLim, actLim)
			}
		})
	}
}

func TestOptOffset(t *testing.T) {
	var offsetTests = []struct {
		Name   string
		Offset int
		ExpOff string
		ExpErr error
	}{
		{"Offset within range", 20, "20", nil},
		{"Zero offset", 0, "0", nil},
		{"Offset below range", -15, "", ErrOutOfRange},
		{"Offset above range", 100, "", ErrOutOfRange},
	}

	for _, ot := range offsetTests {
		t.Run(ot.Name, func(t *testing.T) {
			opt, err := newOpt()
			if err != nil {
				t.Fatalf(err.Error())
			}
			optFunc := OptOffset(ot.Offset)

			err = optFunc(opt)
			if !reflect.DeepEqual(err, ot.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", ot.ExpErr, err)
			}

			actOff := opt.Values.Get("offset")
			if actOff != ot.ExpOff {
				t.Fatalf("Expected offset '%s', got '%s'", ot.ExpOff, actOff)
			}
		})
	}
}

func TestOptFields(t *testing.T) {
	var fieldsTests = []struct {
		Name      string
		Fields    []string
		ExpFields string
		ExpErr    error
	}{
		{"Single non-empty field", []string{"name"}, "name", nil},
		{"Multiple non-empty fields", []string{"name", "popularity", "rating"}, "name,popularity,rating", nil},
		{"Empty fields slice", []string{}, "", ErrEmptySlice},
		{"Single empty field", []string{""}, "", ErrEmptyField},
		{"Multiple empty fields", []string{"", "", "", ""}, "", ErrEmptyField},
		{"Mixed empty and non-empty fields", []string{"", "id", "", "url"}, "", ErrEmptyField},
	}

	for _, ft := range fieldsTests {
		t.Run(ft.Name, func(t *testing.T) {
			opt, err := newOpt()
			if err != nil {
				t.Fatalf(err.Error())
			}
			optFunc := OptFields(ft.Fields...)

			err = optFunc(opt)
			if !reflect.DeepEqual(err, ft.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", ft.ExpErr, err)
			}

			actFields := opt.Values.Get("fields")
			if actFields != ft.ExpFields {
				t.Fatalf("Expected order '%s', got '%s'", ft.ExpFields, actFields)
			}
		})
	}
}

func TestOptFilter(t *testing.T) {
	var filterTests = []struct {
		Name      string
		Field     string
		Op        operator
		Val       string
		ExpFilter string
		ExpErr    error
	}{
		{"Non-empty field and non-empty value", "rating", OpGreaterThanEqual, "60", "", nil},
		{"Non-empty field and empty value", "name", OpPrefix, "", "", ErrEmptyField},
		{"Empty field and non-empty value", "", OpEquals, "Megaman X1", "", ErrEmptyField},
		{"Empty field and empty value", "", OpIn, "", "", ErrEmptyField},
	}

	for _, ft := range filterTests {
		t.Run(ft.Name, func(t *testing.T) {
			opt, err := newOpt()
			if err != nil {
				t.Fatalf(err.Error())
			}
			optFunc := OptFilter(ft.Field, ft.Op, ft.Val)

			err = optFunc(opt)
			if !reflect.DeepEqual(err, ft.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", ft.ExpErr, err)
			}

			actFilter := opt.Values.Get(fmt.Sprintf("[%s][%s]", ft.Field, ft.Op))
			if actFilter != ft.ExpFilter {
				t.Fatalf("Expected order '%s', got '%s'", ft.ExpFilter, actFilter)
			}
		})
	}
}

func TestOptScroll(t *testing.T) {
	var scrollTests = []struct {
		Name    string
		Page    int
		ExpPage string
		ExpErr  error
	}{
		{"Page within range", 5, "5", nil},
		{"Zero page", 0, "", ErrOutOfRange},
		{"Page below range", -15, "", ErrOutOfRange},
	}

	for _, st := range scrollTests {
		t.Run(st.Name, func(t *testing.T) {
			opt, err := newOpt()
			if err != nil {
				t.Fatalf(err.Error())
			}
			optFunc := OptScroll(st.Page)

			err = optFunc(opt)
			if !reflect.DeepEqual(err, st.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", st.ExpErr, err)
			}

			actPage := opt.Values.Get("scroll")
			if actPage != st.ExpPage {
				t.Fatalf("Expected offset '%s', got '%s'", st.ExpPage, actPage)
			}
		})
	}
}

func TestOptSearch(t *testing.T) {
	var searchTests = []struct {
		Name   string
		Qry    string
		ExpQry string
		ExpErr error
	}{
		{"Non-empty query", "zelda", "zelda", nil},
		{"Non-Empty query with spaces", "the legend of zelda", "the legend of zelda", nil},
		{"Empty query", "", "", ErrEmptyQuery},
	}

	for _, st := range searchTests {
		t.Run(st.Name, func(t *testing.T) {
			opt, err := newOpt()
			if err != nil {
				t.Fatalf(err.Error())
			}
			optFunc := optSearch(st.Qry)

			err = optFunc(opt)
			if !reflect.DeepEqual(err, st.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", st.ExpErr, err)
			}

			actQry := opt.Values.Get("search")
			if actQry != st.ExpQry {
				t.Fatalf("Expected offset '%s', got '%s'", st.ExpQry, actQry)
			}
		})
	}
}

func TestOptOverlap(t *testing.T) {
	var overlapTests = []struct {
		Name     string
		OptFuncs []optionFunc
		ExpErr   error
	}{
		{"OptOrder overlap", []optionFunc{OptOrder("popularity", DescendingOrder), OptOrder("rating", AscendingOrder)}, ErrOptionSet},
		{"OptLimit overlap", []optionFunc{OptLimit(5), OptLimit(40)}, ErrOptionSet},
		{"OptOffset overlap", []optionFunc{OptOffset(0), OptOffset(25)}, ErrOptionSet},
		{"OptFields overlap", []optionFunc{OptFields("id"), OptFields("name")}, ErrOptionSet},
		{"OptFilter overlap", []optionFunc{OptFilter("rating", OpLessThan, "50"), OptFilter("popularity", OpGreaterThan, "50")}, nil},
		{"OptScroll overlap", []optionFunc{OptScroll(1), OptScroll(10)}, ErrOptionSet},
		{"OptSearch overlap", []optionFunc{optSearch("zelda"), optSearch("link")}, ErrOptionSet},
		{"OptOffset and OptScroll overlap", []optionFunc{OptOffset(40), OptScroll(4)}, ErrExclusiveOption},
		{"OptScroll and OptOffset overlap", []optionFunc{OptScroll(3), OptOffset(30)}, ErrExclusiveOption},
	}

	for _, ot := range overlapTests {
		t.Run(ot.Name, func(t *testing.T) {
			_, err := newOpt(ot.OptFuncs...)
			if !reflect.DeepEqual(err, ot.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", ot.ExpErr, err)
			}
		})
	}
}
