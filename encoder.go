package igdb

import "strings"

// Encoder is implemented by any values that has an encode method, which
// returns the encoded format for that value. The Encode method is used
// to print a case-sensitive key value map used for query parameters or
// form values as a string.
type Encoder interface {
	Encode() string
}

// encodeURL encodes the url with the query
// parameters provided by the encoder.
func encodeURL(enc Encoder, url string) string {
	url = strings.Replace(url, " ", "", -1)

	if values := enc.Encode(); values != "" {
		url += "?" + values
	}

	return url
}
