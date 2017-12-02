package igdb

import (
	"net/http"
	"testing"
)

func TestVersionsGet(t *testing.T) {
	var versionTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/versions_get.txt", 1, ""},
		{"Invalid ID", "test_data/empty.txt", -10, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 1, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range versionTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()
			v, err := c.Versions.Get(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			eID := tt.ID
			aID := v.ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			eURL := URL("https://www.igdb.com/games/assassins-creed-origins/")
			aURL := v.URL
			if aURL != eURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			el := 19
			al := len(v.Features)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}
		})
	}
}

func TestVersionsList(t *testing.T) {
	var versionTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/versions_list.txt", []int{100, 200}, []FuncOption{SetLimit(5)}, ""},
		{"Zero IDs", "test_data/versions_list.txt", nil, nil, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-500}, nil, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", []int{100, 200}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{100, 200}, []FuncOption{SetOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{100, 200}, nil, ErrNoResults.Error()},
	}
	for _, tt := range versionTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			v, err := c.Versions.List(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(v)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eg := 7346
			ag := v[0].Game
			if ag != eg {
				t.Errorf("Expected Game ID %d, got %d", eg, ag)
			}

			egID := []int{1748, 1751, 1749, 1743}
			agID := v[0].Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}

			eID := 200
			aID := v[1].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			et := "Big Daddy Statuette"
			at := v[1].Features[0].Title
			if at != et {
				t.Errorf("Expected title '%s', got '%s'", et, at)
			}

		})
	}

}

func TestVersionsCount(t *testing.T) {
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

			count, err := c.Versions.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestVersionsListFields(t *testing.T) {
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

			fields, err := c.Versions.ListFields()
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
