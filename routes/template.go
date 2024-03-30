package routes

import "github.com/labstack/echo/v4"

func TemplateHandler() *echo.Echo {

    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index", "")
    })

    e.GET("/login", func(c echo.Context) error {
        return c.Render(200, "login", "")
    })

    e.GET("/signup", func(c echo.Context) error {
        return c.Render(200, "signup", "")
    })

    e.Static("/public", "views/public")
    e.Static("/css", "views/css")
    e.Static("/js", "views/js")

    return e
}
