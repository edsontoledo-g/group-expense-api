package templates

import (
	"html/template"
	"io/fs"
)

func Load(filesystems ...fs.FS) *template.Template {
	t := template.New("")
	for _, filesystem := range filesystems {
		parsed, err := t.ParseFS(filesystem, "*.html")
		if err != nil {
			panic(err)
		}
		t = parsed
	}
	return t
}
