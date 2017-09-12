package igdb

import (
	"net/http"
	"testing"
)

const getFeedResp = `
[{
	"id": 128033,
	"created_at": 1500916658000,
	"updated_at": 1500917216678,
	"url": "https://www.igdb.com/feed/2qsh",
	"content": "amiibo Functionality Revealed for Metroid: Samus Returns + Fusion Difficulty http://www.youtube.com/watch?v=sR1mAv7dcsc",
	"category": 7,
	"meta": "{\"aggregator\":\"youtube\",\"external_id\":\"sR1mAv7dcsc\"}"
}]
`

const getFeedsResp = `
[{
	"id": 62732,
	"created_at": 1494010804000,
	"updated_at": 1494013354420,
	"url": "https://www.igdb.com/feed/1cek",
	"content": "eSports Ready - Rainbow Six Siege http://www.youtube.com/watch?v=jRYVzfQz9nU",
	"category": 7,
	"meta": "{\"aggregator\":\"youtube\",\"external_id\":\"jRYVzfQz9nU\"}"
},
{
	"id": 132484,
	"created_at": 1501156914070,
	"updated_at": 1501156914070,
	"url": "https://www.igdb.com/feed/2u84",
	"category": 1,
	"pulse": 252065
},
{
	"id": 143318,
	"created_at": 1503654296057,
	"updated_at": 1503654296057,
	"url": "https://www.igdb.com/feed/32l2",
	"category": 1,
	"pulse": 261098
}]
`

func TestGetFeed(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getFeedResp)
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
		t.Errorf("Expected unix epoch of %d, got %d", eu, au)
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
	ts, c := startTestServer(http.StatusOK, getFeedsResp)
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
		t.Errorf("Expected unix epoch of %d, got %d", eu, au)
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
