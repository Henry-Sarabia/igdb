package igdb

import (
	"net/http"
	"testing"
)

func TestReleaseDatesGet(t *testing.T) {
	var releaseDateTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/releasedates_get.txt", 1073, ""},
		{"Invalid ID", "test_data/empty.txt", -1000, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 1073, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range releaseDateTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			rd, err := c.ReleaseDates.Get(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			eID := 1073
			aID := rd.ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			eg := 475
			ag := rd.Game
			if ag != eg {
				t.Errorf("Expected Game ID %d, got %d", eg, ag)
			}

			ec := 1303935024000
			ac := rd.CreatedAt
			if ac != ec {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
			}
		})
	}
}

func TestReleaseDatesList(t *testing.T) {
	var releaseDateTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/releasedates_list.txt", []int{62566, 32350, 1077}, []FuncOption{SetLimit(5)}, ""},
		{"Zero IDs", "test_data/releasedates_list.txt", nil, nil, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-123}, nil, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", []int{62566, 32350, 1077}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{62566, 32350, 1077}, []FuncOption{SetOffset(99999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{0, 9999999}, nil, ErrNoResults.Error()},
	}
	for _, tt := range releaseDateTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			rd, err := c.ReleaseDates.List(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 3
			al := len(rd)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			ec := DateCategory(2)
			ac := rd[0].Category
			if ac != ec {
				t.Errorf("Expected date category %d, got %d", ec, ac)
			}

			ep := 27
			ap := rd[0].Platform
			if ap != ep {
				t.Errorf("Expected platform %d, got %d", ep, ap)
			}

			ed := 978220800000
			ad := rd[1].Date
			if ad != ed {
				t.Errorf("Expected Unix time in milliseconds %d, got %d", ed, ad)
			}

			ey := 2000
			ay := rd[1].Year
			if ay != ey {
				t.Errorf("Expected year %d, got %d", ey, ay)
			}

			em := 10
			am := rd[2].Month
			if am != em {
				t.Errorf("Expected month %d, got %d", em, am)
			}

			eh := "2010-Oct-26"
			ah := rd[2].Human
			if ah != eh {
				t.Errorf("Expected date '%s', got '%s'", eh, ah)
			}
		})
	}
}

func TestReleaseDatesCount(t *testing.T) {
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

			count, err := c.ReleaseDates.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestReleaseDatesListFields(t *testing.T) {
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

			fields, err := c.ReleaseDates.ListFields()
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
