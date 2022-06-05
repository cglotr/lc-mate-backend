package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidUsernames(t *testing.T) {
	invalids := []string{
		"support",
		"jobs",
		"bugbounty",
		"student",
		"terms",
		"privacy",
		"region",
		"explore",
		"contest",
		"discuss",
		"interview",
		"store",
		"profile",
	}
	for _, username := range invalids {
		assert.True(t, IsInvalidUsername(username))
	}

	// making sure there are no duplicates
	assert.Equal(t, len(invalidUsernames), len(invalids))
}
