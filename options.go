package igdb

import (
	"net/url"
	"strconv"
)

// Options contains a value map to store optional
// parameters for the various API calls.
type Options struct {
	Values url.Values
}

// OptionFunc is a first-order function that is
// returned by functional options and is later
// used in API calls to set individual options.
type OptionFunc func(*Options)

// Type order specifies in which order to place
// results from an API call. The two available
// constants of this type are Asc and Desc.
type order string

const (
	// Asc is used as an argument in the SetOrder optional function
	// to set the results from an API call in ascending order.
	Asc order = ":asc"
	// Desc is used as an argument in the SetOrder optional function
	// to set the results from an API call in descending order.
	Desc order = ":desc"
)

// OptOrder is a functional option used to set
// the order of results from an API call.
func OptOrder(param string, ord order) OptionFunc {
	return func(o *Options) {
		o.Values.Set("order", param+string(ord))
	}
}

// OptLimit is a functional option used to set
// the limit of results from an API call. The
// correct way to use this function is to pass
// it as a parameter to an API call.
func OptLimit(lim int) OptionFunc {
	return func(o *Options) {
		o.Values.Set("limit", strconv.Itoa(lim))
	}
}

// OptOffset is a functional option used to set
// the offset of results from an API call. The
// correct way to use this function is to pass
// it as a parameter to an API call.
func OptOffset(off int) OptionFunc {
	return func(o *Options) {
		o.Values.Set("offset", strconv.Itoa(off))
	}
}
