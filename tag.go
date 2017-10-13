package igdb

// Tag numbers are generated numbers which
// provide a compact and fast way to do
// complex filtering on the IGDB API.
type Tag int

// tagType represents the IGDB Object ID
// of a particular IGDB object type.
type tagType int

// The following tagTypes correspond
// to their respective IGDB Object
// Type IDs. For the list of these
// IDs and more information, visit:
// https://igdb.github.io/api/references/tag-numbers/
const (
	TagTheme tagType = iota
	TagGenre
	TagKeyword
	TagGame
	TagPerspective
)

// GenerateTag uses the ID of an IGDB object type and
// the ID of an IGDB object to generate a Tag addressed
// to that object. Negative ID values are considered
// invalid.
func GenerateTag(typeID tagType, objectID int) (Tag, error) {
	if typeID < 0 || objectID < 0 {
		return 0, ErrOutOfRange
	}

	tag := int(typeID) << 28
	tag |= objectID

	return Tag(tag), nil
}
