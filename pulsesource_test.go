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

func TestGetPulseSource(t *testing.T) {
	var pulseSourceTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/get_pulsesource.txt", 4943, ""},
		{"Invalid ID", "test_data/empty.txt", -4900, ErrNegativeID.Error()},
		{"Empty Response", "test_data/empty.txt", 4943, errEndOfJSON.Error()},
	}
	for _, tt := range pulseSourceTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ps, err := c.GetPulseSource(tt.ID)
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

func TestGetPulseSources(t *testing.T) {
	var pulseSourceTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/get_pulsesources.txt", []int{6, 10}, []OptionFunc{OptLimit(5)}, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-10}, nil, ErrNegativeID.Error()},
		{"Zero IDs", "test_data/empty.txt", nil, nil, ErrEmptyIDs.Error()},
		{"Empty Response", "test_data/empty.txt", []int{6, 10}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{6, 10}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range pulseSourceTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ps, err := c.GetPulseSources(tt.IDs, tt.Opts...)
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

func TestSearchPulseSources(t *testing.T) {
	var pulseSourceTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/search_pulsesources.txt", "game", []OptionFunc{OptLimit(50)}, ""},
		{"Empty query", "test_data/search_pulsesources.txt", "", []OptionFunc{OptLimit(50)}, ""},
		{"Empty response", "test_data/empty.txt", "game", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "game", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range pulseSourceTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			ps, err := c.SearchPulseSources(tt.Qry, tt.Opts...)
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
