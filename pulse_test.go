package igdb

import (
	"net/http"
	"testing"
)

func TestPulsesGet(t *testing.T) {
	var pulseTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/pulses_get.txt", 145346, ""},
		{"Invalid ID", "test_data/empty.txt", -25000, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 145346, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range pulseTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			p, err := c.Pulses.Get(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			et := "Nintendo announces new Mario, Zelda amiibo"
			at := p.Title
			if at != et {
				t.Errorf("Expected title '%s', got '%s'", et, at)
			}

			ep := 2
			ap := p.PulseSource
			if ap != ep {
				t.Errorf("Expected Pulse Source %d, got %d", ep, ap)
			}

			eURL := URL("//images.igdb.com/igdb/image/upload/t_thumb/i1fti435exzyu1ydftu4.jpg")
			aURL := p.PulseImage.URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			etID := []Tag{1, 17, 38, 268435468, 268435487, 536871422, 536872221}
			atID := p.Tags
			for i := range atID {
				if atID[i] != etID[i] {
					t.Errorf("Expected Tag ID %d, got %d\n", etID[i], atID[i])
				}
			}
		})
	}
}

func TestPulsesList(t *testing.T) {
	var pulseTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/pulses_list.txt", []int{132354, 257394, 109415}, []FuncOption{SetLimit(5)}, ""},
		{"Zero IDs", "test_data/pulses_list.txt", nil, nil, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-250000}, nil, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", []int{132354, 257394, 109415}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{132354, 257394, 109415}, []FuncOption{SetOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{0, 9999999}, nil, ErrNoResults.Error()},
	}
	for _, tt := range pulseTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			p, err := c.Pulses.List(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 3
			al := len(p)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			et := "Battleborn: a great game with a fatal flaw"
			at := p[0].Title
			if at != et {
				t.Errorf("Expected title '%s', got '%s'", et, at)
			}

			ea := "Darren Nakamura"
			aa := p[0].Author
			if aa != ea {
				t.Errorf("Expected slug '%s', got '%s'", ea, aa)
			}

			eUID := "5fc0c5269a2aa7887d1a2c13a27c5bd2"
			aUID := p[1].UID
			if aUID != eUID {
				t.Errorf("Expected ID '%s', got '%s'", eUID, aUID)
			}

			eURL := URL("//images.igdb.com/igdb/image/upload/t_thumb/ibmrifg0uxp8w3y6hfzo.jpg")
			aURL := p[1].PulseImage.URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			ep := 8
			ap := p[2].PulseSource
			if ap != ep {
				t.Errorf("Expected Pulse Source %d, got %d", ep, ap)
			}

			etID := []Tag{1, 18, 27, 268435461, 268435468, 268435471, 536871198}
			atID := p[2].Tags
			for i := range atID {
				if atID[i] != etID[i] {
					t.Errorf("Expected Tag ID %d, got %d\n", etID[i], atID[i])
				}
			}
		})
	}
}

func TestPulsesSearch(t *testing.T) {
	var pulseTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []FuncOption
		ExpErr string
	}{
		{"Happy path", "test_data/pulses_search.txt", "megaman", []FuncOption{SetLimit(50)}, ""},
		{"Empty query", "test_data/empty.txt", "", []FuncOption{SetLimit(50)}, ErrEmptyQuery.Error()},
		{"Empty response", "test_data/empty.txt", "megaman", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "megaman", []FuncOption{SetOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", "non-existant entry", nil, ErrNoResults.Error()},
	}
	for _, tt := range pulseTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			p, err := c.Pulses.Search(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 3
			al := len(p)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eCat := 10
			aCat := p[0].Category
			if aCat != eCat {
				t.Errorf("Expected category %d, got %d", eCat, aCat)
			}

			ec := 1502176226691
			ac := p[0].CreatedAt
			if ac != ec {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
			}

			et := "Retroid talks Mega Man tonight"
			at := p[1].Title
			if at != et {
				t.Errorf("Expected title '%s', got '%s'", et, at)
			}

			eURL := URL("http://feedproxy.google.com/~r/Destructoid/~3/ChzHjztL10M/retroid-talks-mega-man-tonight-456282.phtml")
			aURL := p[1].URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			ep := 1479339000000
			ap := p[2].PublishedAt
			if ap != ep {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", ep, ap)
			}

			eID := "y4wwcqqbkuyeteoq4l2n"
			aID := p[2].PulseImage.ID
			if aID != eID {
				t.Errorf("Expected ID '%s', got '%s'", eID, aID)
			}
		})
	}
}

func TestPulsesCount(t *testing.T) {
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

			count, err := c.Pulses.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestPulsesListFields(t *testing.T) {
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

			fields, err := c.Pulses.ListFields()
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
