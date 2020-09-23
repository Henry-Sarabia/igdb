package igdb

import (
	"fmt"
	"github.com/Henry-Sarabia/apicalypse"
	"github.com/pkg/errors"
	"log"
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
		{"Multiple options", []Option{SetLimit(20), SetFields("name", "id"), SetFilter("hypes", OpLessThan, "50")}, []string{"50", "name,id", "where hypes < 50"}, nil},
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

func TestUnwrapOptions(t *testing.T) {
	tests := []struct {
		name     string
		opts     []Option
		wantOpts []string
		wantErr  error
	}{
		{
			"Zero valid options",
			nil,
			nil,
			nil,
		},
		{
			"Single valid option",
			[]Option{SetLimit(10)},
			[]string{"limit 10;"},
			nil,
		},
		{
			"Multiple valid options",
			[]Option{SetLimit(10), SetOffset(20), SetFields("name")},
			[]string{"limit 10;", "offset 20", "fields name;"},
			nil,
		},
		{
			"Single invalid option",
			[]Option{SetLimit(-99999)},
			nil,
			ErrOutOfRange,
		},
		{
			"Multiple invalid options",
			[]Option{SetLimit(-99999), SetOffset(-99999)},
			nil,
			ErrOutOfRange,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts, err := unwrapOptions(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if test.wantErr != nil {
				return
			}

			qry, err := apicalypse.Query(opts...)
			if err != nil {
				t.Fatal(err)
			}

			for _, want := range test.wantOpts {
				if !strings.Contains(qry, want) {
					t.Errorf("got: <%v>, want: <%v>", qry, want)
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
		{"Multiple non-empty fields", []string{"name", "hypes", "rating"}, "name,hypes,rating", nil},
		{"Empty fields slice", []string{}, "", ErrEmptyFields},
		{"Single empty field", []string{"  "}, "", ErrEmptyFields},
		{"Multiple empty fields", []string{"", " ", "", ""}, "", ErrEmptyFields},
		{"Mixed empty and non-empty fields", []string{"", "id", "  ", "url"}, "", ErrEmptyFields},
		{"Single expanded field", []string{"game.name"}, "", ErrExpandedField},
		{"Multiple expanded fields", []string{"game.name", "game.id"}, "", ErrExpandedField},
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
		{"Multiple non-empty fields", []string{"name", "hypes", "rating"}, "name,hypes,rating", nil},
		{"Empty fields slice", []string{}, "", ErrEmptyFields},
		{"Single empty field", []string{"  "}, "", ErrEmptyFields},
		{"Multiple empty fields", []string{"", " ", "", ""}, "", ErrEmptyFields},
		{"Mixed empty and non-empty fields", []string{"", "id", "  ", "url"}, "", ErrEmptyFields},
		{"Single expanded field", []string{"game.name"}, "", ErrExpandedField},
		{"Multiple expanded fields", []string{"game.name", "game.id"}, "", ErrExpandedField},
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

func ExampleComposeOptions() {
	c := NewClient("YOUR_CLIENT_ID", "YOUR_APP_ACCESS_TOKEN", nil)

	// Composing FuncOptions to filter for top 5 popular games
	composed := ComposeOptions(
		SetLimit(5),
		SetFields("name", "cover"),
		SetOrder("hypes", OrderDescending),
		SetFilter("category", OpEquals, "0"),
	)

	// Using composed FuncOptions
	PS4, err := c.Games.Index(
		composed,
		SetFilter("platforms", OpEquals, "48"), // only PS4 games
	)
	if err != nil {
		log.Fatal(err)
	}

	// Reusing composed FuncOptions
	XBOX, err := c.Games.Index(
		composed,
		SetFilter("platforms", OpEquals, "49"), // only XBOX games
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Top 5 PS4 Games: ")
	for _, v := range PS4 {
		fmt.Println(*v)
	}

	fmt.Println("Top 5 Xbox Games: ")
	for _, v := range XBOX {
		fmt.Println(*v)
	}
}

func ExampleSetOrder() {
	c := NewClient("YOUR_CLIENT_ID", "YOUR_APP_ACCESS_TOKEN", nil)

	// Retrieve most relevant games - default
	c.Games.Search("zelda")

	// Retrieve most hyped games
	c.Games.Search("zelda", SetOrder("hypes", OrderDescending))

	// Retrieve least hyped games
	c.Games.Search("zelda", SetOrder("hypes", OrderAscending))
}

func ExampleSetLimit() {
	c := NewClient("YOUR_CLIENT_ID", "YOUR_APP_ACCESS_TOKEN", nil)

	// Retrieve up to 10 results - default
	c.Characters.Search("snake")

	// Retrieve up to 50 results
	c.Characters.Search("snake", SetLimit(50))

	// Retrieve up to 1 result
	c.Characters.Search("snake", SetLimit(1))
}

func ExampleSetOffset() {
	c := NewClient("YOUR_CLIENT_ID", "YOUR_APP_ACCESS_TOKEN", nil)

	batchLimit := SetLimit(50)

	// Retrieve first batch of results - default
	c.Games.Index(batchLimit)

	// Retrieve second batch of results
	c.Games.Index(batchLimit, SetOffset(50))

	// Retrieve third batch of results
	c.Games.Index(batchLimit, SetOffset(100))

	// Retrieve fourth batch of results
	c.Games.Index(batchLimit, SetOffset(150))
}

func ExampleSetFields() {
	c := NewClient("YOUR_CLIENT_ID", "YOUR_APP_ACCESS_TOKEN", nil)

	// Retrieve name field
	c.Characters.Search("mario", SetFields("name"))

	// Retrieve gender field
	c.Characters.Search("mario", SetFields("gender"))

	// Retrieve both name and gender field
	c.Characters.Search("mario", SetFields("name", "gender"))

	// Retrieve whole mug_shot field
	c.Characters.Search("mario", SetFields("mug_shot"))

	// Retrieve any number of fields
	c.Characters.Search("mario", SetFields("name", "gender", "url", "species", "games", "mug_shot"))

	// Retrieve all available fields
	c.Characters.Search("mario", SetFields("*"))
}

func ExampleSetExclude() {
	c := NewClient("YOUR_CLIENT_ID", "YOUR_APP_ACCESS_TOKEN", nil)

	// Exclude name field
	c.Characters.Search("mario", SetFields("name"))

	// Exclude gender field
	c.Characters.Search("mario", SetFields("gender"))

	// Exclude both name and gender field
	c.Characters.Search("mario", SetFields("name", "gender"))

	// Exclude whole mug_shot field
	c.Characters.Search("mario", SetFields("mug_shot"))

	// Exclude any number of fields
	c.Characters.Search("mario", SetFields("name", "gender", "url", "species", "games", "mug_shot"))

	// Exclude all available fields
	c.Characters.Search("mario", SetFields("*"))
}

func ExampleSetFilter() {
	c := NewClient("YOUR_CLIENT_ID", "YOUR_APP_ACCESS_TOKEN", nil)

	// Retrieve unfiltered games - default
	c.Games.Index()

	// Retrieve games with popularity above 50
	c.Games.Index(SetFilter("hypes", OpGreaterThan, "50"))

	// Retrieve games with cover art
	c.Games.Index(SetFilter("cover", OpNotEquals, "null"))

	// Retrieve games released on PS4 (platform ID of 48)
	c.Games.Index(SetFilter("platforms", OpEquals, "48"))

	// Retrieve games whose name does not match "Horizon: Zero Dawn"
	c.Games.Index(SetFilter("name", OpNotEquals, "Horizon: Zero Dawn"))

	// Retrieve games which have the Adventure genre (Genre ID of 31)
	c.Games.Index(SetFilter("genres", OpContainsAtLeast, "31"))

	// Retrieve games that meet all the previous requirements
	c.Games.Index(
		SetFilter("hypes", OpGreaterThan, "50"),
		SetFilter("cover", OpNotEquals, "null"),
		SetFilter("platforms", OpEquals, "48"),
		SetFilter("name", OpNotEquals, "Horizon: Zero Dawn"),
		SetFilter("genres", OpContainsAtLeast, "31"),
	)
}
