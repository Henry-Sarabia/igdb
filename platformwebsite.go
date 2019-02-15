package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

// PlatformWebsite represents the main website for a particular platform.
// For more information visit: https://api-docs.igdb.com/#platform-website
type PlatformWebsite struct {
	ID       int             `json:"id"`
	Category WebsiteCategory `json:"category"`
	Trusted  bool            `json:"trusted"`
	URL      string          `json:"url"`
}

// PlatformWebsiteService handles all the API calls for the IGDB PlatformWebsite endpoint.
type PlatformWebsiteService service

// Get returns a single PlatformWebsite identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PlatformWebsites, an error is returned.
func (ps *PlatformWebsiteService) Get(id int, opts ...Option) (*PlatformWebsite, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var web []*PlatformWebsite

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &web, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PlatformWebsite with ID %v", id)
	}

	return web[0], nil
}

// List returns a list of PlatformWebsites identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a PlatformWebsite is ignored. If none of the IDs
// match a PlatformWebsite, an error is returned.
func (ps *PlatformWebsiteService) List(ids []int, opts ...Option) ([]*PlatformWebsite, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var web []*PlatformWebsite

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := ps.client.get(ps.end, &web, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PlatformWebsites with IDs %v", ids)
	}

	return web, nil
}

// Index returns an index of PlatformWebsites based solely on the provided functional
// options used to sort, filter, and paginate the results. If no PlatformWebsites can
// be found using the provided options, an error is returned.
func (ps *PlatformWebsiteService) Index(opts ...Option) ([]*PlatformWebsite, error) {
	var web []*PlatformWebsite

	err := ps.client.get(ps.end, &web, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of PlatformWebsites")
	}

	return web, nil
}

// Count returns the number of PlatformWebsites available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which PlatformWebsites to count.
func (ps *PlatformWebsiteService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count PlatformWebsites")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB PlatformWebsite object.
func (ps *PlatformWebsiteService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get PlatformWebsite fields")
	}

	return f, nil
}
