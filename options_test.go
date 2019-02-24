package igdb

import (
	"github.com/Henry-Sarabia/apicalypse"
	"github.com/pkg/errors"
	"strings"
	"testing"
)

func TestComposeOptions(t *testing.T) {
	var optTests = []struct {
		name        string
		opts        []Option
		wantFilters []string
		wantErr     error
	}{
		{"Zero options", nil, nil, nil},
		{"Single option", []Option{SetLimit(20)}, []string{"20"}, nil},
		{"Multiple options", []Option{SetLimit(20), SetFields("name", "id"), SetFilter("popularity", OpLessThan, "50")}, []string{"50", "name,id", "where popularity < 50"}, nil},
		{"Single invalid option", []Option{SetOffset(-500)}, nil, ErrOutOfRange},
		{"Multiple invalid options", []Option{SetOffset(-500), SetLimit(-999)}, nil, ErrOutOfRange},
	}

	for _, test := range optTests {
		t.Run(test.name, func(t *testing.T) {
			fn, err := ComposeOptions(test.opts...)()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if test.wantErr != nil {
				return
			}

			q, err := apicalypse.Query(fn)
			if err != nil {
				t.Fatal(err)
			}

			for _, f := range test.wantFilters {
				if !strings.Contains(q, f) {
					t.Errorf("got: <%v>, want: <%v>", q, f)
				}

			}
		})
	}
}

func TestSetOrder(t *testing.T) {
	var tests = []struct {
		name    string
		field   string
		order   order
		wantOrd string
		wantErr error
	}{
		{"Non-empty field with ascending order", "rating", OrderAscending, "rating asc", nil},
		{"Non-empty field with descending order", "rating", OrderDescending, "rating desc", nil},
		{"Empty field with ascending order", "  ", OrderAscending, "", ErrEmptyFields},
		{"Empty field with descending order", "  ", OrderDescending, "", ErrEmptyFields},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fn, err := SetOrder(test.field, test.order)()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if test.wantErr != nil {
				return
			}

			q, err := apicalypse.Query(fn)
			if err != nil {
				t.Fatal(err)
			}

			if !strings.Contains(q, test.wantOrd) {
				t.Errorf("got: <%v>, want: <%v>", q, test.wantOrd)
			}
		})
	}
}

func TestSetLimit(t *testing.T) {
	var tests = []struct {
		name      string
		limit     int
		wantLimit string
		wantErr   error
	}{
		{"Limit within range", 5, "5", nil},
		{"Zero limit", 0, "", ErrOutOfRange},
		{"Limit below range", -10, "", ErrOutOfRange},
		{"Limit above range", 5001, "", ErrOutOfRange},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fn, err := SetLimit(test.limit)()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if test.wantErr != nil {
				return
			}

			q, err := apicalypse.Query(fn)
			if err != nil {
				t.Fatal(err)
			}

			if !strings.Contains(q, test.wantLimit) {
				t.Errorf("got: <%v>, want: <%v>", q, test.wantLimit)
			}
		})
	}
}

func TestSetOffset(t *testing.T) {
	var tests = []struct {
		name       string
		offset     int
		wantOffset string
		wantErr    error
	}{
		{"Offset within range", 20, "20", nil},
		{"Zero offset", 0, "0", nil},
		{"Offset below range", -15, "", ErrOutOfRange},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fn, err := SetOffset(test.offset)()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if test.wantErr != nil {
				return
			}

			q, err := apicalypse.Query(fn)
			if err != nil {
				t.Fatal(err)
			}

			if !strings.Contains(q, test.wantOffset) {
				t.Errorf("got: <%v>, want: <%v>", q, test.wantOffset)
			}
		})
	}
}

func TestSetFields(t *testing.T) {
	var tests = []struct {
		name       string
		fields     []string
		wantFields string
		wantErr    error
	}{
		{"Single non-empty field", []string{"name"}, "name", nil},
		{"Multiple non-empty fields", []string{"name", "popularity", "rating"}, "name,popularity,rating", nil},
		{"Empty fields slice", []string{}, "", ErrEmptyFields},
		{"Single empty field", []string{"  "}, "", ErrEmptyFields},
		{"Multiple empty fields", []string{"", " ", "", ""}, "", ErrEmptyFields},
		{"Mixed empty and non-empty fields", []string{"", "id", "  ", "url"}, "", ErrEmptyFields},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fn, err := SetFields(test.fields...)()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if test.wantErr != nil {
				return
			}

			q, err := apicalypse.Query(fn)
			if err != nil {
				t.Fatal(err)
			}

			if !strings.Contains(q, test.wantFields) {
				t.Errorf("got: <%v>, want: <%v>", q, test.wantFields)
			}
		})
	}
}

