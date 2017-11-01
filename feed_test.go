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
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_feed.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	f, err := c.GetFeed(132482)
	if err != nil {
		t.Error(err)
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
}

func TestGetFeeds(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_feeds.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	ids := []int{62732, 132484, 143318}
	f, err := c.GetFeeds(ids)
	if err != nil {
		t.Error(err)
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

	em := "eSports Ready - Rainbow Six Siege http://www.youtube.com/watch?v=jRYVzfQz9nU"
	am := f[0].Content
	if am != em {
		t.Errorf("Expected content '%s', got '%s'", em, am)
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
}
