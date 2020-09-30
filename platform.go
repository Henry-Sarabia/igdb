package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Platform -add-tags json -w

// Platform represents the hardware used to run the game
// or game delivery network.
// For more information visit: https://api-docs.igdb.com/#platform
type Platform struct {
	ID              int              `json:"id"`
	Abbreviation    string           `json:"abbreviation"`
	AlternativeName string           `json:"alternative_name"`
	Category        PlatformCategory `json:"category"`
	CreatedAt       int              `json:"created_at"`
	Generation      int              `json:"generation"`
	Name            string           `json:"name"`
	PlatformLogo    int              `json:"platform_logo"`
	ProductFamily   int              `json:"product_family"`
	Slug            string           `json:"slug"`
	Summary         string           `json:"summary"`
	UpdatedAt       int              `json:"updated_at"`
	URL             string           `json:"url"`
	Versions        []int            `json:"versions"`
	Websites        []int            `json:"websites"`
}

//go:generate stringer -type=PlatformCategory

// PlatformCategory specifies a type of platform.
type PlatformCategory int

// Expected PlatformCategory enums from the IGDB.
const (
	PlatformConsole PlatformCategory = iota + 1
	PlatformArcade
	PlatformPlatform
	PlatformOperatingSystem
	PlatformPortableConsole
	PlatformComputer
)

// PlatformService handles all the API calls for the IGDB Platform endpoint.
type PlatformService service

// Get returns a single Platform identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Platforms, an error is returned.
func (ps *PlatformService) Get(id int, opts ...Option) (*Platform, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var plat []*Platform

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.post(ps.end, &plat, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Platform with ID %v", id)
	}

	return plat[0], nil
}

// List returns a list of Platforms identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Platform is ignored. If none of the IDs
// match a Platform, an error is returned.
func (ps *PlatformService) List(ids []int, opts ...Option) ([]*Platform, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var plat []*Platform

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ps.client.post(ps.end, &plat, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Platforms with IDs %v", ids)
	}

	return plat, nil
}

// Index returns an index of Platforms based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Platforms can
// be found using the provided options, an error is returned.
func (ps *PlatformService) Index(opts ...Option) ([]*Platform, error) {
	var plat []*Platform

	err := ps.client.post(ps.end, &plat, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Platforms")
	}

	return plat, nil
}

// Search returns a list of Platforms found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no Platforms are found using the provided query, an error is returned.
func (ps *PlatformService) Search(qry string, opts ...Option) ([]*Platform, error) {
	var plat []*Platform

	opts = append(opts, setSearch(qry))
	err := ps.client.post(ps.end, &plat, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Platform with query %s", qry)
	}

	return plat, nil
}

// Count returns the number of Platforms available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Platforms to count.
func (ps *PlatformService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Platforms")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Platform object.
func (ps *PlatformService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Platform fields")
	}

	return f, nil
}
