package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

const getThemeResp = `
[{
	"id": 17,
	"name": "Fantasy",
	"created_at": 1322605338000,
	"updated_at": 1323289216000,
	"slug": "fantasy",
	"url": "https://www.igdb.com/themes/fantasy",
	"games": [
		799,
		651,
		901,
		929,
		939,
		800,
		931,
		876
	]
}]
`

const getThemesResp = `
[{
	"id": 20,
	"name": "Thriller",
	"created_at": 1322605338000,
	"updated_at": 1323289216000,
	"slug": "thriller",
	"url": "https://www.igdb.com/themes/thriller",
	"games": [
		18,
		19,
		21,
		493,
		109,
		955,
		844
	]
},
{
	"id": 23,
	"name": "Stealth",
	"created_at": 1322605338000,
	"updated_at": 1323289216000,
	"slug": "stealth",
	"url": "https://www.igdb.com/themes/stealth",
	"games": [
		4,
		820,
		43,
		500,
		501,
		433,
		250,
		377,
		375
	]
}]
`

const searchThemesResp = `
[{
	"id": 19,
	"name": "Horror",
	"created_at": 1322605338000,
	"updated_at": 1323289216000,
	"slug": "horror",
	"url": "https://www.igdb.com/themes/horror",
	"games": [
		967,
		973,
		971,
		970,
		861,
		985,
		37
	]
}]
`

func TestThemeTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	r := Theme{}
	typ := reflect.ValueOf(r).Type()

	err := c.validateStruct(typ, ThemeEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetTheme(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getThemeResp)
	defer ts.Close()

	th, err := c.GetTheme(17)
	if err != nil {
		t.Error(err)
	}

	en := "Fantasy"
	an := th.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eID := 17
	aID := th.ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	eURL := URL("https://www.igdb.com/themes/fantasy")
	aURL := th.URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	egID := []int{799, 651, 901, 929, 939, 800, 931, 876}
	agID := th.Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
		}
	}
}

func TestGetThemes(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getThemesResp)
	defer ts.Close()

	ids := []int{20, 23}
	th, err := c.GetThemes(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(th)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Thriller"
	an := th[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	es := "thriller"
	as := th[0].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	eu := 1323289216000
	au := th[1].UpdatedAt
	if au != eu {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
	}

	egID := []int{4, 820, 43, 500, 501, 433, 250, 377, 375}
	agID := th[1].Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
		}
	}
}

func TestSearchThemes(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, searchThemesResp)
	defer ts.Close()

	th, err := c.SearchThemes("horror")
	if err != nil {
		t.Error(err)
	}

	el := 1
	al := len(th)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	ec := 1322605338000
	ac := th[0].CreatedAt
	if ac != ec {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
	}

	eURL := URL("https://www.igdb.com/themes/horror")
	aURL := th[0].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	es := "horror"
	as := th[0].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	egl := 7
	agl := len(th[0].Games)
	if agl != egl {
		t.Errorf("Expected Games length %d, got %d", egl, agl)
	}
}
