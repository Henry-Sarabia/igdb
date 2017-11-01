package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestFranchiseTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	f := Franchise{}
	typ := reflect.ValueOf(f).Type()

	err := c.validateStruct(typ, FranchiseEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetFranchise(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_franchise.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	f, err := c.GetFranchise(596)
	if err != nil {
		t.Error(err)
	}

	en := "The Legend of Zelda"
	an := f.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eURL := URL("https://www.igdb.com/franchises/the-legend-of-zelda")
	aURL := f.URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	egID := []int{11607, 1036, 18017, 18066, 7346, 25840, 8534, 41829, 1628, 9602}
	agID := f.Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
		}
	}
}

func TestGetFranchises(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_franchises.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	ids := []int{9, 22}
	f, err := c.GetFranchises(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(f)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Red Dead"
	an := f[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eURL := URL("https://www.igdb.com/franchises/red-dead")
	aURL := f[0].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	eu := 1479418914178
	au := f[1].UpdatedAt
	if au != eu {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
	}

	egID := []int{25639, 26546, 26180, 28368, 44157}
	agID := f[1].Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
		}
	}
}

func TestSearchFranchises(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/search_franchises.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	f, err := c.SearchFranchises("super")
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(f)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Super Man"
	an := f[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	ec := 1381669592350
	ac := f[0].CreatedAt
	if ac != ec {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
	}

	eID := 860
	aID := f[1].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	eURL := URL("https://www.igdb.com/franchises/super-mario")
	aURL := f[1].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	es := "marvel-super-hero-squad"
	as := f[2].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	egID := []int{4997, 5188}
	agID := f[2].Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
		}
	}
}
