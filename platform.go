package igdb

import (
	"strconv"
	"strings"
)

// RegionCode codes
type RegionCode int

// PlatformDate type
type PlatformDate struct {
	Date   int        `json:"date"` //unix epoch
	Region RegionCode `json:"region"`
}

// PlatformCompany type
type PlatformCompany struct {
	Company int    `json:"company"`
	Comment string `json:"comment"`
}

// PlatformVersion type
type PlatformVersion struct {
	ID            int               `json:"id"`
	Name          string            `json:"name"`
	Slug          string            `json:"slug"`
	URL           URL               `json:"url"`
	Manufacturer  string            `json:"manufacturer"`
	OS            string            `json:"os"`
	CPU           string            `json:"cpu"`
	Media         string            `json:"media"`
	Memory        string            `json:"memory"`
	Online        string            `json:"online"`
	Output        string            `json:"output"`
	Storage       string            `json:"storage"`
	Graphics      string            `json:"graphics"`
	Resolutions   string            `json:"resolutions"`
	Connectivity  string            `json:"connectivity"`
	Sound         string            `json:"sound"`
	Logo          Image             `json:"logo"`
	Summary       string            `json:"summary"`
	ReleaseDates  []PlatformDate    `json:"release_dates"`
	Developers    []PlatformCompany `json:"developers"`
	Manufacturers []PlatformCompany `json:"manufacturers"`
}

// Platform type
type Platform struct {
	ID            int               `json:"id"`
	Name          string            `json:"name"`
	Slug          string            `json:"slug"`
	URL           URL               `json:"url"`
	CreatedAt     int               `json:"created_at"` //unix epoch
	UpdatedAt     int               `json:"updated_at"` //unix epoch
	Shortcut      string            `json:"shortcut"`
	Logo          Image             `json:"logo"`
	Website       string            `json:"website"`
	Summary       string            `json:"summary"`
	AltName       string            `json:"alternative_name"`
	Generation    int               `json:"generation"`
	Category      int               `json:"category"`
	ProductFamily int               `json:"product_family"`
	Games         []int             `json:"games"`
	Versions      []PlatformVersion `json:"versions"`
}

// GetPlatform gets IGDB information for a platform identified by its unique IGDB ID.
func (c *Client) GetPlatform(id int, opts ...OptionFunc) (*Platform, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "platforms/" + strconv.Itoa(id)
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var p []Platform

	err := c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return &p[0], nil
}

// GetPlatforms gets IGDB information for a list of platforms identified by their
// unique IGDB IDs.
func (c *Client) GetPlatforms(ids []int, opts ...OptionFunc) ([]*Platform, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	str := intsToStrings(ids)
	url := c.rootURL + "platforms/" + strings.Join(str, ",")
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "?" + values
		}
	}

	var p []*Platform

	err := c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// SearchPlatforms searches the IGDB using the given query and returns IGDB information
// for the results. Use functional options for pagination and to sort results by parameter.
func (c *Client) SearchPlatforms(qry string, opts ...OptionFunc) ([]*Platform, error) {
	opt := newOpt()
	for _, optFunc := range opts {
		optFunc(&opt)
	}

	url := c.rootURL + "platforms/?search=" + qry
	if opts != nil {
		if values := opt.Values.Encode(); values != "" {
			url += "&" + values
		}
	}

	var p []*Platform

	err := c.get(url, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
