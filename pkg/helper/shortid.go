package helper

import "github.com/teris-io/shortid"

// GenerateShortID generate a unique ID using an external library
func GenerateShortID() (string, error) {
	return shortid.Generate()
}
