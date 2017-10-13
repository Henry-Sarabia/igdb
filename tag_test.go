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
		{"OjectID below range", TagKeyword, -1234, 0, ErrOutOfRange},
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
