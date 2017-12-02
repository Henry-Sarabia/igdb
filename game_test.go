package igdb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGamesGet(t *testing.T) {
	var gameTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/games_get.txt", 7346, ""},
		{"Invalid ID", "test_data/empty.txt", -1, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 7346, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range gameTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			g, err := c.Games.Get(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			eID := 7346
			aID := g.ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			en := "The Legend of Zelda: Breath of the Wild"
			an := g.Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'\n", en, an)
			}

			er := 98.5
			ar := g.AggregateRating
			if ar != er {
				t.Errorf("Expected rating of %f, got %f", er, ar)
			}

			ew := 2709
			aw := g.Covers.Width
			if aw != ew {
				t.Errorf("Expected width of %d, got %d\n", ew, aw)
			}

			var ev = []YoutubeVideo{
				{"Trailer", "Z6BeAtdoELY"},
				{"Trailer", "1rPxiXXxftE"},
				{"Trailer", "vDFZIUdo764"},
				{"Trailer", "zw47_q9wbBE"}}
			av := g.Videos
			for i := range av {
				if av[i] != ev[i] {
					t.Errorf("Expected video %v, got video %v\n", ev[i], av[i])
				}
			}
		})
	}
}

func TestGamesList(t *testing.T) {
	var gameTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/games_list.txt", []int{1721, 2777}, []FuncOption{OptLimit(5)}, ""},
		{"Zero IDs", "test_data/games_list.txt", nil, nil, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-500}, nil, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", []int{1721, 2777}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{1721, 2777}, []FuncOption{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{0, 9999999}, nil, ErrNoResults.Error()},
	}
	for _, tt := range gameTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			g, err := c.Games.List(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(g)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eURL := URL("https://www.igdb.com/games/mega-man-8")
			aURL := g[0].URL
			if aURL != eURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			ec := 1352059102919
			ac := g[0].CreatedAt
			if ac != ec {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
			}

			ed := 762
			ad := g[1].Developers[0]
			if ad != ed {
				t.Errorf("Expected developer ID %d, got %d", ed, ad)
			}

			eID := "etjab1sgankzyq6p6qgf"
			aID := g[1].Covers.ID
			if aID != eID {
				t.Errorf("Expected cloudinary ID '%s', got '%s'", eID, aID)
			}
		})
	}
}

func TestGamesSearch(t *testing.T) {
	var gameTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/games_search.txt", "mario", []FuncOption{OptLimit(50)}, ""},
		{"Empty query", "test_data/empty.txt", "", []FuncOption{OptLimit(50)}, ErrEmptyQuery.Error()},
		{"Empty response", "test_data/empty.txt", "mario", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "mario", []FuncOption{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", "non-existant entry", nil, ErrNoResults.Error()},
	}
	for _, tt := range gameTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			g, err := c.Games.Search(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 3
			al := len(g)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			en := "Super Mario 64"
			an := g[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			es := "super-mario-64"
			as := g[0].Slug
			if as != es {
				t.Errorf("Expected slug '%s', got '%s'", es, as)
			}

			er := 85.5273965373542
			ar := g[1].Rating
			if ar != er {
				t.Errorf("Expected rating of %f, got %f", er, ar)
			}

			ep := 3.666666666666667
			ap := g[1].Popularity
			if ap != ep {
				t.Errorf("Expected popularity of %f, got %f", ep, ap)
			}

			ed := 864
			ad := g[2].Developers[0]
			if ad != ed {
				t.Errorf("Expected developer ID %d, got %d", ed, ad)
			}

			eURL := URL("//images.igdb.com/igdb/image/upload/t_thumb/clmh270eov5rimiggwrk.jpg")
			aURL := g[2].Covers.URL
			if aURL != eURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}
		})
	}
}

func TestGamesCount(t *testing.T) {
	var countTests = []struct {
		Name     string
		Resp     string
		Opts     []FuncOption
		ExpCount int
		ExpErr   string
	}{
		{"Happy path", `{"count": 100}`, []FuncOption{OptFilter("popularity", OpGreaterThan, "75")}, 100, ""},
		{"Empty response", "", nil, 0, errEndOfJSON.Error()},
		{"Invalid option", "", []FuncOption{OptLimit(100)}, 0, ErrOutOfRange.Error()},
		{"No results", "[]", nil, 0, ErrNoResults.Error()},
	}

	for _, tt := range countTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, tt.Resp)
			defer ts.Close()

			count, err := c.Games.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestGamesListFields(t *testing.T) {
	var fieldTests = []struct {
		Name      string
		Resp      string
		ExpFields []string
		ExpErr    string
	}{
		{"Happy path", `["name", "slug", "url"]`, []string{"url", "slug", "name"}, ""},
		{"Dot operator", `["logo.url", "background.id"]`, []string{"background.id", "logo.url"}, ""},
		{"Asterisk", `["*"]`, []string{"*"}, ""},
		{"Empty response", "", nil, errEndOfJSON.Error()},
		{"No results", "[]", nil, ""},
	}

	for _, tt := range fieldTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, tt.Resp)
			defer ts.Close()

			fields, err := c.Games.ListFields()
			assertError(t, err, tt.ExpErr)

			ok, err := equalSlice(fields, tt.ExpFields)
			if err != nil {
				t.Fatal(err)
			}
			if !ok {
				t.Fatalf("Expected fields '%v', got '%v'", tt.ExpFields, fields)
			}
		})
	}
}

func ExampleGameService_Get() {
	c := NewClient(APIkey, nil)

	g, err := c.Games.Get(7346, OptFields("name", "url", "summary", "storyline", "rating", "popularity", "cover"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("IGDB entry for The Legend of Zelda: Breath of the Wild\n", *g)
}

func ExampleGameService_List() {
	c := NewClient(APIkey, nil)

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

func ExampleGameService_Search() {
	c := NewClient("YOUR_API_KEY", nil)

	g, err := c.Games.Search("mario",
		OptFields("*"),
		OptFilter("rating", OpGreaterThanEqual, "80"),
		OptOrder("rating", OrderDescending),
		OptLimit(3))
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

	ct, err := c.Games.Count(OptFilter("release_dates.date", OpGreaterThan, "1993-12-15"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Number of games released after December 15, 1993: ", ct)
}

func ExampleGameService_ListFields() {
	c := NewClient("YOUR_API_KEY", nil)

	fl, err := c.Games.ListFields()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("List of available fields for the IGDB Game object: ", fl)
}
