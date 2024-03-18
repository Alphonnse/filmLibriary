package validator

import (
	"net/url"

	"github.com/Alphonnse/filmLibriary/internal/model"
)

func ValidateUserAddRequest(model *model.Register) url.Values {
	errs := url.Values{}
	if model.Name == "" {
		errs.Add("name", "the filed is required")
	}
	if model.Password == "" {
		errs.Add("password", "the filed is required")
	}
	return errs
}

func ValidateUserGetRequest(model *model.Login) url.Values {
	errs := url.Values{}
	if model.Name == "" {
		errs.Add("name", "the filed is required")
	}
	if model.Password == "" {
		errs.Add("password", "the filed is required")
	}
	return errs
}

