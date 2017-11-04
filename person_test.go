package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestPersonTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	p := Person{}
	typ := reflect.ValueOf(p).Type()

	err := c.validateStruct(typ, PersonEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPerson(t *testing.T) {
	var personTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/get_person.txt", 2107, ""},
		{"Invalid ID", "test_data/empty.txt", -2000, ErrNegativeID.Error()},
		{"Empty Response", "test_data/empty.txt", 2107, errEndOfJSON.Error()},
	}
	for _, tt := range personTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			p, err := c.GetPerson(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			en := "Shigeru Miyamoto"
			an := p.Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eDOB := -540432000000
			aDOB := p.DOB
			if aDOB != eDOB {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", eDOB, aDOB)
			}

			eURL := URL("//images.igdb.com/igdb/image/upload/t_thumb/wuyjvwyascmcquyf4qh9.jpg")
			aURL := p.Mugshot.URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			egID := []int{2909, 2777, 2476, 2923, 2350, 7337, 1073, 1070, 1036, 1074, 3365}
			agID := p.Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestGetPersons(t *testing.T) {
	var personTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/get_persons.txt", []int{52302, 84908}, []OptionFunc{OptLimit(5)}, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-50000}, nil, ErrNegativeID.Error()},
		{"Zero IDs", "test_data/empty.txt", nil, nil, ErrEmptyIDs.Error()},
		{"Empty Response", "test_data/empty.txt", []int{52302, 84908}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{52302, 84908}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range personTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			p, err := c.GetPersons(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(p)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			en := "Sean Murray"
			an := p[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			es := "sean-murray"
			as := p[0].Slug
			if as != es {
				t.Errorf("Expected slug '%s', got '%s'", es, as)
			}

			eg := GenderCode(0)
			ag := p[1].Gender
			if ag != eg {
				t.Errorf("Expected Gender code %d, got %d", eg, ag)
			}

			egID := []int{5845, 496, 1348, 1913, 1430, 1035, 11131, 1723, 6914}
			agID := p[1].Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestSearchPersons(t *testing.T) {
	var personTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/search_persons.txt", "hideokojima", []OptionFunc{OptLimit(50)}, ""},
		{"Empty query", "test_data/search_persons.txt", "", []OptionFunc{OptLimit(50)}, ""},
		{"Empty response", "test_data/empty.txt", "hideokojima", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "hideokojima", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range personTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			p, err := c.SearchPersons(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(p)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eID := 34056
			aID := p[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			ec := CountryCode(392)
			ac := p[0].Country
			if ac != ec {
				t.Errorf("Expected Country code %d, got %d", ec, ac)
			}

			ev := []int{5328, 1985}
			av := p[0].VoiceActed
			for i := range av {
				if av[i] != ev[i] {
					t.Errorf("Expected Game ID %d, got %d\n", ev[i], av[i])
				}
			}

			eURL := URL("https://www.igdb.com/people/hideo-kohima")
			aURL := p[1].URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			es := "hideo-kohima"
			as := p[1].Slug
			if as != es {
				t.Errorf("Expected slug '%s', got '%s'", es, as)
			}

			eg := GenderCode(0)
			ag := p[1].Gender
			if ag != eg {
				t.Errorf("Expected Gender code %d, got %d", eg, ag)
			}
		})
	}
}
