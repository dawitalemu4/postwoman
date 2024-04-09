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

    e.GET("/handle/navbar/:page/:token", handlers.RenderNavbar)

    e.GET("/handle/username/:token", handlers.RenderUsername)
    e.GET("/handle/shortcut/:token", handlers.RenderHomeShortcuts)
    e.GET("/handle/request/new", handlers.RenderNewRequest)
    e.GET("/handle/request/response", handlers.RenderRequestResponse)
    // e.GET("/handle/request/history", handlers.RenderHistory)
    // e.GET("/handle/request/favorites", handlers.RenderFavorites)

    e.GET("/handle/login/:token", handlers.RenderLogin)
    e.GET("/handle/signup/:token", handlers.RenderSignup)

    e.GET("/handle/profile/info/:token", handlers.RenderProfileInfo)
    e.GET("/handle/profile/update/:token", handlers.RenderProfileUpdate)
    e.GET("/handle/profile/delete/:deleted", handlers.RenderProfileDelete)

    e.Static("/public", "views/public")
    e.Static("/robots.txt", "views/public/robots.txt")
    e.Static("/favicon.ico", "views/public/favicon.ico")
    e.Static("/css", "views/css")
    e.Static("/js", "views/js")

    return e
}
