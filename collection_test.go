package igdb

import (
	"net/http"
	"testing"
)

func TestCollectionsGet(t *testing.T) {
	var collectionTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/collections_get.txt", 2404, ""},
		{"Invalid ID", "test_data/empty.txt", -500, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 2404, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range collectionTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}

			defer ts.Close()

			col, err := c.Collections.Get(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
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
		})
	}
}

func TestCollectionsList(t *testing.T) {
	var collectionTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/collections_list.txt", []int{338, 1}, []OptionFunc{OptLimit(5)}, ""},
		{"Zero IDs", "test_data/collections_list.txt", nil, nil, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-123}, nil, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", []int{338, 1}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{338, 1}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{0, 9999999}, nil, ErrNoResults.Error()},
	}
	for _, tt := range collectionTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			col, err := c.Collections.List(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
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
		})
	}
}

func TestCollectionsSearch(t *testing.T) {
	var collectionTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/collections_search.txt", "mario", []OptionFunc{OptLimit(50)}, ""},
		{"Empty query", "test_data/empty.txt", "", []OptionFunc{OptLimit(50)}, ErrEmptyQuery.Error()},
		{"Empty response", "test_data/empty.txt", "mario", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "mario", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", "non-existant entry", nil, ErrNoResults.Error()},
	}
	for _, tt := range collectionTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			col, err := c.Collections.Search(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
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
		})
	}
}

func TestCollectionsCount(t *testing.T) {
	var countTests = []struct {
		Name     string
		Resp     string
		Opts     []OptionFunc
		ExpCount int
		ExpErr   string
	}{
		{"Happy path", `{"count": 100}`, []OptionFunc{OptFilter("popularity", OpGreaterThan, "75")}, 100, ""},
		{"Empty response", "", nil, 0, errEndOfJSON.Error()},
		{"Invalid option", "", []OptionFunc{OptLimit(100)}, 0, ErrOutOfRange.Error()},
		{"No results", "[]", nil, 0, ErrNoResults.Error()},
	}

	for _, tt := range countTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, tt.Resp)
			defer ts.Close()

			count, err := c.Collections.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestCollectionsListFields(t *testing.T) {
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

			fields, err := c.Collections.ListFields()
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
