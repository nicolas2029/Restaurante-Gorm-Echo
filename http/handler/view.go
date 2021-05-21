package handler

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewRender() *Template {
	route := "../public/views/"
	renderer := &Template{
		templates: template.Must(template.ParseGlob(route + "table/*.html")),
	}
	return renderer
}