func TestSetExclude(t *testing.T) {
	var tests = []struct {
		name       string
		fields     []string
		wantFields string
		wantErr    error
	}{
		{"Single non-empty field", []string{"name"}, "name", nil},
		{"Multiple non-empty fields", []string{"name", "popularity", "rating"}, "name,popularity,rating", nil},
		{"Empty fields slice", []string{}, "", ErrEmptyFields},
		{"Single empty field", []string{"  "}, "", ErrEmptyFields},
		{"Multiple empty fields", []string{"", " ", "", ""}, "", ErrEmptyFields},
		{"Mixed empty and non-empty fields", []string{"", "id", "  ", "url"}, "", ErrEmptyFields},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fn, err := SetExclude(test.fields...)()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if test.wantErr != nil {
				return
			}

			q, err := apicalypse.Query(fn)
			if err != nil {
				t.Fatal(err)
			}

			if !strings.Contains(q, test.wantFields) {
				t.Errorf("got: <%v>, want: <%v>", q, test.wantFields)
			}
		})
	}
}

func TestSetFilter(t *testing.T) {
	var tests = []struct {
		name       string
		field      string
		op         operator
		vals       []string
		wantFilter string
		wantErr    error
	}{
		{"Non-empty field and non-empty value", "rating", OpGreaterThanEqual, []string{"60"}, "60", nil},
		{"Non-empty field and non-empty values", "games", OpContainsAll, []string{"123", "456"}, "123,456", nil},
		{"Non-empty field and empty value", "name", OpNotEquals, []string{""}, "", ErrEmptyFilterVals},
		{"Non-empty field and empty values", "name", OpNotEquals, []string{"", ""}, "", ErrEmptyFilterVals},
		{"Non-empty field and no values", "rating", OpGreaterThanEqual, nil, "", ErrEmptyFilterVals},
		{"Empty field and non-empty value", "", OpEquals, []string{"Megaman X1"}, "", ErrEmptyFields},
		{"Empty field and empty value", "", OpEquals, []string{""}, "", ErrEmptyFields},
		{"Empty field and no values", "", OpEquals, nil, "", ErrEmptyFields},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fn, err := SetFilter(test.field, test.op, test.vals...)()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if test.wantErr != nil {
				return
			}

			q, err := apicalypse.Query(fn)
			if err != nil {
				t.Fatal(err)
			}

			if !strings.Contains(q, test.wantFilter) {
				t.Errorf("got: <%v>, want: <%v>", q, test.wantFilter)
			}
		})
	}
}

func TestSetSearch(t *testing.T) {
	var tests = []struct {
		name    string
		qry     string
		wantQry string
		wantErr error
	}{
		{"Non-empty query", "zelda", "zelda", nil},
		{"Non-Empty query with spaces", "the legend of zelda", "the legend of zelda", nil},
		{"Empty query", "", "", ErrEmptyQry},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fn, err := setSearch(test.qry)()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if test.wantErr != nil {
				return
			}

			q, err := apicalypse.Query(fn)
			if err != nil {
				t.Fatal(err)
			}

			if !strings.Contains(q, test.wantQry) {
				t.Errorf("got: <%v>, want: <%v>", q, test.wantQry)
			}
		})
	}
}

