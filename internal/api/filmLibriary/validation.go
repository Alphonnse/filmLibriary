package filmlibriary

import (
	"net/url"
	"strconv"

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
	if model.UUID.String() == "" || model.UUID.String() == "00000000-0000-0000-0000-000000000000" {
		errs.Add("UUID", "wrong UUID")
	}
	return errs
}

func ValidateAddFilmInfoRequest(model *model.AddFilmInfoRequest) url.Values {
	errs := url.Values{}
	if len(model.Title) < 1 || len(model.Title) > 150 {
		errs.Add("title", "should be length of 1;150")
	}
	if len(model.Description) < 1 || len(model.Description) > 1000 {
		errs.Add("description", "is not proper long")
	}
	if model.Rate < 0 || model.Rate > 10 {
		errs.Add("rate", "should be in range 1;10")
	}
	if model.ReleaseDate == "" {
		errs.Add("releaseDate", "the field is required")
	}
	if len(model.Actors) < 1 {
		errs.Add("actors", "the field is required")
	} else {
		for i, actor := range model.Actors {
			actorErrs := ValidateActorAddRequest(&actor)
			if len(actorErrs) > 0 {
				for field, err := range actorErrs {
					errs.Add("actors["+strconv.Itoa(i)+"]["+field+"]", err[0])
				}
			}
		}
	}
	return errs
}

func ValidateChangeFilmInfoRequest(model *model.ChangeFilmInfoRequest) url.Values {
	errs := url.Values{}
	if model.UUID.String() == "" || model.UUID.String() == "00000000-0000-0000-0000-000000000000" {
		errs.Add("UUID", "wrong UUID")
	}
	if len(model.Title) < 1 || len(model.Title) > 150 {
		errs.Add("title", "should be length of 1;150")
	}
	if len(model.Description) < 1 || len(model.Description) > 1000 {
		errs.Add("description", "is not proper long")
	}
	if model.Rate < 0 || model.Rate > 10 {
		errs.Add("rate", "should be in range 1;10")
	}
	for i, actor := range model.Actors {
		actorErrs := ValidateChangeActorRequest(&actor)
		if len(actorErrs) > 0 {
			for field, err := range actorErrs {
				errs.Add("actors["+strconv.Itoa(i)+"]["+field+"]", err[0])
			}
		}
	}
	return errs
}


func ValidateGetFilmsList(model *model.GetFilmsListRequest) url.Values {
	errs := url.Values{}
	if model.SortBy != ""  && model.SortBy != "title" && model.SortBy != "rate" && model.SortBy != "releaseDate"{
		errs.Add("sortredBy", "Should be blank || title || rate || releaseDate")
	}
	if model.Order != "" && model.Order != "descending" && model.Order != "ascending" {
		errs.Add("order", "Should be blank || descending || ascending")
	}
	return errs
}


func ValidateGetFilmsListByFragment(model *model.GetFilmsListByFragmentRequest) url.Values {
	errs := url.Values{}
	if model.FragmentOf != "" && model.FragmentOf != "film title" && model.FragmentOf != "actor name" {
		errs.Add("fragmentOf", "Should be blank || film title || actor name ")
	}
	if model.Fragment == "" {
		errs.Add("fragment", "the field is required")
	}
	return errs
}
