package igdb

// Tag is a generated number that represents a specific IGDB object. Tag
// provides a quick and compact way to do complex filtering on the IGDB API.
type Tag int

// tagType represents the IGDB Object ID of a particular IGDB object type.
type tagType int

// The following tagTypes correspond to their respective IGDB Object Type IDs.
//
// For the list of these IDs and other information,
// visit: https://igdb.github.io/api/references/tag-numbers/
const (
	TagTheme tagType = iota
	TagGenre
	TagKeyword
	TagGame
	TagPerspective
)

// GenerateTag uses the ID of an IGDB object type and the ID of an IGDB 
// object to generate a Tag addressed to that object. Negative ID values 
// are considered invalid.
func GenerateTag(typeID tagType, objectID int) (Tag, error) {
	if typeID < 0 || objectID < 0 {
		return 0, ErrNegativeID
	}

	tag := int(typeID) << 28
	tag |= objectID

	return Tag(tag), nil
}
