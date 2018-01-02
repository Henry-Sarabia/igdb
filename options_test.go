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
				t.Fatal(err)
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
				t.Fatal(err)
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
				t.Fatal(err)
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
				t.Fatal(err)
			}
			funcOpt := SetFields(ft.Fields...)

			err = funcOpt(opt)
			if !reflect.DeepEqual(err, ft.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", ft.ExpErr, err)
			}

			actFields := opt.Values.Get("fields")
			if actFields != ft.ExpFields {
				t.Fatalf("Expected fields '%s', got '%s'", ft.ExpFields, actFields)
			}
		})
	}
}

func TestSetFilter(t *testing.T) {
	var filterTests = []struct {
		Name      string
		Field     string
		Op        operator
		Vals      []string
		ExpFilter string
		ExpErr    error
	}{
		{"Non-empty field and non-empty value", "rating", OpGreaterThanEqual, []string{"60"}, "60", nil},
		{"Non-empty field and non-empty values", "rating", OpGreaterThanEqual, []string{"60", "80"}, "", ErrTooManyArgs},
		{"Non-empty field and empty value", "name", OpPrefix, []string{""}, "", ErrEmptyFilterValue},
		{"Non-empty field and empty values", "name", OpPrefix, []string{"", ""}, "", ErrEmptyFilterValue},
		{"Non-empty field and no values", "rating", OpGreaterThanEqual, nil, "", ErrEmptyFilterValue},
		{"Empty field and non-empty value", "", OpEquals, []string{"Megaman X1"}, "", ErrEmptyField},
		{"Empty field and empty value", "", OpEquals, []string{""}, "", ErrEmptyField},
		{"Empty field and no values", "", OpEquals, nil, "", ErrEmptyField},
		{"OpExists, non-empty field, and non-empty value", "version_parent", OpExists, []string{"not empty"}, "", ErrTooManyArgs},
		{"OpExists, non-empty field, and non-empty values", "version_parent", OpExists, []string{"not", "empty"}, "", ErrTooManyArgs},
		{"OpExists, non-empty field, and empty value", "version_parent", OpExists, []string{""}, "", ErrTooManyArgs},
		{"OpExists, non-empty field, and no values", "version_parent", OpExists, nil, "1", nil},
		{"OpExists, empty field, and non-empty value", "", OpExists, []string{"not empty"}, "", ErrEmptyField},
		{"OpExists, empty field, and empty value", "", OpExists, []string{""}, "", ErrEmptyField},
		{"OpExists, empty field, and no values", "", OpExists, nil, "", ErrEmptyField},
		{"OpIn, non-empty field, and non-empty value", "platforms", OpIn, []string{"not empty"}, "not empty", nil},
		{"OpIn, non-empty field, and non-empty values", "platforms", OpIn, []string{"not", "empty"}, "not,empty", nil},
		{"OpIn, non-empty field, and empty value", "platforms", OpIn, []string{""}, "", ErrEmptyFilterValue},
		{"OpIn, non-empty field, and no values", "platforms", OpIn, nil, "", ErrEmptyFilterValue},
		{"OpIn, empty field, and non-empty value", "", OpIn, []string{"not empty"}, "", ErrEmptyField},
		{"OpIn, empty field, and empty value", "", OpIn, []string{""}, "", ErrEmptyField},
		{"OpIn, empty field, and no values", "", OpIn, nil, "", ErrEmptyField},
	}

	for _, ft := range filterTests {
		t.Run(ft.Name, func(t *testing.T) {
			opt, err := newOpt()
			if err != nil {
				t.Fatal(err)
			}
			funcOpt := SetFilter(ft.Field, ft.Op, ft.Vals...)

			err = funcOpt(opt)
			if !reflect.DeepEqual(err, ft.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", ft.ExpErr, err)
			}

			actFilter := opt.Values.Get(fmt.Sprintf("filter[%s][%s]", ft.Field, ft.Op))
			if actFilter != ft.ExpFilter {
				t.Fatalf("Expected filter '%s', got '%s'", ft.ExpFilter, actFilter)
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
				t.Fatal(err)
			}
			funcOpt := setSearch(st.Qry)

			err = funcOpt(opt)
			if !reflect.DeepEqual(err, st.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", st.ExpErr, err)
			}

			actQry := opt.Values.Get("search")
			if actQry != st.ExpQry {
				t.Fatalf("Expected query '%s', got '%s'", st.ExpQry, actQry)
			}
		})
	}
}

