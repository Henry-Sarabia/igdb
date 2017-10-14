package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

const getTitleResp = `
[{
	"id": 16549,
	"name": "Dev. Support Engineer",
	"created_at": 1433110945497,
	"updated_at": 1433110945600,
	"slug": "dev-support-engineer",
	"url": "https://www.igdb.com/titles/dev-support-engineer",
	"games": [
		1000
	]
}]
`

const getTitlesResp = `
[{
	"id": 7839,
	"name": "Web Dev",
	"created_at": 1413999800195,
	"updated_at": 1413999800195,
	"slug": "web-dev",
	"url": "https://www.igdb.com/titles/web-dev",
	"games": [
		4754
	]
},
{
	"id": 25381,
	"name": "Embedded QA",
	"created_at": 1461934414039,
	"updated_at": 1472328870944,
	"slug": "embedded-qa",
	"url": "https://www.igdb.com/titles/embedded-qa",
	"games": [
		556,
		15894,
		105,
		11582
	]
}]
`

const searchTitlesResp = `
[{
	"id": 17228,
	"name": "Senior Graphic & Interface Designer",
	"created_at": 1437825411943,
	"updated_at": 1437825412024,
	"slug": "senior-graphic-and-interface-designer",
	"url": "https://www.igdb.com/titles/senior-graphic-and-interface-designer",
	"games": [
		2995
	]
},
{
	"id": 16396,
	"name": "Graphic Coordinator",
	"created_at": 1433023246973,
	"updated_at": 1479009213820,
	"slug": "graphic-coordinator",
	"url": "https://www.igdb.com/titles/graphic-coordinator",
	"games": [
		3824,
		5418
	]
},
{
	"id": 16514,
	"name": "Character Graphic",
	"created_at": 1433023452820,
	"updated_at": 1433023454932,
	"slug": "character-graphic",
	"url": "https://www.igdb.com/titles/character-graphic",
	"games": [
		9141
	]
}]
`

func TestTitleTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	r := Title{}
	typ := reflect.ValueOf(r).Type()

	err := c.validateStruct(typ, TitleEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetTitle(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getTitleResp)
	defer ts.Close()

	ti, err := c.GetTitle(16549)
	if err != nil {
		t.Error(err)
	}

	en := "Dev. Support Engineer"
	an := ti.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eID := 16549
	aID := ti.ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	eURL := URL("https://www.igdb.com/titles/dev-support-engineer")
	aURL := ti.URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	egID := []int{1000, 1001}
	agID := ti.Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
		}
	}
}

func TestGetTitles(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getTitlesResp)
	defer ts.Close()

	ids := []int{7839, 25381}
	ti, err := c.GetTitles(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(ti)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Web Dev"
	an := ti[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	es := "web-dev"
	as := ti[0].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	eu := 1472328870944
	au := ti[1].UpdatedAt
	if au != eu {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
	}

	egID := []int{556, 15894, 105, 11582}
	agID := ti[1].Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
		}
	}
}

func TestSearchTitles(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, searchTitlesResp)
	defer ts.Close()

	ti, err := c.SearchTitles("graphic")
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(ti)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	ec := 1437825411943
	ac := ti[0].CreatedAt
	if ac != ec {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
	}

	eURL := URL("https://www.igdb.com/titles/senior-graphic-and-interface-designer")
	aURL := ti[0].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	es := "graphic-coordinator"
	as := ti[1].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	egl := 2
	agl := len(ti[1].Games)
	if agl != egl {
		t.Errorf("Expected Games lengti %d, got %d", egl, agl)
	}

	en := "Character Graphic"
	an := ti[2].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eu := 1433023454932
	au := ti[2].UpdatedAt
	if au != eu {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
	}
}
