package handlers

import (
    "html/template"
    "io"

    "github.com/labstack/echo/v4"
)

type Template struct {
    templates *template.Template
}

func (template *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return template.templates.ExecuteTemplate(w, name, data)
}

func RenderTemplate() *Template {

    return &Template{
        templates: template.Must(template.ParseGlob("views/html/*.html")),
    }
}
