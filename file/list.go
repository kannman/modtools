package file

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/kannman/modtools/output"
)

// ParseDepsListJSON parses the dependency file into a data structure.
func ParseDepsListJSON(raw string) []ModItem {
	decoder := json.NewDecoder(strings.NewReader(raw))
	data := make([]ModItem, 0, 10)

	for {
		var dep ModItem
		err := decoder.Decode(&dep)

		if err != nil {
			if err == io.EOF {
				break
			}
			output.OnError(err, "Error decoding dependency json")
		}

		data = append(data, dep)
	}

	return data
}

type ModItem struct {
	Path    string
	Version string
	Replace *ModReplace
	Time    string
	Dir     string
	GoMod   string
}

type ModReplace struct {
	Path    string
	Version string
	Time    string
	Dir     string
	GoMod   string
}
