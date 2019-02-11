package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

// CompanyWebsite represents a website for a specific company.
// For more information visit: https://api-docs.igdb.com/#company-website
type CompanyWebsite struct {
	ID       int             `json:"id"`
	Category WebsiteCategory `json:"category"`
	Trusted  bool            `json:"trusted"`
	URL      string          `json:"url"`
}

// CompanyWebsiteService handles all the API calls for the IGDB CompanyWebsite endpoint.
type CompanyWebsiteService service

// Get returns a single CompanyWebsite identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any CompanyWebsites, an error is returned.
func (zs *CompanyWebsiteService) Get(id int, opts ...FuncOption) (*CompanyWebsite, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var web []*CompanyWebsite

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := zs.client.get(zs.end, &web, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CompanyWebsite with ID %v", id)
	}

	return web[0], nil
}

// List returns a list of CompanyWebsites identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a CompanyWebsite is ignored. If none of the IDs
// match a CompanyWebsite, an error is returned.
func (zs *CompanyWebsiteService) List(ids []int, opts ...FuncOption) ([]*CompanyWebsite, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var web []*CompanyWebsite

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := zs.client.get(zs.end, &web, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CompanyWebsites with IDs %v", ids)
	}

	return web, nil
}

// Index returns an index of CompanyWebsites based solely on the provided functional
// options used to sort, filter, and paginate the results. If no CompanyWebsites can
// be found using the provided options, an error is returned.
func (zs *CompanyWebsiteService) Index(opts ...FuncOption) ([]*CompanyWebsite, error) {
	var web []*CompanyWebsite

	err := zs.client.get(zs.end, &web, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of CompanyWebsites")
	}

	return web, nil
}

// Count returns the number of CompanyWebsites available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which CompanyWebsites to count.
func (zs *CompanyWebsiteService) Count(opts ...FuncOption) (int, error) {
	ct, err := zs.client.getCount(zs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count CompanyWebsites")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB CompanyWebsite object.
func (zs *CompanyWebsiteService) Fields() ([]string, error) {
	f, err := zs.client.getFields(zs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get CompanyWebsite fields")
	}

	return f, nil
}