//func ExampleComposeOptions() {
//	c := NewClient("YOUR_API_KEY", nil)
//
//	// Composing FuncOptions to filter out unpopular results
//	composedOpts := ComposeOptions(
//		SetFields("title", "username", "game", "likes", "content"),
//		SetFilter("likes", OpGreaterThanEqual, "10"),
//		SetFilter("views", OpGreaterThanEqual, "200"),
//		SetLimit(50),
//	)
//
//	// Using composed FuncOptions
//	mario, err := c.Reviews.Search("mario", composedOpts)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// Reusing composed FuncOptions
//	sonic, err := c.Reviews.Search("sonic", composedOpts)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println("Popular reviews related to Mario")
//	for _, v := range mario {
//		fmt.Println(*v)
//	}
//
//	fmt.Println("Popular reviews related to Sonic")
//	for _, v := range sonic {
//		fmt.Println(*v)
//	}
//}
//
//func ExampleSetOrder() {
//	c := NewClient("YOUR_API_KEY", nil)
//
//	// Retrieve most relevant reviews - default
//	c.Reviews.Search("zelda")
//
//	// Retrieve most viewed reviews
//	c.Reviews.Search("zelda", SetOrder("views", OrderDescending))
//
//	// Retrieve least liked reviews
//	c.Reviews.Search("zelda", SetOrder("likes", OrderAscending))
//
//	// Retrieve arliest released games by their initial release version
//	c.Games.List(nil, SetOrder("release_dates.date", OrderDescending, SubMin))
//}
//
//func ExampleSetLimit() {
//	c := NewClient("YOUR_API_KEY", nil)
//
//	// Retrieve up to 10 results - default
//	c.Characters.Search("snake")
//
//	// Retrieve up to 50 results
//	c.Characters.Search("snake", SetLimit(50))
//
//	// Retrieve up to 1 result
//	c.Characters.Search("snake", SetLimit(1))
//}
//
//func ExampleSetOffset() {
//	c := NewClient("YOUR_API_KEY", nil)
//
//	batchLimit := SetLimit(50)
//
//	// Retrieve first batch of results - default
//	c.People.List(nil, batchLimit)
//
//	// Retrieve second batch of results
//	c.People.List(nil, batchLimit, SetOffset(50))
//
//	// Retrieve third batch of results
//	c.People.List(nil, batchLimit, SetOffset(100))
//
//	// Retrieve fourth batch of results
//	c.People.List(nil, batchLimit, SetOffset(150))
//}
//
//func ExampleSetFields() {
//	c := NewClient("YOUR_API_KEY", nil)
//
//	// Retrieve name field
//	c.Characters.Search("mario", SetFields("name"))
//
//	// Retrieve gender field
//	c.Characters.Search("mario", SetFields("gender"))
//
//	// Retrieve both name and gender field
//	c.Characters.Search("mario", SetFields("name", "gender"))
//
//	// Retrieve whole mug_shot field
//	c.Characters.Search("mario", SetFields("mug_shot"))
//
//	// Retrieve only mug_shot.width field
//	c.Characters.Search("mario", SetFields("mug_shot.width"))
//
//	// Retrieve any number of fields
//	c.Characters.Search("mario", SetFields("name", "gender", "url", "species", "games", "mug_shot.width", "mug_shot.height"))
//
//	// Retrieve all available fields
//	c.Characters.Search("mario", SetFields("*"))
//}
//
//func ExampleSetFilter() {
//	c := NewClient("YOUR_API_KEY", nil)
//
//	// Retrieve unfiltered games - default
//	c.Games.List(nil)
//
//	// Retrieve games with popularity above 50
//	c.Games.List(nil, SetFilter("popularity", OpGreaterThan, "50"))
//
//	// Retrieve games with cover art
//	c.Games.List(nil, SetFilter("cover", OpExists))
//
//	// Retrieve games released on PS4 (platform ID of 48)
//	c.Games.List(nil, SetFilter("platforms", OpIn, "48"))
//
//	// Retrieve games whose ID is not 1234
//	// (This is a special case where ID can be used for filtering,
//	// as it is not normally allowed except for filtering out a
//	// specific entry)
//	c.Games.List(nil, SetFilter("id", OpNotIn, "1234"))
//
//	// Retrieve games whose name does not match "Horizon: Zero Dawn"
//	c.Games.List(nil, SetFilter("name", OpNotEquals, "Horizon: Zero Dawn"))
//
//	// Retrieve games whose ESRB synopsis begins with "Contains adult themes"
//	c.Games.List(nil, SetFilter("esrb.synopsis", OpPrefix, "Contains adult themes"))
//
//	// Retrieve games that meet all the previous requirements
//	c.Games.List(
//		nil,
//		SetFilter("popularity", OpGreaterThan, "50"),
//		SetFilter("cover", OpExists),
//		SetFilter("platforms", OpIn, "48"),
//		SetFilter("id", OpNotIn, "1234"),
//		SetFilter("name", OpNotEquals, "Horizon: Zero Dawn"),
//		SetFilter("esrb.synopsis", OpPrefix, "Contains adult themes"),
//	)
//}
