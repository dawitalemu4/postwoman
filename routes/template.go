package routes

import "github.com/labstack/echo/v4"

func TemplateHandler() *echo.Echo {

    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index", map[string]string{"home": "navbar-active", "login": "", "signup": ""})
    })

    e.GET("/login", func(c echo.Context) error {
        return c.Render(200, "login", map[string]string{"home": "", "login": "navbar-active", "signup": ""})
    })

    e.GET("/signup", func(c echo.Context) error {
        return c.Render(200, "signup", map[string]string{"home": "", "login": "", "signup": "navbar-active"})
    })

    e.Static("/public", "views/public")
    e.Static("/robots.txt", "views/public/robots.txt")
    e.Static("/css", "views/css")
    e.Static("/js", "views/js")

    return e
}
