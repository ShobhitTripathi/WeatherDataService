package validation

import (
	"WeatherDataService/constant"
	"regexp"
)

func Validate(zipCode string, days int) bool {

	if zipCode == constant.EMPTY_STRING {
		return false
	}

	if !validateZipCode(zipCode) {
		return false
	}
	return true
}

func validateZipCode(zipCode string) bool {
	if zipCode == constant.EMPTY_STRING {
		return false
	}

	zipCodePattern := `^\d{5}(?:-\d{4})?$`
	match, _ := regexp.MatchString(zipCodePattern, zipCode)
	return match
}
