package igdb

type Artwork struct {
	AlphaChannel bool   `json:"alpha_channel,omitempty"`
	Animated     bool   `json:"animated,omitempty"`
	Height       int    `json:"height,omitempty"`
	ImageID      string `json:"image_id,omitempty"`
	URL          string `json:"url,omitempty"`
	Width        int    `json:"width,omitempty"`
}

//func (c *Client) GetArtwork(opts ...FuncOption) ([]Artwork, error) {
//	req, err := c.request(GameEndpoint, opts...)
//	if err != nil {
//		return nil, err
//	}
//
//	var a []Artwork
//
//	err = c.do(req, &a)
//	if err != nil {
//		return nil, errors.Wrap(err, "cannot make request")
//	}
//
//	return a, nil
//}
