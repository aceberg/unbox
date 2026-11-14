package parse

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/aceberg/unbox/internal/check"
	"github.com/aceberg/unbox/internal/share"
)

func getLinksFromFile() (string, bool) {

	f, err := os.ReadFile(share.ParseSettings.FilePath)
	if check.IfError(err) {
		log.Println("ERROR: Input file error. Exiting...")
		return "", false
	}

	return string(f), true
}

func insertToTemplate() string {

	res := strings.Join(result, ",\n")

	if share.ParseSettings.TemplatePath == "" {
		return res
	}

	t, err := template.ParseFiles(share.ParseSettings.TemplatePath)
	if check.IfError(err) {
		log.Println("ERROR: Template file error")
		return res
	}

	data := map[string]interface{}{
		"Unbox_outbounds": res,
		"Unbox_tags":      strings.Join(tags, ","),
	}

	var buf bytes.Buffer

	err = t.Execute(&buf, data)
	check.IfError(err)

	return buf.String()
}

func outToFile(out string) {

	err := os.WriteFile(share.ParseSettings.OutPath, []byte(out), 0644)
	if check.IfError(err) {
		log.Println("ERROR: Output file error")
	}
}

func valIndent(raw string) string {
	var out bytes.Buffer

	if !json.Valid([]byte(raw)) {
		log.Println("ERROR: JSON is not valid!")
		return raw
	}

	err := json.Indent(&out, []byte(raw), "", "  ")
	check.IfError(err)

	return out.String()
}
