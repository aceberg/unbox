package file

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"github.com/aceberg/unbox/pkg/check"
)

func getLinksFromFile() (string, bool) {

	f, err := os.ReadFile(Config.FilePath)
	if check.IfError(err) {
		log.Println("File error. Exiting...")
		return "", false
	}

	return string(f), true
}

func insertToTemplate(res, tags string) string {

	t, err := template.ParseFiles(Config.TemplatePath)
	check.IfError(err)

	data := map[string]interface{}{
		"Unbox_outbounds": res,
		"Unbox_tags":      tags,
	}

	var buf bytes.Buffer

	err = t.Execute(&buf, data)
	check.IfError(err)

	return buf.String()
}

func outToFile(out string) {

	err := os.WriteFile(Config.OutPath, []byte(out), 0644)
	check.IfError(err)
}
