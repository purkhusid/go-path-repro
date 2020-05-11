package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceShouldReturnData(t *testing.T) {

	_, err := GetData("id")

	assert.NoError(t, err)
}

func TestServiceShouldReturnErrorIfFail(t *testing.T) {

	_, err := GetData("fail")

	assert.EqualError(t, err, "Failure")
}
