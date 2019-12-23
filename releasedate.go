package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct ReleaseDate -add-tags json -w

// ReleaseDate represents the release date for a particular game.
// Used to dig deeper into release dates, platforms, and versions.
// For more information visit: https://api-docs.igdb.com/#release-date
type ReleaseDate struct {
	ID        int            `json:"id,omitempty"`
	Category  DateCategory   `json:"category,omitempty"`
	CreatedAt int            `json:"created_at,omitempty"`
	Date      int            `json:"date,omitempty"`
	Game      int            `json:"game,omitempty"`
	Human     string         `json:"human,omitempty"`
	M         int            `json:"m,omitempty"`
	Platform  int            `json:"platform,omitempty"`
	Region    RegionCategory `json:"region,omitempty"`
	UpdatedAt int            `json:"updated_at,omitempty"`
	Y         int            `json:"y,omitempty"`
}

//go:generate stringer -type=DateCategory,RegionCategory

// DateCategory specifies the format of a release date.
type DateCategory int

// Expected DateCategory enums from the IGDB.
const (
	DateYYYYMMMMDD DateCategory = iota
	DateYYYYMMMM
	DateYYYY
	DateYYYYQ1
	DateYYYYQ2
	DateYYYYQ3
	DateYYYYQ4
	DateTBD
)

// RegionCategory specifies a specific geographic region.
type RegionCategory int

// Expected RegionCategory enums from the IGDB.
const (
	RegionEurope RegionCategory = iota + 1
	RegionNorthAmerica
	RegionAustralia
	RegionNewZealand
	RegionJapan
	RegionChina
	RegionAsia
	RegionWorldwide
)

// ReleaseDateService handles all the API calls for the IGDB ReleaseDate endpoint.
type ReleaseDateService service

// Get returns a single ReleaseDate identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any ReleaseDates, an error is returned.
func (rs *ReleaseDateService) Get(id int, opts ...Option) (*ReleaseDate, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var date []*ReleaseDate

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := rs.client.get(rs.end, &date, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get ReleaseDate with ID %v", id)
	}

	return date[0], nil
}

// List returns a list of ReleaseDates identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a ReleaseDate is ignored. If none of the IDs
// match a ReleaseDate, an error is returned.
func (rs *ReleaseDateService) List(ids []int, opts ...Option) ([]*ReleaseDate, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var date []*ReleaseDate

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := rs.client.get(rs.end, &date, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get ReleaseDates with IDs %v", ids)
	}

	return date, nil
}

// Index returns an index of ReleaseDates based solely on the provided functional
// options used to sort, filter, and paginate the results. If no ReleaseDates can
// be found using the provided options, an error is returned.
func (rs *ReleaseDateService) Index(opts ...Option) ([]*ReleaseDate, error) {
	var date []*ReleaseDate

	err := rs.client.get(rs.end, &date, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of ReleaseDates")
	}

	return date, nil
}

// Count returns the number of ReleaseDates available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which ReleaseDates to count.
func (rs *ReleaseDateService) Count(opts ...Option) (int, error) {
	ct, err := rs.client.getCount(rs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count ReleaseDates")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB ReleaseDate object.
func (rs *ReleaseDateService) Fields() ([]string, error) {
	f, err := rs.client.getFields(rs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get ReleaseDate fields")
	}

	return f, nil
}
