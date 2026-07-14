package parse

import (
	"fmt"
)

// SettingsType contains unbox settings for parse
type SettingsType struct {
	FilePath     string
	OutPath      string
	TemplatePath string
	RenameTags   bool
	ValidateJSON bool
}

// Settings contains unbox settings for parse
var Settings SettingsType

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

	if Settings.ValidateJSON {
		resStr = valIndent(resStr)
	}

	if Settings.OutPath != "" {
		outToFile(resStr)
	} else {
		fmt.Println(resStr)
	}
}