func TestSetOverlap(t *testing.T) {
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

func ExampleComposeOptions() {
	c := NewClient("YOUR_API_KEY", nil)

	// Composing FuncOptions to filter out unpopular results
	composedOpts := ComposeOptions(
		SetFields("title", "username", "game", "likes", "content"),
		SetFilter("likes", OpGreaterThanEqual, "10"),
		SetFilter("views", OpGreaterThanEqual, "200"),
		SetLimit(50),
	)

	// Using composed FuncOptions
	mario, err := c.Reviews.Search("mario", composedOpts)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Reusing composed FuncOptions
	sonic, err := c.Reviews.Search("sonic", composedOpts)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Popular reviews related to Mario")
	for _, v := range mario {
		fmt.Println(*v)
	}

	fmt.Println("Popular reviews related to Sonic")
	for _, v := range sonic {
		fmt.Println(*v)
	}
}

func ExampleSetOrder() {
	c := NewClient("YOUR_API_KEY", nil)

	// Retrieve most relevant reviews - default
	c.Reviews.Search("zelda")

	// Retrieve most viewed reviews
	c.Reviews.Search("zelda", SetOrder("views", OrderDescending))

	// Retrieve least liked reviews
	c.Reviews.Search("zelda", SetOrder("likes", OrderAscending))

	// Retrieve arliest released games by their initial release version
	c.Games.List(nil, SetOrder("release_dates.date", OrderDescending, SubMin))
}

func ExampleSetLimit() {
	c := NewClient("YOUR_API_KEY", nil)

	// Retrieve up to 10 results - default
	c.Characters.Search("snake")

	// Retrieve up to 50 results
	c.Characters.Search("snake", SetLimit(50))

	// Retrieve up to 1 result
	c.Characters.Search("snake", SetLimit(1))
}

func ExampleSetOffset() {
	c := NewClient("YOUR_API_KEY", nil)

	batchLimit := SetLimit(50)

	// Retrieve first batch of results - default
	c.People.List(nil, batchLimit)

	// Retrieve second batch of results
	c.People.List(nil, batchLimit, SetOffset(50))

	// Retrieve third batch of results
	c.People.List(nil, batchLimit, SetOffset(100))

	// Retrieve fourth batch of results
	c.People.List(nil, batchLimit, SetOffset(150))
}

func ExampleSetFields() {
	c := NewClient("YOUR_API_KEY", nil)

	// Retrieve name field
	c.Characters.Search("mario", SetFields("name"))

	// Retrieve gender field
	c.Characters.Search("mario", SetFields("gender"))

	// Retrieve both name and gender field
	c.Characters.Search("mario", SetFields("name", "gender"))

	// Retrieve whole mug_shot field
	c.Characters.Search("mario", SetFields("mug_shot"))

	// Retrieve only mug_shot.width field
	c.Characters.Search("mario", SetFields("mug_shot.width"))

	// Retrieve any number of fields
	c.Characters.Search("mario", SetFields("name", "gender", "url", "species", "games", "mug_shot.width", "mug_shot.height"))

	// Retrieve all available fields
	c.Characters.Search("mario", SetFields("*"))
}

func ExampleSetFilter() {
	c := NewClient("YOUR_API_KEY", nil)

	// Retrieve unfiltered games - default
	c.Games.List(nil)

	// Retrieve games with popularity above 50
	c.Games.List(nil, SetFilter("popularity", OpGreaterThan, "50"))

	// Retrieve games with cover art
	c.Games.List(nil, SetFilter("cover", OpExists))

	// Retrieve games released on PS4 (platform ID of 48)
	c.Games.List(nil, SetFilter("platforms", OpIn, "48"))

	// Retrieve games whose ID is not 1234
	// (This is a special case where ID can be used for filtering,
	// as it is not normally allowed except for filtering out a
	// specific entry)
	c.Games.List(nil, SetFilter("id", OpNotIn, "1234"))

	// Retrieve games whose name does not match "Horizon: Zero Dawn"
	c.Games.List(nil, SetFilter("name", OpNotEquals, "Horizon: Zero Dawn"))

	// Retrieve games whose ESRB synopsis begins with "Contains adult themes"
	c.Games.List(nil, SetFilter("esrb.synopsis", OpPrefix, "Contains adult themes"))

	// Retrieve games that meet all the previous requirements
	c.Games.List(
		nil,
		SetFilter("popularity", OpGreaterThan, "50"),
		SetFilter("cover", OpExists),
		SetFilter("platforms", OpIn, "48"),
		SetFilter("id", OpNotIn, "1234"),
		SetFilter("name", OpNotEquals, "Horizon: Zero Dawn"),
		SetFilter("esrb.synopsis", OpPrefix, "Contains adult themes"),
	)
}
