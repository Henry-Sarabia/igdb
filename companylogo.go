package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

// CompanyLogo represents the logo of a developer or publisher.
// For more information visit: https://api-docs.igdb.com/#company-logo
type CompanyLogo struct {
	Image
	ID int `json:"id"`
}

// CompanyLogoService handles all the API calls for the IGDB CompanyLogo endpoint.
type CompanyLogoService service

// Get returns a single CompanyLogo identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any CompanyLogos, an error is returned.
func (cs *CompanyLogoService) Get(id int, opts ...Option) (*CompanyLogo, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var logo []*CompanyLogo

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.get(cs.end, &logo, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CompanyLogo with ID %v", id)
	}

	return logo[0], nil
}

// List returns a list of CompanyLogos identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a CompanyLogo is ignored. If none of the IDs
// match a CompanyLogo, an error is returned.
func (cs *CompanyLogoService) List(ids []int, opts ...Option) ([]*CompanyLogo, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var logo []*CompanyLogo

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := cs.client.get(cs.end, &logo, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get CompanyLogos with IDs %v", ids)
	}

	return logo, nil
}

// Index returns an index of CompanyLogos based solely on the provided functional
// options used to sort, filter, and paginate the results. If no CompanyLogos can
// be found using the provided options, an error is returned.
func (cs *CompanyLogoService) Index(opts ...Option) ([]*CompanyLogo, error) {
	var logo []*CompanyLogo

	err := cs.client.get(cs.end, &logo, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of CompanyLogos")
	}

	return logo, nil
}

// Count returns the number of CompanyLogos available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which CompanyLogos to count.
func (cs *CompanyLogoService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count CompanyLogos")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB CompanyLogo object.
func (cs *CompanyLogoService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get CompanyLogo fields")
	}

	return f, nil
}
