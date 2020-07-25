package validation

import (
	"net/url"

	"github.com/thedevsaddam/govalidator"
)

// Validate : Validate struct by options.
func Validate(opts govalidator.Options) url.Values {
	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) == 0 {
		return nil
	}

	return e
}

// MergeErrors : Takes a list of `url.Values` merge them together when multiple validations occuerr.
func MergeErrors(values ...url.Values) url.Values {
	out := url.Values{}

	for _, v := range values {
		for k, vv := range v {
			for _, vvv := range vv {
				out.Add(k, vvv)
			}
		}
	}

	if len(out) == 0 {
		return nil
	}

	return out
}
