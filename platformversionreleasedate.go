package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct PlatformVersionReleaseDate -add-tags json -w

// PlatformVersionReleaseDate describes a platform release date.
// Used to dig deeper into release dates, platforms, and versions.
// For more information visit: https://api-docs.igdb.com/#platform-version-release-date
type PlatformVersionReleaseDate struct {
	ID              int            `json:"id"`
	Category        DateCategory   `json:"category"`
	CreatedAt       int            `json:"created_at"`
	Date            int            `json:"date"`
	Human           string         `json:"human"`
	M               int            `json:"m"`
	PlatformVersion int            `json:"platform_version"`
	Region          RegionCategory `json:"region"`
	UpdatedAt       int            `json:"updated_at"`
	Y               int            `json:"y"`
}

// PlatformVersionReleaseDateService handles all the API calls for the IGDB PlatformVersionReleaseDate endpoint.
type PlatformVersionReleaseDateService service

// Get returns a single PlatformVersionReleaseDate identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PlatformVersionReleaseDates, an error is returned.
func (ps *PlatformVersionReleaseDateService) Get(id int, opts ...Option) (*PlatformVersionReleaseDate, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var date []*PlatformVersionReleaseDate

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &date, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PlatformVersionReleaseDate with ID %v", id)
	}

	return date[0], nil
}

// List returns a list of PlatformVersionReleaseDates identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a PlatformVersionReleaseDate is ignored. If none of the IDs
// match a PlatformVersionReleaseDate, an error is returned.
func (ps *PlatformVersionReleaseDateService) List(ids []int, opts ...Option) ([]*PlatformVersionReleaseDate, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var date []*PlatformVersionReleaseDate

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := ps.client.get(ps.end, &date, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PlatformVersionReleaseDates with IDs %v", ids)
	}

	return date, nil
}

// Index returns an index of PlatformVersionReleaseDates based solely on the provided functional
// options used to sort, filter, and paginate the results. If no PlatformVersionReleaseDates can
// be found using the provided options, an error is returned.
func (ps *PlatformVersionReleaseDateService) Index(opts ...Option) ([]*PlatformVersionReleaseDate, error) {
	var date []*PlatformVersionReleaseDate

	err := ps.client.get(ps.end, &date, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of PlatformVersionReleaseDates")
	}

	return date, nil
}

// Count returns the number of PlatformVersionReleaseDates available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which PlatformVersionReleaseDates to count.
func (ps *PlatformVersionReleaseDateService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count PlatformVersionReleaseDates")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB PlatformVersionReleaseDate object.
func (ps *PlatformVersionReleaseDateService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get PlatformVersionReleaseDate fields")
	}

	return f, nil
}
