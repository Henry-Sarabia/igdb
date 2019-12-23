package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct ReviewVideo -add-tags json -w

// ReviewVideo represents a user-created review video.
// For more information visit: https://api-docs.igdb.com/#review-video
type ReviewVideo struct {
	ID      int    `json:"id,omitempty"`
	Trusted bool   `json:"trusted,omitempty"`
	URL     string `json:"url,omitempty"`
}

// ReviewVideoService handles all the API calls for the IGDB ReviewVideo endpoint.
type ReviewVideoService service

// Get returns a single ReviewVideo identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any ReviewVideos, an error is returned.
func (rs *ReviewVideoService) Get(id int, opts ...Option) (*ReviewVideo, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var vid []*ReviewVideo

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := rs.client.get(rs.end, &vid, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get ReviewVideo with ID %v", id)
	}

	return vid[0], nil
}

// List returns a list of ReviewVideos identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a ReviewVideo is ignored. If none of the IDs
// match a ReviewVideo, an error is returned.
func (rs *ReviewVideoService) List(ids []int, opts ...Option) ([]*ReviewVideo, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var vid []*ReviewVideo

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := rs.client.get(rs.end, &vid, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get ReviewVideos with IDs %v", ids)
	}

	return vid, nil
}

// Index returns an index of ReviewVideos based solely on the provided functional
// options used to sort, filter, and paginate the results. If no ReviewVideos can
// be found using the provided options, an error is returned.
func (rs *ReviewVideoService) Index(opts ...Option) ([]*ReviewVideo, error) {
	var vid []*ReviewVideo

	err := rs.client.get(rs.end, &vid, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of ReviewVideos")
	}

	return vid, nil
}

// Count returns the number of ReviewVideos available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which ReviewVideos to count.
func (rs *ReviewVideoService) Count(opts ...Option) (int, error) {
	ct, err := rs.client.getCount(rs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count ReviewVideos")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB ReviewVideo object.
func (rs *ReviewVideoService) Fields() ([]string, error) {
	f, err := rs.client.getFields(rs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get ReviewVideo fields")
	}

	return f, nil
}
