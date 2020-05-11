package testing

import (
	"errors"
)

//GetData returns test data
func GetData(id string) (string, error) {
	if id == "fail" {
		return "", errors.New("Failure")
	}

	return "yay", nil
}
