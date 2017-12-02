package igdb

// ReleaseDateService handles all the API
// calls for the IGDB ReleaseDate endpoint.
type ReleaseDateService service

// ReleaseDate contains information on an IGDB entry for a particular release
// date. ReleaseDate is used primarily as an extension to the Game endpoint.
// ReleaseDate does not support the search function.
//
// For more information, visit: https://igdb.github.io/api/endpoints/release-date/
type ReleaseDate struct {
	ID        int          `json:"id"`
	Game      int          `json:"game"`
	Category  DateCategory `json:"category"`
	Platform  int          `json:"platform"`
	Human     string       `json:"human"`
	UpdatedAt int          `json:"updated_at"` // Unix time in milliseconds
	CreatedAt int          `json:"created_at"` // Unix time in milliseconds
	Date      int          `json:"date"`       // Unix time in milliseconds
	Region    int          `json:"region"`
	Year      int          `json:"y"`
	Month     int          `json:"m"`
}

// Get returns a single ReleaseDate identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any ReleaseDates, an error is returned.
func (rds *ReleaseDateService) Get(id int, opts ...FuncOption) (*ReleaseDate, error) {
	url, err := rds.client.singleURL(ReleaseDateEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var rd []ReleaseDate

	err = rds.client.get(url, &rd)
	if err != nil {
		return nil, err
	}

	return &rd[0], nil
}

// List returns a list of ReleaseDates identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate  the results. Omitting
// IDs will instead retrieve an index of ReleaseDates based solely on the provided
// options. Any ID that does not match a ReleaseDate is ignored. If none of the IDs
// match a ReleaseDate, an error is returned.
func (rds *ReleaseDateService) List(ids []int, opts ...FuncOption) ([]*ReleaseDate, error) {
	url, err := rds.client.multiURL(ReleaseDateEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var rd []*ReleaseDate

	err = rds.client.get(url, &rd)
	if err != nil {
		return nil, err
	}

	return rd, nil
}

// Count returns the number of ReleaseDates available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which ReleaseDates to count.
func (rds *ReleaseDateService) Count(opts ...FuncOption) (int, error) {
	ct, err := rds.client.getEndpointCount(ReleaseDateEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB ReleaseDate object.
func (rds *ReleaseDateService) ListFields() ([]string, error) {
	fl, err := rds.client.getEndpointFieldList(ReleaseDateEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
