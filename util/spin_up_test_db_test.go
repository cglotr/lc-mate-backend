package util_test

import (
	"testing"

	"github.com/cglotr/lc-mate-backend/util"
	"github.com/hooligram/kifu"
)

func TestSpinUpTestDbError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			kifu.Info("Recovered")
		}
	}()
	util.SpinUpTestDb(".")
}
