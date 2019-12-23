package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct PlatformVersion -add-tags json -w

// PlatformVersion represents a particular version of a platform.
// For more information visit: https://api-docs.igdb.com/#platform-version
type PlatformVersion struct {
	ID                          int    `json:"id,omitempty"`
	Companies                   []int  `json:"companies,omitempty"`
	Connectivity                string `json:"connectivity,omitempty"`
	CPU                         string `json:"cpu,omitempty"`
	Graphics                    string `json:"graphics,omitempty"`
	MainManufacturer            int    `json:"main_manufacturer,omitempty"`
	Media                       string `json:"media,omitempty"`
	Memory                      string `json:"memory,omitempty"`
	Name                        string `json:"name,omitempty"`
	OS                          string `json:"os,omitempty"`
	Output                      string `json:"output,omitempty"`
	PlatformLogo                int    `json:"platform_logo,omitempty"`
	PlatformVersionReleaseDates []int  `json:"platform_version_release_dates,omitempty"`
	Resolutions                 string `json:"resolutions,omitempty"`
	Slug                        string `json:"slug,omitempty"`
	Sound                       string `json:"sound,omitempty"`
	Storage                     string `json:"storage,omitempty"`
	Summary                     string `json:"summary,omitempty"`
	URL                         string `json:"url,omitempty"`
}

// PlatformVersionService handles all the API calls for the IGDB PlatformVersion endpoint.
type PlatformVersionService service

// Get returns a single PlatformVersion identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PlatformVersions, an error is returned.
func (ps *PlatformVersionService) Get(id int, opts ...Option) (*PlatformVersion, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var ver []*PlatformVersion

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &ver, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PlatformVersion with ID %v", id)
	}

	return ver[0], nil
}

// List returns a list of PlatformVersions identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a PlatformVersion is ignored. If none of the IDs
// match a PlatformVersion, an error is returned.
func (ps *PlatformVersionService) List(ids []int, opts ...Option) ([]*PlatformVersion, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var ver []*PlatformVersion

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ps.client.get(ps.end, &ver, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PlatformVersions with IDs %v", ids)
	}

	return ver, nil
}

// Index returns an index of PlatformVersions based solely on the provided functional
// options used to sort, filter, and paginate the results. If no PlatformVersions can
// be found using the provided options, an error is returned.
func (ps *PlatformVersionService) Index(opts ...Option) ([]*PlatformVersion, error) {
	var ver []*PlatformVersion

	err := ps.client.get(ps.end, &ver, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of PlatformVersions")
	}

	return ver, nil
}

// Count returns the number of PlatformVersions available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which PlatformVersions to count.
func (ps *PlatformVersionService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count PlatformVersions")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB PlatformVersion object.
func (ps *PlatformVersionService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get PlatformVersion fields")
	}

	return f, nil
}
