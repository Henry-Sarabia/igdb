package igdb

import (
	"errors"
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
type OptionFunc func(*Options) error

// newOpt returns a new Options object
// mutated by the OptionFunc arguments.
// Only one of each type can be passed
// per API call. OptFilter is the only
// exception to this rule.
func newOpt(ofs ...OptionFunc) (*Options, error) {
	opt := &Options{Values: url.Values{}}

	for _, of := range ofs {
		err := of(opt)
		if err != nil {
			return nil, err
		}
	}

	return opt, nil
}

// Type order specifies in which order to place
// results from an API call. The two available
// constants of this type are Asc and Desc.
type order string

const (
	// AscendingOrder is used as an argument in the SetOrder optional function
	// to set the results from an API call in ascending order.
	AscendingOrder order = ":asc"
	// DescendingOrder is used as an argument in the SetOrder optional function
	// to set the results from an API call in descending order.
	DescendingOrder order = ":desc"
)

// OptOrder is a functional option used to set
// the order of results from an API call. The default
// ordering is based on relevance.
func OptOrder(field string, ord order) OptionFunc {
	return func(o *Options) error {
		if len(field) == 0 {
			return errors.New("field value is empty")
		}
		if o.Values.Get("order") != "" {
			return errors.New("order option already set")
		}
		o.Values.Set("order", field+string(ord))
		return nil
	}
}

// OptLimit is a functional option used to set
// the limit of results from an API call. The
// correct way to use this function is to pass
// it as a parameter to an API call. The default
// limit is 10. The maximum limit is 50.
func OptLimit(lim int) OptionFunc {
	return func(o *Options) error {
		if lim <= 0 || lim > 50 {
			return errors.New("limit value not from 1 through 50")
		}
		if o.Values.Get("limit") != "" {
			return errors.New("limit option already set")
		}
		o.Values.Set("limit", strconv.Itoa(lim))
		return nil
	}
}

// OptOffset is a functional option used to set
// the offset of results from an API call. The
// correct way to use this function is to pass
// it as a parameter to an API call. The default
// offset is 0. The maximum offset is 10,000.
// For results larger than 10,000 elements, use
// the Scroll option. This option and OptScroll
// are mutually exclusive options; only one can
// be used per API call.
func OptOffset(off int) OptionFunc {
	return func(o *Options) error {
		if off < 0 || off > 50 {
			return errors.New("offset value not from 0 through 50")
		}
		if o.Values.Get("scroll") != "" {
			return errors.New("offset option is mutually exclusive with scroll option")
		}
		if o.Values.Get("offset") != "" {
			return errors.New("offset option already set")
		}
		o.Values.Set("offset", strconv.Itoa(off))
		return nil
	}
}

// OptFields is a functional option used to specify
// which struct fields from the requested type you
// want the API response to contain. Subfields are
// accessed with a dot operator. The default
// is set to all available fields.
func OptFields(fields ...string) OptionFunc {
	return func(o *Options) error {
		if len(fields) == 0 {
			return errors.New("fields value is empty")
		}
		if o.Values.Get("fields") != "" {
			return errors.New("fields option already set")
		}
		fs := strings.Join(fields, ",")
		if prev, ok := o.Values["fields"]; ok {
			fs = prev[0] + "," + fs
		}
		o.Values.Set("fields", fs)
		return nil
	}
}

type postfix string

const (
	// Equals checks for equality. Must match exactly.
	Equals postfix = "eq"
	// NotEquals checks for inequality. Any non-exact match.
	NotEquals postfix = "not_eq"
	// GreaterThan checks if value is greater than another value. Only works on numbers.
	GreaterThan postfix = "gt"
	// GreaterThanEqual checks if value is greater than or equal to another value. Only works on numbers.
	GreaterThanEqual postfix = "gte"
	// LessThan checks if value is less than another value. Only works on numbers.
	LessThan postfix = "lt"
	// LessThanEqual checks if value is less than or equal to another value. Only works on numbers.
	LessThanEqual postfix = "lte"
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
// value as a string to specify the value of the configured filter. This
// is the only option allowed to have more than one of in a single API call.
// For more information visit https://igdb.github.io/api/references/filters/.
func OptFilter(field string, post postfix, val string) OptionFunc {
	return func(o *Options) error {
		if field == "" {
			return errors.New("field value is empty")
		}
		s := fmt.Sprintf("filter[%s][%s]", field, string(post))
		o.Values.Set(s, val)
		return nil
	}
}

// OptScroll is a functional option used
// to paginate the results of an API call
// using IGDB's Scroll API. The given
// integer denotes which page of results
// to retrieve from the API call. When not
// included in an API call, this option has
// no default value. This option and OptScroll
// are mutually exclusive options; only one can
// be used per API call.
func OptScroll(page int) OptionFunc {
	return func(o *Options) error {
		if o.Values.Get("offset") != "" {
			return errors.New("scroll option is mutually exclusive with offset option")
		}
		if o.Values.Get("scroll") != "" {
			return errors.New("scroll option already set")
		}
		o.Values.Set("scroll", strconv.Itoa(page))
		return nil
	}
}

// optSearch is an unexported functional
// option used to search the IGDB for
// the given query.
func optSearch(qry string) OptionFunc {
	return func(o *Options) error {
		if qry == "" {
			return errors.New("qry value is empty")
		}
		o.Values.Set("search", qry)
		return nil
	}
}
