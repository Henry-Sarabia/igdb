package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestFeedTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	f := Feed{}
	typ := reflect.ValueOf(f).Type()

	err := c.validateStruct(typ, FeedEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetFeed(t *testing.T) {
	var feedTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/get_feed.txt", 128033, ""},
		{"Invalid ID", "test_data/empty.txt", -123, ErrNegativeID.Error()},
		{"Empty Response", "test_data/empty.txt", 128033128033, errEndOfJSON.Error()},
	}
	for _, tt := range feedTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			f, err := c.GetFeed(tt.ID)
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

func TestGetFeeds(t *testing.T) {
	var feedTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/get_feeds.txt", []int{62732, 132484, 143318}, []OptionFunc{OptLimit(5)}, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-123}, nil, ErrNegativeID.Error()},
		{"Zero IDs", "test_data/empty.txt", nil, nil, ErrEmptyIDs.Error()},
		{"Empty Response", "test_data/empty.txt", []int{62732, 132484, 143318}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{62732, 132484, 143318}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range feedTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			f, err := c.GetFeeds(tt.IDs, tt.Opts...)
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

func TestSearchFeeds(t *testing.T) {
	var feedTests = []struct {
		Name   string
		Resp   string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/search_feeds.txt", []OptionFunc{OptLimit(3)}, ""},
		{"Empty response", "test_data/empty.txt", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range feedTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			f, err := c.SearchFeeds(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 3
			al := len(f)
			if al != el {
				t.Errorf("Expected lfth of %d, got %d", el, al)
			}

			eID := 48651
			aID := f[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			ec := "Minecraft Hide N Seek! HIDE FROM THE NEIGHBOR! (Minecraft HELLO NEIGHBOR MINIGAME!) http://www.youtube.com/watch?v=kyYz_r-DQ8g"
			ac := f[0].Content
			if ac != ec {
				t.Errorf("Expected content '%s', got '%s'", ec, ac)
			}

			eu := 1492105740351
			au := f[1].UpdatedAt
			if au != eu {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
			}

			eURL := URL("https://www.igdb.com/feed/11jg")
			aURL := f[1].URL
			if aURL != eURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			eCat := FeedCategory(7)
			aCat := f[2].Category
			if aCat != eCat {
				t.Errorf("Expected category %d, got %d", eCat, aCat)
			}

			em := "{\"aggregator\":\"youtube\",\"external_id\":\"tLZpWNItcq0\"}"
			am := f[2].Meta
			if am != em {
				t.Errorf("Expected meta '%s', got '%s'", em, am)
			}
		})
	}
}
