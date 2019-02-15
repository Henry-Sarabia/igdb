package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

// PageLogo represents the logo of a specific page.
// For more information visit: https://api-docs.igdb.com/#page-logo
type PageLogo struct {
	Image
	ID int `json:"id"`
}

// PageLogoService handles all the API calls for the IGDB PageLogo endpoint.
type PageLogoService service

// Get returns a single PageLogo identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PageLogos, an error is returned.
func (ps *PageLogoService) Get(id int, opts ...Option) (*PageLogo, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var logo []*PageLogo

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &logo, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PageLogo with ID %v", id)
	}

	return logo[0], nil
}

// List returns a list of PageLogos identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a PageLogo is ignored. If none of the IDs
// match a PageLogo, an error is returned.
func (ps *PageLogoService) List(ids []int, opts ...Option) ([]*PageLogo, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var logo []*PageLogo

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := ps.client.get(ps.end, &logo, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PageLogos with IDs %v", ids)
	}

	return logo, nil
}

// Index returns an index of PageLogos based solely on the provided functional
// options used to sort, filter, and paginate the results. If no PageLogos can
// be found using the provided options, an error is returned.
func (ps *PageLogoService) Index(opts ...Option) ([]*PageLogo, error) {
	var logo []*PageLogo

	err := ps.client.get(ps.end, &logo, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of PageLogos")
	}

	return logo, nil
}

// Count returns the number of PageLogos available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which PageLogos to count.
func (ps *PageLogoService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count PageLogos")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB PageLogo object.
func (ps *PageLogoService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get PageLogo fields")
	}

	return f, nil
}
