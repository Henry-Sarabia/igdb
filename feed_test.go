package igdb

import (
	"net/http"
	"testing"
)

func TestFeedsGet(t *testing.T) {
	var feedTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/feeds_get.txt", 128033, ""},
		{"Invalid ID", "test_data/empty.txt", -123, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 128033128033, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range feedTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			f, err := c.Feeds.Get(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			eID := 128033
			aID := f.ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			eu := 1500917216678
			au := f.UpdatedAt
			if au != eu {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
			}

			eURL := URL("https://www.igdb.com/feed/2qsh")
			aURL := f.URL
			if aURL != eURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			eCat := FeedCategory(7)
			aCat := f.Category
			if aCat != eCat {
				t.Errorf("Expected category %d, got %d", eCat, aCat)
			}

			em := "{\"aggregator\":\"youtube\",\"external_id\":\"sR1mAv7dcsc\"}"
			am := f.Meta
			if am != em {
				t.Errorf("Expected meta '%s', got '%s'", em, am)
			}
		})
	}
}

func TestFeedsList(t *testing.T) {
	var feedTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/feeds_list.txt", []int{62732, 132484, 143318}, []FuncOption{SetLimit(5)}, ""},
		{"Zero IDs", "test_data/feeds_list.txt", nil, nil, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-123}, nil, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", []int{62732, 132484, 143318}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{62732, 132484, 143318}, []FuncOption{SetOffset(99999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{0, 9999999}, nil, ErrNoResults.Error()},
	}
	for _, tt := range feedTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			f, err := c.Feeds.List(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 3
			al := len(f)
			if al != el {
				t.Errorf("Expected lfth of %d, got %d", el, al)
			}

			eID := 62732
			aID := f[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			ec := "eSports Ready - Rainbow Six Siege http://www.youtube.com/watch?v=jRYVzfQz9nU"
			ac := f[0].Content
			if ac != ec {
				t.Errorf("Expected content '%s', got '%s'", ec, ac)
			}

			eu := 1501156914070
			au := f[1].UpdatedAt
			if au != eu {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
			}

			eURL := URL("https://www.igdb.com/feed/2u84")
			aURL := f[1].URL
			if aURL != eURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			eCat := FeedCategory(1)
			aCat := f[2].Category
			if aCat != eCat {
				t.Errorf("Expected category %d, got %d", eCat, aCat)
			}

			ep := 261098
			ap := f[2].Pulse
			if ap != ep {
				t.Errorf("Expected Pulse ID %d, got %d", ep, ap)
			}
		})
	}

}

func TestFeedsCount(t *testing.T) {
	var countTests = []struct {
		Name     string
		Resp     string
		Opts     []FuncOption
		ExpCount int
		ExpErr   string
	}{
		{"Happy path", `{"count": 100}`, []FuncOption{SetFilter("popularity", OpGreaterThan, "75")}, 100, ""},
		{"Empty response", "", nil, 0, errEndOfJSON.Error()},
		{"Invalid option", "", []FuncOption{SetLimit(100)}, 0, ErrOutOfRange.Error()},
		{"No results", "[]", nil, 0, ErrNoResults.Error()},
	}

	for _, tt := range countTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, tt.Resp)
			defer ts.Close()

			count, err := c.Feeds.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestFeedsListFields(t *testing.T) {
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

			fields, err := c.Feeds.ListFields()
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
