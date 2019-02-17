package render

import (
	"net/http"
	"html/template"
)

func Render(w http.ResponseWriter) error {
	t, err := template.ParseFiles(
		"src/project/WWW/index.html",
	)
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(w, "index", nil)
	return err
}
