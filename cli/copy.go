package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/kannman/modtools/output"
)

func CopyFile(src, dst string) int64 {
	sourceFileStat, err := os.Stat(src)
	output.OnError(err, "stat "+src)

	if !sourceFileStat.Mode().IsRegular() {
		output.OnError(err, src+" is not a regular file")
	}

	source, err := os.Open(src)
	output.OnError(err, "open "+src)
	defer source.Close()

	destination, err := os.Create(dst)
	output.OnError(err, "create "+dst)
	defer destination.Close()

	bts, err := io.Copy(destination, source)
	output.OnError(err, fmt.Sprintf("copy src: '%s' dst: '%s'", src, dst))

	return bts
}
