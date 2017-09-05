package igdb

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
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

// newOpt returns a basic Options object
func newOpt() Options {
	return Options{Values: url.Values{}}
}

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
// the order of results from an API call. The default
// ordering is based on relevance.
func OptOrder(field string, ord order) OptionFunc {
	return func(o *Options) {
		o.Values.Set("order", field+string(ord))
	}
}

// OptLimit is a functional option used to set
// the limit of results from an API call. The
// correct way to use this function is to pass
// it as a parameter to an API call. The default
// limit is 10.
func OptLimit(lim int) OptionFunc {
	return func(o *Options) {
		o.Values.Set("limit", strconv.Itoa(lim))
	}
}

// OptOffset is a functional option used to set
// the offset of results from an API call. The
// correct way to use this function is to pass
// it as a parameter to an API call. The default
// offset is 0.
func OptOffset(off int) OptionFunc {
	return func(o *Options) {
		o.Values.Set("offset", strconv.Itoa(off))
	}
}

// OptFields is a functional option used to specify
// which struct fields from the requested type you
// want the API response to contain. Subfields are
// accessed with a dot operator. The default
// is set to all available fields.
func OptFields(fields ...string) OptionFunc {
	return func(o *Options) {
		if len(fields) == 0 {
			return
		}
		fs := strings.Join(fields, ",")
		if prev, ok := o.Values["fields"]; ok {
			fs = prev[0] + "," + fs
		}
		o.Values.Set("fields", fs)
	}
}

type postfix string

const (
	// EQ stands for equal. Must match exactly.
	EQ postfix = "eq"
	// NotEQ stands for not equal. Any non-exact match.
	NotEQ postfix = "not_eq"
	// GT stands for greater than. Only works on numbers.
	GT postfix = "gt"
	// GTE stands for greater than or equal. Only works on numbers.
	GTE postfix = "gte"
	// LT stands for less than. Only works on numbers.
	LT postfix = "lt"
	// LTE stands for less than or equal. Only works on numbers.
	LTE postfix = "lte"
	// Prefix only works on strings.
	Prefix postfix = "prefix"
	// Exists checks for a non-null value.
	Exists postfix = "exists"
	// NotExists checks for a null value.
	NotExists postfix = "not_exists"
	// In checks if the value exists within an array and between values.
	In postfix = "in"
	// NotIn checks if the values do not not exist within an array and between values.
	NotIn postfix = "not_in"
	// Any checks if the value has any within the array or between values.
	Any postfix = "any"
)

// OptFilter is a functional option used to filter the results from
// an API call. Provide a field name to specify what property you
// want to filter with. Provide a postfix to specify how you want
// to filter the results using the given field name. Provide a concrete
// value as a string to specify the value of the configured filter.
// For more information visit https://igdb.github.io/api/references/filters/.
func OptFilter(field string, post postfix, val string) OptionFunc {
	return func(o *Options) {
		s := fmt.Sprintf("filter[%s][%s]", field, string(post))
		o.Values.Set(s, val)
	}
}
