package filmlibriary

import (
	"net/url"

	"github.com/Alphonnse/filmLibriary/internal/model"
)

func ValidateActorAddRequest(model *model.AddActorInfoRequest) url.Values {
	errs := url.Values{}
	if model.Name == "" {
		errs.Add("name", "the filed is required")
	}
	if model.Sex == "" {
		errs.Add("sex", "the field is required")
	}
	if model.Birthday == "" {
		errs.Add("birthday", "the field is required")
	}
	return errs
}

func ValidateChangeActorRequest(model *model.ChangeActorInfoRequest) url.Values {
	errs := url.Values{}
	if model.UUID.String() == "" {
		errs.Add("UUID", "wrong UUID")
	}
	return errs
}

