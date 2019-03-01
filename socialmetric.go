package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct SocialMetric -add-tags json -w

// SocialMetric represents a particular social media metric such as
// follows, likes, shares, views, favorites, etc.
// For more information visit: https://api-docs.igdb.com/#social-metric
type SocialMetric struct {
	ID                 int                  `json:"id"`
	Category           SocialMetricCategory `json:"category"`
	CreatedAt          int                  `json:"created_at"`
	SocialMetricSource int                  `json:"social_metric_source"`
	Value              int                  `json:"value"`
}

//go:generate stringer -type=SocialMetricCategory

// SocialMetricCategory specifies a particular type of social metric.
type SocialMetricCategory int

// Expected SocialMetricCategory enums from the IGDB.
const (
	SocialFollows SocialMetricCategory = iota + 1
	SocialLikes
	SocialHates
	SocialShares
	SocialViews
	SocialComments
	SocialFavorites
)

// SocialMetricService handles all the API calls for the IGDB SocialMetric endpoint.
type SocialMetricService service

// Get returns a single SocialMetric identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any SocialMetrics, an error is returned.
func (ss *SocialMetricService) Get(id int, opts ...Option) (*SocialMetric, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var met []*SocialMetric

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ss.client.get(ss.end, &met, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get SocialMetric with ID %v", id)
	}

	return met[0], nil
}

// List returns a list of SocialMetrics identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a SocialMetric is ignored. If none of the IDs
// match a SocialMetric, an error is returned.
func (ss *SocialMetricService) List(ids []int, opts ...Option) ([]*SocialMetric, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var met []*SocialMetric

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ss.client.get(ss.end, &met, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get SocialMetrics with IDs %v", ids)
	}

	return met, nil
}

// Index returns an index of SocialMetrics based solely on the provided functional
// options used to sort, filter, and paginate the results. If no SocialMetrics can
// be found using the provided options, an error is returned.
func (ss *SocialMetricService) Index(opts ...Option) ([]*SocialMetric, error) {
	var met []*SocialMetric

	err := ss.client.get(ss.end, &met, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of SocialMetrics")
	}

	return met, nil
}

// Count returns the number of SocialMetrics available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which SocialMetrics to count.
func (ss *SocialMetricService) Count(opts ...Option) (int, error) {
	ct, err := ss.client.getCount(ss.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count SocialMetrics")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB SocialMetric object.
func (ss *SocialMetricService) Fields() ([]string, error) {
	f, err := ss.client.getFields(ss.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get SocialMetric fields")
	}

	return f, nil
}
