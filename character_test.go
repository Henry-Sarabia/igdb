package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestCharacterTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	ch := Character{}
	typ := reflect.ValueOf(ch).Type()

	err := c.validateStruct(typ, CharacterEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetCharacter(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_character.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	ch, err := c.GetCharacter(10617)
	if err != nil {
		t.Error(err)
	}

	eID := 10617
	aID := ch.ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	en := "Princess Zelda"
	an := ch.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'\n", en, an)
	}

	eh := 479
	ah := ch.Mugshot.Height
	if ah != eh {
		t.Errorf("Expected height of %d, got %d\n", eh, ah)
	}
}

func TestGetCharacters(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_characters.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	ids := []int{3726, 9580}
	ch, err := c.GetCharacters(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(ch)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Mario"
	an := ch[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eAKA := "Super Mario"
	aAKA := ch[0].AKAs[0]
	if eAKA != aAKA {
		t.Errorf("Expected AKA '%s', got '%s'", eAKA, aAKA)
	}

	eID := 9580
	aID := ch[1].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	eURL := URL("https://www.igdb.com/characters/samus-aran")
	aURL := ch[1].URL
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

}

func TestSearchCharacters(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/search_characters.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	ch, err := c.SearchCharacters("snake")
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(ch)
	if al != el {
		t.Errorf("Expected length of %d, got %d\n", el, al)
	}

	eURL := URL("https://www.igdb.com/characters/snake")
	aURL := ch[0].URL
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'\n", eURL, aURL)
	}

	en := "Solid Snake"
	an := ch[1].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'\n", en, an)
	}

	eID := 5378
	aID := ch[2].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d\n", eID, aID)
	}
}
