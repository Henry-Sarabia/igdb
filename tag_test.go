package igdb

import (
	"reflect"
	"testing"
)

func TestGenerateTag(t *testing.T) {
	var tagTests = []struct {
		Name     string
		TypeID   tagType
		ObjectID int
		ExpTag   Tag
		ExpErr   error
	}{
		{"ObjectID at zero", TagTheme, 0, 0, nil},
		{"ObjectID within range", TagGenre, 5, 268435461, nil},
		{"OjectID below range", TagKeyword, -1234, 0, ErrNegativeID},
	}

	for _, tt := range tagTests {
		t.Run(tt.Name, func(t *testing.T) {
			tag, err := GenerateTag(tt.TypeID, tt.ObjectID)
			if !reflect.DeepEqual(err, tt.ExpErr) {
				t.Fatalf("Expected error '%v', got '%v'", tt.ExpErr, err)
			}

			if tag != tt.ExpTag {
				t.Fatalf("Expected tag %d, got %d", tt.ExpTag, tag)
			}
		})
	}
}

func TestTagString(t *testing.T) {
	var tagTests = []struct {
		Name      string
		Tag       Tag
		ExpString string
	}{
		{"Zero Tag", 0, "0"},
		{"One Tag", 1, "1"},
		{"Million Tag", 1000000, "1000000"},
		{"Billion Tag", 1000000000, "1000000000"},
		{"Trillion Tag", 1000000000000, "1000000000000"},
		{"Negative Tag", -1234, "-1234"},
	}

	for _, tt := range tagTests {
		t.Run(tt.Name, func(t *testing.T) {
			actString := tt.Tag.String()
			if actString != tt.ExpString {
				t.Fatalf("Expected tag string '%s', got '%s'", tt.ExpString, actString)
			}
		})
	}
}
