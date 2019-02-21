package igdb

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

const (
	testGameGet    string = "test_data/game_get.json"
	testGameList   string = "test_data/game_list.json"
	testGameSearch string = "test_data/game_search.json"
)

func TestGameService_Get(t *testing.T) {
	f, err := ioutil.ReadFile(testGameGet)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Game, 1)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name     string
		file     string
		id       int
		opts     []Option
		wantGame *Game
		wantErr  error
	}{
		{"Valid response", testGameGet, 7346, []Option{SetFields("name")}, init[0], nil},
		{"Invalid ID", testFileEmpty, -1, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, 7346, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, 7346, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, 0, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			g, err := c.Games.Get(test.id, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(g, test.wantGame) {
				t.Errorf("got: <%v>, \nwant: <%v>", g, test.wantGame)
			}
		})
	}
}

func TestGameService_List(t *testing.T) {
	f, err := ioutil.ReadFile(testGameList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Game, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name      string
		file      string
		ids       []int
		opts      []Option
		wantGames []*Game
		wantErr   error
	}{
		{"Valid response", testGameList, []int{105842, 32478, 98774, 104945, 69530}, []Option{SetLimit(5)}, init, nil},
		{"Zero IDs", testFileEmpty, nil, nil, nil, ErrEmptyIDs},
		{"Invalid ID", testFileEmpty, []int{-500}, nil, nil, ErrNegativeID},
		{"Empty response", testFileEmpty, []int{105842, 32478, 98774, 104945, 69530}, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []int{105842, 32478, 98774, 104945, 69530}, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, []int{0, 9999999}, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			g, err := c.Games.List(test.ids, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(g, test.wantGames) {
				t.Errorf("got: <%v>, \nwant: <%v>", g, test.wantGames)
			}
		})
	}
}

func TestGameService_Index(t *testing.T) {
	f, err := ioutil.ReadFile(testGameList)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Game, 0)
	json.Unmarshal(f, &init)

	tests := []struct {
		name      string
		file      string
		opts      []Option
		wantGames []*Game
		wantErr   error
	}{
		{"Valid response", testGameList, []Option{SetLimit(5)}, init, nil},
		{"Empty response", testFileEmpty, nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			g, err := c.Games.Index(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(g, test.wantGames) {
				t.Errorf("got: <%v>, \nwant: <%v>", g, test.wantGames)
			}
		})
	}
}

func TestGameService_Search(t *testing.T) {
	f, err := ioutil.ReadFile(testGameSearch)
	if err != nil {
		t.Fatal(err)
	}

	init := make([]*Game, 0)
	json.Unmarshal(f, &init)

	var tests = []struct {
		name      string
		file      string
		qry       string
		opts      []Option
		wantGames []*Game
		wantErr   error
	}{
		{"Valid response", testGameSearch, "mario", []Option{SetLimit(5)}, init, nil},
		{"Empty query", testFileEmpty, "", []Option{SetLimit(5)}, nil, ErrEmptyQry},
		{"Empty response", testFileEmpty, "mario", nil, nil, errInvalidJSON},
		{"Invalid option", testFileEmpty, "mario", []Option{SetOffset(99999)}, nil, ErrOutOfRange},
		{"No results", testFileEmptyArray, "non-existent entry", nil, nil, ErrNoResults},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			g, err := c.Games.Search(test.qry, test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(g, test.wantGames) {
				t.Errorf("got: <%v>, \nwant: <%v>", g, test.wantGames)
			}
		})
	}
}

func TestGameService_Count(t *testing.T) {
	var tests = []struct {
		name      string
		resp      string
		opts      []Option
		wantCount int
		wantErr   error
	}{
		{"Happy path", `{"count": 100}`, []Option{SetFilter("popularity", OpGreaterThan, "75")}, 100, nil},
		{"Empty response", "", nil, 0, errInvalidJSON},
		{"Invalid option", "", []Option{SetLimit(100)}, 0, ErrOutOfRange},
		{"No results", "[]", nil, 0, ErrNoResults},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, test.resp)
			defer ts.Close()

			count, err := c.Games.Count(test.opts...)
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if count != test.wantCount {
				t.Fatalf("got: <%v>, want: <%v>", count, test.wantCount)
			}
		})
	}
}

func TestGameService_Fields(t *testing.T) {
	var tests = []struct {
		name       string
		resp       string
		wantFields []string
		wantErr    error
	}{
		{"Happy path", `["name", "slug", "url"]`, []string{"url", "slug", "name"}, nil},
		{"Dot operator", `["logo.url", "background.id"]`, []string{"background.id", "logo.url"}, nil},
		{"Asterisk", `["*"]`, []string{"*"}, nil},
		{"Empty response", "", nil, errInvalidJSON},
		{"No results", "[]", nil, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, test.resp)
			defer ts.Close()

			fields, err := c.Games.Fields()
			if errors.Cause(err) != test.wantErr {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !equalSlice(fields, test.wantFields) {
				t.Fatalf("Expected fields '%v', got '%v'", test.wantFields, fields)
			}
		})
	}
}

func ExampleGameService_Get() {
	c := NewClient("YOUR_API_KEY", nil)

	g, err := c.Games.Get(7346, SetFields("name", "url", "summary", "storyline", "rating", "popularity", "cover"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("IGDB entry for The Legend of Zelda: Breath of the Wild\n", *g)
}

func ExampleGameService_List() {
	c := NewClient("YOUR_API_KEY", nil)

	g, err := c.Games.List([]int{1721, 2777, 1074})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("IGDB entries for Megaman 8, Kirby Air Ride, and Super Mario 64")
	for _, v := range g {
		fmt.Println(*v)
	}
}

func ExampleGameService_Index() {
	c := NewClient("YOUR_API_KEY", nil)

	g, err := c.Games.Index(
		SetLimit(5),
		SetFilter("popularity", OpGreaterThan, "80"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("IGDB entries for 5 Games with popularity above 80")
	for _, v := range g {
		fmt.Println(*v)
	}
}

func ExampleGameService_Search() {
	c := NewClient("YOUR_API_KEY", nil)

	g, err := c.Games.Search(
		"mario",
		SetFields("*"),
		SetFilter("rating", OpGreaterThanEqual, "80"),
		SetOrder("rating", OrderDescending),
		SetLimit(3))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("IGDB entries for Super Mario 64, Mario Kart 8, and Mario Party")
	for _, v := range g {
		fmt.Println(*v)
	}
}

func ExampleGameService_Count() {
	c := NewClient("YOUR_API_KEY", nil)

	ct, err := c.Games.Count(SetFilter("release_dates.date", OpGreaterThan, "1993-12-15"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Number of games released after December 15, 1993: ", ct)
}

func ExampleGameService_Fields() {
	c := NewClient("YOUR_API_KEY", nil)

	fl, err := c.Games.Fields()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("List of available fields for the IGDB Game object: ", fl)
}
