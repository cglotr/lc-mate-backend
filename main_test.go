package main

import (
	"testing"

	"github.com/hooligram/kifu"
)

func TestMain(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			kifu.Info("%v", r)
		}
	}()
	main()
}
