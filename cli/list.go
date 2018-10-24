package cli

import (
	"path/filepath"

	"github.com/kannman/modtools/output"
	"github.com/mattn/go-zglob"
)

func BuildModVendorList(copyPat []string, modDir string) []string {
	var vendorList []string

	for _, pat := range copyPat {
		matches, err := zglob.Glob(filepath.Join(modDir, pat))
		output.OnError(err, "glob match failure, path: "+filepath.Join(modDir, pat))

		for _, m := range matches {
			vendorList = append(vendorList, m)
		}
	}

	return vendorList
}
