package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestPulseSourceTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	ps := PulseSource{}
	typ := reflect.ValueOf(ps).Type()

	err := c.validateStruct(typ, PulseSourceEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestPulseSourcesGet(t *testing.T) {
	var pulseSourceTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/pulsesources_get.txt", 4943, ""},
		{"Invalid ID", "test_data/empty.txt", -4900, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 4943, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range pulseSourceTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ps, err := c.PulseSources.Get(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			en := "Kotaku"
			an := ps.Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eID := 1
			aID := ps.ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			ep := 501
			ap := ps.Page
			if ap != ep {
				t.Errorf("Expected Page ID %d, got %d", ep, ap)
			}
		})
	}
}

func TestPulseSourcesList(t *testing.T) {
	var pulseSourceTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/pulsesources_list.txt", []int{6, 10}, []OptionFunc{OptLimit(5)}, ""},
		{"Zero IDs", "test_data/pulsesources_list.txt", nil, nil, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-10}, nil, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", []int{6, 10}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{6, 10}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{0, 9999999}, nil, ErrNoResults.Error()},
	}
	for _, tt := range pulseSourceTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ps, err := c.PulseSources.List(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(ps)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			en := "Escapist"
			an := ps[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eID := 6
			aID := ps[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			ep := 552
			ap := ps[1].Page
			if ap != ep {
				t.Errorf("Expected Page ID %d, got %d", ep, ap)
			}
		})
	}
}

func TestPulseSourcesSearch(t *testing.T) {
	var pulseSourceTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/pulsesources_search.txt", "game", []OptionFunc{OptLimit(50)}, ""},
		{"Empty query", "test_data/empty.txt", "", []OptionFunc{OptLimit(50)}, ErrEmptyQuery.Error()},
		{"Empty response", "test_data/empty.txt", "game", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "game", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", "non-existant entry", nil, ErrNoResults.Error()},
	}
	for _, tt := range pulseSourceTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ps, err := c.PulseSources.Search(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 3
			al := len(ps)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eID := 3
			aID := ps[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			en := "GameIndustry.biz"
			an := ps[1].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			ep := 564
			ap := ps[2].Page
			if ap != ep {
				t.Errorf("Expected Page ID %d, got %d", ep, ap)
			}
		})
	}
}

func TestPulseSourcesCount(t *testing.T) {
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

			count, err := c.PulseSources.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestPulseSourcesListFields(t *testing.T) {
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

			fields, err := c.PulseSources.ListFields()
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
