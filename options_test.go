package igdb

import (
	"fmt"
	"reflect"
	"testing"
)

func TestComposeOptions(t *testing.T) {
	var optTests = []struct {
		Name     string
		FuncOpts []FuncOption
	}{
		{"Zero options", nil},
		{"Single option", []FuncOption{SetLimit(20)}},
		{"Multiple options", []FuncOption{SetLimit(20), SetFields("name", "id"), SetFilter("popularity", OpLessThan, "50")}},
		{"Single invalid option", []FuncOption{SetOffset(-500)}},
		{"Multiple invalid options", []FuncOption{SetOffset(-500), SetLimit(999)}},
	}

	for _, tt := range optTests {
		t.Run(tt.Name, func(t *testing.T) {
			comp := ComposeOptions(tt.FuncOpts...)

			expOpt, expErr := newOpt(tt.FuncOpts...)
			actOpt, actErr := newOpt(comp)
			if !reflect.DeepEqual(actErr, expErr) {
				t.Fatalf("Expected error '%v', got '%v'", expErr, actErr)
			}
			if !reflect.DeepEqual(actOpt, expOpt) {
				t.Fatalf("Expected options '%v', got '%v'", expOpt, actOpt)
			}
		})
	}
}

func TestNewOpt(t *testing.T) {
	var optTests = []struct {
		Name     string
		FuncOpts []FuncOption
		ExpCount int
		ExpErr   error
	}{
		{"Empty option", []FuncOption{}, 0, nil},
		{"Single option", []FuncOption{SetLimit(4)}, 1, nil},
		{"Multiple options", []FuncOption{SetOffset(10), SetLimit(50), SetFields("id", "rating"), SetFilter("rating", OpLessThan, "40"), SetOrder("rating", OrderAscending)}, 5, nil},
		{"Multiple filter options", []FuncOption{SetFilter("popularity", OpLessThan, "50"), SetFilter("rating", OpGreaterThan, "50")}, 2, nil},
	}

	for _, ot := range optTests {
		t.Run(ot.Name, func(t *testing.T) {
			opt, err := newOpt(ot.FuncOpts...)
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

func TestSetOrder(t *testing.T) {
	var orderTests = []struct {
		Name   string
		Field  string
		Order  order
		Sub    []subfilter
		ExpOrd string
		ExpErr error
	}{
		{"Non-empty field with single subfilter", "release_dates.date", OrderDescending, []subfilter{SubMin}, "release_dates.date:desc:min", nil},
		{"Non-empty field with no subfilter", "rating", OrderAscending, nil, "rating:asc", nil},
		{"Non-empty field with multiple subfilters", "release_dates.date", OrderDescending, []subfilter{SubMin, SubMax}, "", ErrTooManyArgs},
		{"Empty field with single subfilter", "", OrderAscending, []subfilter{SubAverage}, "", ErrEmptyField},
		{"Empty field with no subfilter", "  ", OrderDescending, nil, "", ErrEmptyField},
		{"Empty field with multiple subfilters", "    ", OrderAscending, []subfilter{SubMedian, SubSum}, "", ErrEmptyField},
	}

	for _, ot := range orderTests {
		t.Run(ot.Name, func(t *testing.T) {
			opt, err := newOpt()
			if err != nil {
				t.Fatalf(err.Error())
			}
			funcOpt := SetOrder(ot.Field, ot.Order, ot.Sub...)

			err = funcOpt(opt)
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

func TestSetLimit(t *testing.T) {
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
			funcOpt := SetLimit(lt.Limit)

			err = funcOpt(opt)
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

func TestSetOffset(t *testing.T) {
	var offsetTests = []struct {
		Name   string
		Offset int
		ExpOff string
		ExpErr error
	}{
		{"Offset within range", 20, "20", nil},
		{"Zero offset", 0, "0", nil},
		{"Offset below range", -15, "", ErrOutOfRange},
		{"Offset above range", 10001, "", ErrOutOfRange},
	}

	for _, ot := range offsetTests {
		t.Run(ot.Name, func(t *testing.T) {
			opt, err := newOpt()
			if err != nil {
				t.Fatalf(err.Error())
			}
			funcOpt := SetOffset(ot.Offset)

			err = funcOpt(opt)
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

func TestSetFields(t *testing.T) {
	var fieldsTests = []struct {
		Name      string
		Fields    []string
		ExpFields string
		ExpErr    error
	}{
		{"Single non-empty field", []string{"name"}, "name", nil},
		{"Multiple non-empty fields", []string{"name", "popularity", "rating"}, "name,popularity,rating", nil},
		{"Empty fields slice", []string{}, "", ErrEmptySlice},
		{"Single empty field", []string{"  "}, "", ErrEmptyField},
		{"Multiple empty fields", []string{"", " ", "", ""}, "", ErrEmptyField},
		{"Mixed empty and non-empty fields", []string{"", "id", "  ", "url"}, "", ErrEmptyField},
	}

	for _, ft := range fieldsTests {
		t.Run(ft.Name, func(t *testing.T) {
			opt, err := newOpt()
			if err != nil {
				t.Fatalf(err.Error())
			}
			funcOpt := SetFields(ft.Fields...)

			err = funcOpt(opt)
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

func TestSetFilter(t *testing.T) {
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
			funcOpt := SetFilter(ft.Field, ft.Op, ft.Val)

			err = funcOpt(opt)
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

func TestSetSearch(t *testing.T) {
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
			funcOpt := setSearch(st.Qry)

			err = funcOpt(opt)
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
		FuncOpts []FuncOption
		ExpErr   error
	}{
		{"SetOrder overlap", []FuncOption{SetOrder("popularity", OrderDescending), SetOrder("rating", OrderAscending)}, ErrOptionSet},
		{"SetLimit overlap", []FuncOption{SetLimit(5), SetLimit(40)}, ErrOptionSet},
		{"SetOffset overlap", []FuncOption{SetOffset(0), SetOffset(25)}, ErrOptionSet},
		{"SetFields overlap", []FuncOption{SetFields("id"), SetFields("name")}, ErrOptionSet},
		{"SetFilter overlap", []FuncOption{SetFilter("rating", OpLessThan, "50"), SetFilter("popularity", OpGreaterThan, "50")}, nil},
		{"SetSearch overlap", []FuncOption{setSearch("zelda"), setSearch("link")}, ErrOptionSet},
	}

	for _, ot := range overlapTests {
		t.Run(ot.Name, func(t *testing.T) {
			_, err := newOpt(ot.FuncOpts...)
			if !reflect.DeepEqual(err, ot.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", ot.ExpErr, err)
			}
		})
	}
}
