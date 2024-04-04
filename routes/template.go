package routes

import (
    "github.com/labstack/echo/v4"

    "postwoman/handlers"
)

func TemplateRoutes() *echo.Echo {

    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "postwoman", map[string]string{"screen": "index"})
    })

    e.GET("/login", func(c echo.Context) error {
        return c.Render(200, "postwoman", map[string]string{"screen": "login"})
    })

    e.GET("/signup", func(c echo.Context) error {
        return c.Render(200, "postwoman", map[string]string{"screen": "signup"})
    })

    e.GET("/profile", func(c echo.Context) error {
        return c.Render(200, "postwoman", map[string]string{"screen": "profile"})
    })

    e.GET("/handle/username/:token", handlers.RenderUsername)
    e.GET("/handle/navbar/:token", handlers.RenderNavbar)
    e.GET("/handle/login/:token", handlers.RenderLogin)
    e.GET("/handle/signup/:token", handlers.RenderSignup)

    e.Static("/public", "views/public")
    e.Static("/robots.txt", "views/public/robots.txt")
    e.Static("/css", "views/css")
    e.Static("/js", "views/js")

    return e
}
