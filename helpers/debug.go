package helpers

import (
	"os"

	"github.com/davecgh/go-spew/spew"
)

// Pre exit running project.
// Param: interface{}
// Param: ...interface{}
func Pre(x interface{}, y ...interface{}) {
	spew.Dump(x)
	os.Exit(1)
}
