package igdb

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Errors returned by an OptionFunc.
var (
	ErrOptionSet       = errors.New("igdb.OptionFunc: option already set")
	ErrEmptyField      = errors.New("igdb.OptionFunc: field value empty")
	ErrEmptySlice      = errors.New("igdb.OptionFunc: slice empty")
	ErrOutOfRange      = errors.New("igdb.OptionFunc: value out of range")
	ErrExclusiveOption = errors.New("igdb.OptionFunc: multiple set options are mutually exclusive")
	ErrEmptyQuery      = errors.New("igdb.OptionFunc: query value empty")
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
			return ErrEmptyField
		}
		if o.Values.Get("order") != "" {
			return ErrOptionSet
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
			return ErrOutOfRange
		}
		if o.Values.Get("limit") != "" {
			return ErrOptionSet
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
			return ErrOutOfRange
		}
		if o.Values.Get("scroll") != "" {
			return ErrExclusiveOption
		}
		if o.Values.Get("offset") != "" {
			return ErrOptionSet
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
			return ErrEmptySlice
		}
		for _, f := range fields {
			if f == "" {
				return ErrEmptyField
			}
		}
		if o.Values.Get("fields") != "" {
			return ErrOptionSet
		}
		fs := strings.Join(fields, ",")
		if prev, ok := o.Values["fields"]; ok {
			fs = prev[0] + "," + fs
		}
		o.Values.Set("fields", fs)
		return nil
	}
}

type operator string

const (
	// OpEquals checks for equality. Must match exactly.
	OpEquals operator = "eq"
	// OpNotEquals checks for inequality. Any non-exact match.
	OpNotEquals operator = "not_eq"
	// OpGreaterThan checks if value is greater than another value. Only works on numbers.
	OpGreaterThan operator = "gt"
	// OpGreaterThanEqual checks if value is greater than or equal to another value. Only works on numbers.
	OpGreaterThanEqual operator = "gte"
	// OpLessThan checks if value is less than another value. Only works on numbers.
	OpLessThan operator = "lt"
	// OpLessThanEqual checks if value is less than or equal to another value. Only works on numbers.
	OpLessThanEqual operator = "lte"
	// OpPrefix only works on strings.
	OpPrefix operator = "prefix"
	// OpExists checks for a non-null value.
	OpExists operator = "exists"
	// OpNotExists checks for a null value.
	OpNotExists operator = "not_exists"
	// OpIn checks if the value exists within an array and between values.
	OpIn operator = "in"
	// OpNotIn checks if the values do not not exist within an array and between values.
	OpNotIn operator = "not_in"
	// OpAny checks if the value has any within the array or between values.
	OpAny operator = "any"
)

// OptFilter is a functional option used to filter the results from
// an API call. Provide a field name to specify what property you
// want to filter with. Provide an operator to specify how you want
// to filter the results using the given field name. Provide a concrete
// value as a string to specify the value of the configured filter. This
// is the only option allowed to have more than one of in a single API call.
// For more information visit https://igdb.github.io/api/references/filters/.
func OptFilter(field string, op operator, val string) OptionFunc {
	return func(o *Options) error {
		if field == "" || val == "" {
			return ErrEmptyField
		}
		s := fmt.Sprintf("filter[%s][%s]", field, string(op))
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
		if page < 1 {
			return ErrOutOfRange
		}
		if o.Values.Get("offset") != "" {
			return ErrExclusiveOption
		}
		if o.Values.Get("scroll") != "" {
			return ErrOptionSet
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
			return ErrEmptyQuery
		}
		o.Values.Set("search", qry)
		return nil
	}
}
