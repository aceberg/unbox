package parse

import (
	"fmt"

	"github.com/aceberg/unbox/internal/share"
)

var (
	result []string
	tags   []string
	i      uint16
)

// Start converting
func Start() {

	parseFile()

	if i < 2 {
		return
	}

	resStr := insertToTemplate()

	if share.ParseSettings.ValidateJSON {
		resStr = valIndent(resStr)
	}

	if share.ParseSettings.OutPath != "" {
		outToFile(resStr)
	} else {
		fmt.Println(resStr)
	}
}
