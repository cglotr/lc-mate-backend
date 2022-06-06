package util_test

import (
	"testing"

	"github.com/cglotr/lc-mate-backend/util"
	"github.com/stretchr/testify/assert"
)

func TestGetUsernamesQuery(t *testing.T) {
	s := util.GetUsernamesQuery([]string{"awice", "numb3r5"})
	assert.Equal(t, `"awice","numb3r5"`, s)
}
