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

func args(vs ...interface{}) []interface{} { return vs }

func RenderTemplate() *Template {

    return &Template{
        templates: template.Must(template.New("").Funcs(template.FuncMap{"args": args}).ParseGlob("views/html/*.html")),
    }
}
