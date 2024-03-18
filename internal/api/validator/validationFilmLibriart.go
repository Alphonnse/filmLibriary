package validator 

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/Alphonnse/filmLibriary/internal/model"
)

func ValidateActorAddRequest(model *model.AddActorInfoRequest) url.Values {
	errs := url.Values{}
	if model.Name == "" || strings.Contains(model.Name, "/") {
		errs.Add("name", "the filed is required")
	}
	if model.Sex == "" {
		errs.Add("sex", "the field is required")
	}
	if !isValidDateFormat(model.Birthday) {
		errs.Add("birthday", "invalid format or blank")
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
	if len(model.Title) < 1 || len(model.Title) > 150 || strings.Contains(model.Title, "/") {
		errs.Add("title", "should be length of 1;150")
	}
	if len(model.Description) < 1 || len(model.Description) > 1000 {
		errs.Add("description", "is not proper long")
	}
	if model.Rate < 1 || model.Rate > 10 {
		errs.Add("rate", "should be in range 1;10")
	}
	if !isValidDateFormat(model.ReleaseDate) {
		errs.Add("releaseDate", "invalid format or blank")
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
	if len(model.Title) < 1 || len(model.Title) > 150 || strings.Contains(model.Title, "/"){
		errs.Add("title", "should be length of 1;150")
	}
	if len(model.Description) < 1 || len(model.Description) > 1000 {
		errs.Add("description", "is not proper long")
	}
	if !isValidDateFormat(model.ReleaseDate) && model.ReleaseDate != "" {
		errs.Add("releaseDate", "invalid format or blank")
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

func ValidateGetFilmsList(urlReq *url.URL) (url.Values, string, string) {
	errs := url.Values{}
	var sortBy string
	var order string
	UrlParts := strings.Split(urlReq.String(), "/")
	if len(UrlParts) > 5 {
		errs.Add("malformed url", "too long. It should look like /film/getList/{sortBy}/{orderBy}")
	}
	if len(UrlParts) > 3 {
		sortBy = UrlParts[3]
		if sortBy != "" && sortBy != "title" && sortBy != "rate" && sortBy != "releaseDate" {
			errs.Add("sortredBy", "Should be blak || title || rate || releaseDate")
		}
	}
	if len(UrlParts) == 5 {
		order = UrlParts[4]
		if order != "" && order != "descending" && order != "ascending" {
			errs.Add("order", "Should be blank || descending || ascending")
		}
	}
	return errs, sortBy, order
}

func ValidateGetFilmsListByFragment(urlReq *url.URL) (url.Values, string, string) {
	errs := url.Values{}
	var fragmentOf string
	var fragment string
	UrlParts := strings.Split(urlReq.String(), "/")
	if len(UrlParts) > 5 {
		errs.Add("malformed url", "too long. It should look like /film/findFilm/{fragmentOf}/{fragment}")
	}
	if len(UrlParts) > 3 {
		fragmentOf = UrlParts[3]
		if fragmentOf != "film_title" && fragmentOf != "actor_name" {
			errs.Add("fragmentOf", "Should be film_title || actor_name ")
		}
	}
	if len(UrlParts) == 5 {
		fragment = UrlParts[4]
		if fragment == "" || strings.Contains(fragment, "/"){
			errs.Add("fragment", "the url part is required")
		}
	}
	return errs, fragmentOf, fragment
}

func isValidDateFormat(birthday string) bool {
	re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	return re.MatchString(birthday)
}
