package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Website -add-tags json -w

// Website represents a website and its URL; usually associated with a game.
// For more information visit: https://api-docs.igdb.com/#website
type Website struct {
	ID       int             `json:"id,omitempty"`
	Category WebsiteCategory `json:"category,omitempty"`
	Trusted  bool            `json:"trusted,omitempty"`
	URL      string          `json:"url,omitempty"`
}

// WebsiteCategory specifies a specific popular website.
type WebsiteCategory int

// Expected WebsiteCategory enums from the IGDB.
const (
	WebsiteOfficial WebsiteCategory = iota + 1
	WebsiteWikia
	WebsiteWikipedia
	WebsiteFacebook
	WebsiteTwitter
	WebsiteTwitch
	_
	WebsiteInstagram
	WebsiteYoutube
	WebsiteIphone
	WebsiteIpad
	WebsiteAndroid
	WebsiteSteam
	WebsiteReddit
	WebsiteDiscord
	WebsiteGooglePlus
	WebsiteTumblr
	WebsiteLinkedin
	WebsitePinterest
	WebsiteSoundcloud
)

// WebsiteService handles all the API calls for the IGDB Website endpoint.
type WebsiteService service

// Get returns a single Website identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Websites, an error is returned.
func (ws *WebsiteService) Get(id int, opts ...Option) (*Website, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var web []*Website

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ws.client.get(ws.end, &web, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Website with ID %v", id)
	}

	return web[0], nil
}

// List returns a list of Websites identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Website is ignored. If none of the IDs
// match a Website, an error is returned.
func (ws *WebsiteService) List(ids []int, opts ...Option) ([]*Website, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var web []*Website

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ws.client.get(ws.end, &web, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Websites with IDs %v", ids)
	}

	return web, nil
}

// Index returns an index of Websites based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Websites can
// be found using the provided options, an error is returned.
func (ws *WebsiteService) Index(opts ...Option) ([]*Website, error) {
	var web []*Website

	err := ws.client.get(ws.end, &web, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Websites")
	}

	return web, nil
}

// Count returns the number of Websites available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Websites to count.
func (ws *WebsiteService) Count(opts ...Option) (int, error) {
	ct, err := ws.client.getCount(ws.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Websites")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Website object.
func (ws *WebsiteService) Fields() ([]string, error) {
	f, err := ws.client.getFields(ws.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Website fields")
	}

	return f, nil
}
