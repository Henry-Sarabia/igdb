package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

// PageWebsite represents the website of a specific page.
// For more information visit: https://api-docs.igdb.com/#page-website
type PageWebsite struct {
	Website
	ID int `json:"id"`
}

// PageWebsiteService handles all the API calls for the IGDB PageWebsite endpoint.
type PageWebsiteService service

// Get returns a single PageWebsite identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PageWebsites, an error is returned.
func (ps *PageWebsiteService) Get(id int, opts ...Option) (*PageWebsite, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var site []*PageWebsite

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &site, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PageWebsite with ID %v", id)
	}

	return site[0], nil
}

// List returns a list of PageWebsites identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a PageWebsite is ignored. If none of the IDs
// match a PageWebsite, an error is returned.
func (ps *PageWebsiteService) List(ids []int, opts ...Option) ([]*PageWebsite, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var site []*PageWebsite

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := ps.client.get(ps.end, &site, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PageWebsites with IDs %v", ids)
	}

	return site, nil
}

// Index returns an index of PageWebsites based solely on the provided functional
// options used to sort, filter, and paginate the results. If no PageWebsites can
// be found using the provided options, an error is returned.
func (ps *PageWebsiteService) Index(opts ...Option) ([]*PageWebsite, error) {
	var site []*PageWebsite

	err := ps.client.get(ps.end, &site, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of PageWebsites")
	}

	return site, nil
}

// Count returns the number of PageWebsites available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which PageWebsites to count.
func (ps *PageWebsiteService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count PageWebsites")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB PageWebsite object.
func (ps *PageWebsiteService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get PageWebsite fields")
	}

	return f, nil
}
