package validation

import (
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

func Validate(opts govalidator.Options) url.Values {
	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) == 0 {
		return nil
	}

	return e
}

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
