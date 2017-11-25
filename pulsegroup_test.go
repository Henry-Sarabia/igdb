package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestPulseGroupTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	pg := PulseGroup{}
	typ := reflect.ValueOf(pg).Type()

	err := c.validateStruct(typ, PulseGroupEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestPulseGroupsGet(t *testing.T) {
	var pulseGroupTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/pulsegroups_get.txt", 4943, ""},
		{"Invalid ID", "test_data/empty.txt", -4000, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 4943, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range pulseGroupTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			pg, err := c.PulseGroups.Get(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			en := "DOTA 2"
			an := pg.Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			eID := 4943
			aID := pg.ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			etl := 14
			atl := len(pg.Tags)
			if atl != etl {
				t.Errorf("Expected Tags length %d, got %d", etl, atl)
			}

			ep := 224467
			ap := pg.Pulses[0]
			if ap != ep {
				t.Errorf("Expected Pulse ID %d, got %d", ep, ap)
			}
		})
	}
}

func TestPulseGroupsList(t *testing.T) {
	var pulseGroupTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/pulsegroups_list.txt", []int{2096, 1108}, []OptionFunc{OptLimit(5)}, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-1000}, nil, ErrNegativeID.Error()},
		{"Zero IDs", "test_data/pulsegroups_list.txt", nil, nil, ""},
		{"Empty response", "test_data/empty.txt", []int{2096, 1108}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{2096, 1108}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{0, 9999999}, nil, ErrNoResults.Error()},
	}
	for _, tt := range pulseGroupTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			pg, err := c.PulseGroups.List(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(pg)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			en := "Battleborn"
			an := pg[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			epID := []int{268536, 268550, 268533}
			apID := pg[0].Pulses
			for i := range apID {
				if apID[i] != epID[i] {
					t.Errorf("Expected Pulse ID %d, got %d", epID[i], apID[i])
				}
			}

			ec := 1
			ac := pg[1].Category
			if ac != ec {
				t.Errorf("Expected category %d, got %d", ec, ac)
			}

			eu := 1501184656575
			au := pg[1].UpdatedAt
			if au != eu {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
			}
		})
	}
}

func TestPulseGroupsSearch(t *testing.T) {
	var pulseGroupTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/pulsegroups_search.txt", "League of Legends", []OptionFunc{OptLimit(50)}, ""},
		{"Empty query", "test_data/empty.txt", "", []OptionFunc{OptLimit(50)}, ErrEmptyQuery.Error()},
		{"Empty response", "test_data/empty.txt", "League of Legends", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "League of Legends", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", "non-existant entry", nil, ErrNoResults.Error()},
	}
	for _, tt := range pulseGroupTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			pg, err := c.PulseGroups.Search(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 3
			al := len(pg)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eID := 3907
			aID := pg[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			en := "League of Legends"
			an := pg[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			ec := 1500792572855
			ac := pg[1].CreatedAt
			if ac != ec {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
			}

			eu := 1500792572855
			au := pg[1].UpdatedAt
			if au != eu {
				t.Errorf("Expected Unix time in milliseconds of %d, %d", eu, au)
			}

			eg := 115
			ag := pg[2].Game
			if ag != eg {
				t.Errorf("Expected Game ID %d, got %d", eg, ag)
			}

			etl := 9
			atl := len(pg[2].Tags)
			if atl != etl {
				t.Errorf("Expected Tags length %d, got %d", etl, atl)
			}
		})
	}
}

func TestPulseGroupsCount(t *testing.T) {
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

			count, err := c.PulseGroups.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestPulseGroupsListFields(t *testing.T) {
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

			fields, err := c.PulseGroups.ListFields()
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
