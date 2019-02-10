package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

// PlatformLogo represents a logo for a particular platform.
// For more information visit: https://api-docs.igdb.com/#platform-logo
type PlatformLogo struct {
	Image
	ID int `json:"id"`
}

// PlatformLogoService handles all the API calls for the IGDB PlatformLogo endpoint.
type PlatformLogoService service

// Get returns a single PlatformLogo identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PlatformLogos, an error is returned.
func (ps *PlatformLogoService) Get(id int, opts ...FuncOption) (*PlatformLogo, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var logo []*PlatformLogo

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &logo, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PlatformLogo with ID %v", id)
	}

	return logo[0], nil
}

// List returns a list of PlatformLogos identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a PlatformLogo is ignored. If none of the IDs
// match a PlatformLogo, an error is returned.
func (ps *PlatformLogoService) List(ids []int, opts ...FuncOption) ([]*PlatformLogo, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var logo []*PlatformLogo

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := ps.client.get(ps.end, &logo, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PlatformLogos with IDs %v", ids)
	}

	return logo, nil
}

// Index returns an index of PlatformLogos based solely on the provided functional
// options used to sort, filter, and paginate the results. If no PlatformLogos can
// be found using the provided options, an error is returned.
func (ps *PlatformLogoService) Index(opts ...FuncOption) ([]*PlatformLogo, error) {
	var logo []*PlatformLogo

	err := ps.client.get(ps.end, &logo, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of PlatformLogos")
	}

	return logo, nil
}

// Count returns the number of PlatformLogos available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which PlatformLogos to count.
func (ps *PlatformLogoService) Count(opts ...FuncOption) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count PlatformLogos")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB PlatformLogo object.
func (ps *PlatformLogoService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get PlatformLogo fields")
	}

	return f, nil
}
