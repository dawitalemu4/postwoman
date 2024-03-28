package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func TemplateHandler() *echo.Echo {

    e.GET("/", func(c echo.Context) error {
        return c.Render(http.StatusOK, "index", "Hello Sandra!")
    })

    e.GET("/login", func(c echo.Context) error {
        return c.Render(http.StatusOK, "login", "log")
    })

    e.GET("/signup", func(c echo.Context) error {
        return c.Render(http.StatusOK, "signup", "signup")
    })

    e.Static("/public", "views/html/public")

    return e
}
