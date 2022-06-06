package util_test

import (
	"testing"

	"github.com/cglotr/lc-mate-backend/util"
	"github.com/stretchr/testify/assert"
)

func TestSpinUpTestDb(t *testing.T) {
	db := util.SpinUpTestDb()
	assert.NotNil(t, db)
}
