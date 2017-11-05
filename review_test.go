package igdb

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestReviewTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	r := Review{}
	typ := reflect.ValueOf(r).Type()

	err := c.validateStruct(typ, ReviewEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetReview(t *testing.T) {
	var reviewTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/get_review.txt", 1462, ""},
		{"Invalid ID", "test_data/empty.txt", -1000, ErrNegativeID.Error()},
		{"Empty Response", "test_data/empty.txt", 1462, errEndOfJSON.Error()},
	}
	for _, tt := range reviewTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			r, err := c.GetReview(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			et := "Almost perfect!"
			at := r.Title
			if at != et {
				t.Errorf("Expected title '%s', got '%s'", et, at)
			}

			eURL := URL("https://www.igdb.com/games/mario-kart-8/reviews/almost-perfect")
			aURL := r.URL
			if eURL != aURL {
				t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
			}

			ev := 250
			av := r.Views
			if av != ev {
				t.Errorf("Expected view count %d, got %d", ev, av)
			}
		})
	}
}

func TestGetReviews(t *testing.T) {
	var reviewTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/get_reviews.txt", []int{1571, 65}, []OptionFunc{OptLimit(5)}, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-1500}, nil, ErrNegativeID.Error()},
		{"Zero IDs", "test_data/empty.txt", nil, nil, ErrEmptyIDs.Error()},
		{"Empty Response", "test_data/empty.txt", []int{1571, 65}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{1571, 65}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range reviewTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			r, err := c.GetReviews(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(r)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			ec := 1
			ac := r[0].Category
			if ac != ec {
				t.Errorf("Expected date category %d, got %d", ec, ac)
			}

			erc := 3
			arc := r[0].RatingCategory
			if arc != erc {
				t.Errorf("Expected rating category %d, got %d", erc, arc)
			}

			ep := 41
			ap := r[0].Platform
			if ap != ep {
				t.Errorf("Expected platform %d, got %d", ep, ap)
			}

			eu := "ZUPERFLY"
			au := r[1].Username
			if au != eu {
				t.Errorf("Expected username '%s', got '%s'", eu, au)
			}

			ePos := "-smash balls\r\n-Subspace Emissary\r\n-gameplay\r\n-custom levels"
			aPos := r[1].PositivePoints
			if aPos != ePos {
				t.Errorf("Expected positive points '%s', got '%s'", ePos, aPos)
			}

			eNeg := "-timers\r\n-Mario"
			aNeg := r[1].NegativePoints
			if aNeg != eNeg {
				t.Errorf("Expected negative points '%s', got '%s'", eNeg, aNeg)
			}
		})
	}
}

func TestSearchReviews(t *testing.T) {
	var reviewTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/search_reviews.txt", "zelda", []OptionFunc{OptLimit(50)}, ""},
		{"Empty query", "test_data/search_reviews.txt", "", []OptionFunc{OptLimit(50)}, ""},
		{"Empty response", "test_data/empty.txt", "zelda", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "zelda", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range reviewTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			r, err := c.SearchReviews(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 3
			al := len(r)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eID := 80
			aID := r[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			es := "zelda-reviewathon-number-1"
			as := r[0].Slug
			if as != es {
				t.Errorf("Expected slug '%s', got '%s'", es, as)
			}

			eIn := "In the first review I started explaining my 'Nostalgia Vortex' theory."
			aIn := r[1].Introduction
			if !strings.Contains(aIn, eIn) {
				t.Errorf("Expected Introduction to contain '%s', got '%s'", eIn, aIn)
			}

			eCont := "I know every last inch of that country and there is no exit."
			aCont := r[1].Content
			if !strings.Contains(aCont, eCont) {
				t.Errorf("Expected Content to contain '%s', got '%s'", eCont, aCont)
			}

			eConc := "but it was also an abrupt change of tone that some gamers might be put off by."
			aConc := r[2].Conclusion
			if !strings.Contains(aConc, eConc) {
				t.Errorf("Expected Conclusion to contain '%s', got '%s'", eConc, aConc)
			}

			elc := 2
			alc := r[2].Likes
			if alc != elc {
				t.Errorf("Expected Likes count %d, got %d", elc, alc)
			}
		})
	}
}
