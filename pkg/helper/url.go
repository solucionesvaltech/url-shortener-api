package helper

import (
	"net/url"
	"regexp"
)

// IsValidURL checks if the URL is valid and correctly formed
func IsValidURL(input string) bool {
	u, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	if u.Host == "" {
		return false
	}

	regex := regexp.MustCompile(`^(http|https)://[^\s$.?#].\S*$`)
	return regex.MatchString(input)
}
