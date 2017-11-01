package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestCollectionTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	col := Collection{}
	typ := reflect.ValueOf(col).Type()

	err := c.validateStruct(typ, CollectionEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetCollection(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_collection.txt")
	if err != nil {
		t.Fatal(err)
	}

	defer ts.Close()

	col, err := c.GetCollection(2404)
	if err != nil {
		t.Error(err)
	}

	en := "Chocobo"
	an := col.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eURL := URL("https://www.igdb.com/collections/chocobo")
	aURL := col.URL
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	eID := []int{22896, 22909, 22905, 22903, 22907, 22906, 22898, 22908, 22901}
	aID := col.Games
	for i := range aID {
		if aID[i] != eID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", eID[i], aID[i])
		}
	}
}

func TestGetCollections(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_collections.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	ids := []int{338, 1}
	col, err := c.GetCollections(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(col)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Mega Man X"
	an := col[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	ec := 1352059968884
	ac := col[0].CreatedAt
	if ac != ec {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
	}

	eURL := URL("https://www.igdb.com/collections/bioshock")
	aURL := col[1].URL
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	eID := []int{538, 14543, 19839, 20, 10047, 21}
	aID := col[1].Games
	for i := range aID {
		if aID[i] != eID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", eID[i], aID[i])
		}
	}
}

func TestSearchCollections(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/search_collections.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	col, err := c.SearchCollections("mario")
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(col)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	eID := 593
	aID := col[0].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	en := "Paper Mario"
	an := col[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	ec := 1388973829025
	ac := col[1].CreatedAt
	if ac != ec {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
	}

	es := "mario-tennis"
	as := col[1].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	eURL := URL("https://www.igdb.com/collections/mario-golf")
	aURL := col[2].URL
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	eIDs := []int{3402, 3400, 3404, 3401, 3405, 3403}
	aIDs := col[2].Games
	for i := range aIDs {
		if aIDs[i] != eIDs[i] {
			t.Errorf("Expected Game ID %d, got %d\n", eIDs[i], aIDs[i])
		}
	}
}
