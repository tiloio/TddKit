package files

import (
	"flag"

	"test-framework/internal/error"

	"github.com/yargevad/filepathx"
)

var dirToSearch = flag.String("path", ".", "Path")
var globPattern = flag.String("glob", "/**/*.test.[tj]s", "Glob")

func Search() (*[]string, error.Common) {
	path := *dirToSearch + *globPattern

	files, err := filepathx.Glob(path)
	if err != nil {
		return nil, error.NewCommon(err, "Could not find path '"+path+"'")
	}

	return &files, nil
}
