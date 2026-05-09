package file

import (
	"fmt"
)

// Conf contains command-line options for unbox
type Conf struct {
	FilePath     string
	OutPath      string
	TemplatePath string
	RenameTags   bool
	ValidateJSON bool
}

// Config - app config
var Config Conf

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

	if Config.ValidateJSON {
		resStr = valIndent(resStr)
	}

	if Config.OutPath != "" {
		outToFile(resStr)
	} else {
		fmt.Println(resStr)
	}
}
