package file

import (
	"os"
	"text/template"

	"github.com/aceberg/unbox/pkg/check"
)

func insertToTemplate(res, tags string) {

	t, err := template.ParseFiles(Config.TemplatePath)
	check.IfError(err)

	data := map[string]interface{}{
		"Unbox_outbounds": res,
		"Unbox_tags":      tags,
	}

	err = t.Execute(os.Stdout, data)
	check.IfError(err)
}
