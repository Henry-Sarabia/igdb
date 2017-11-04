package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestPerspectiveTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	p := Perspective{}
	typ := reflect.ValueOf(p).Type()

	err := c.validateStruct(typ, PerspectiveEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPerspective(t *testing.T) {
	var perspectiveTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/get_perspective.txt", 7, ""},
		{"Invalid ID", "test_data/empty.txt", -10, ErrNegativeID.Error()},
		{"Empty Response", "test_data/empty.txt", 7, errEndOfJSON.Error()},
	}
	for _, tt := range perspectiveTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			p, err := c.GetPerspective(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			en := "Virtual Reality"
			an := p.Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eu := 1462288484243
			au := p.UpdatedAt
			if au != eu {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
			}

			eURL := URL("https://www.igdb.com/player_perspectives/virtual-reality")
			aURL := p.URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			egID := []int{8654, 10724, 6415, 5639, 9254, 17244, 18157, 17986, 12302}
			agID := p.Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestGetPerspectives(t *testing.T) {
	var perspectiveTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/get_perspectives.txt", []int{6, 3}, []OptionFunc{OptLimit(5)}, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-10}, nil, ErrNegativeID.Error()},
		{"Zero IDs", "test_data/empty.txt", nil, nil, ErrEmptyIDs.Error()},
		{"Empty Response", "test_data/empty.txt", []int{6, 3}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{6, 3}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range perspectiveTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			p, err := c.GetPerspectives(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(p)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			en := "Aural"
			an := p[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			es := "aural"
			as := p[0].Slug
			if as != es {
				t.Errorf("Expected slug '%s', got '%s'", es, as)
			}

			eID := 3
			aID := p[1].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			egID := []int{6, 5, 35, 36, 13, 14, 17, 12, 76}
			agID := p[1].Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestSearchPerspectives(t *testing.T) {
	var perspectiveTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/search_perspectives.txt", "person", []OptionFunc{OptLimit(50)}, ""},
		{"Empty query", "test_data/search_perspectives.txt", "", []OptionFunc{OptLimit(50)}, ""},
		{"Empty response", "test_data/empty.txt", "person", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "person", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range perspectiveTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			p, err := c.SearchPerspectives(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(p)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eID := 1
			aID := p[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			ec := 1298968658000
			ac := p[0].CreatedAt
			if ac != ec {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
			}

			eURL := URL("https://www.igdb.com/player_perspectives/third-person")
			aURL := p[1].URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			es := "third-person"
			as := p[1].Slug
			if as != es {
				t.Errorf("Expected slug '%s', got '%s'", es, as)
			}
		})
	}
}
