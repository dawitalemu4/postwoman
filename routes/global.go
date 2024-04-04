package routes

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "postwoman/handlers"
)

func ConfigGlobalRoutes() *echo.Echo {

    e := echo.New()
    
    e.Renderer = handlers.RenderTemplate()

    // e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    //     AllowOrigins: []string{env["CLIENT_URL"], env["LOCAL_URL"]},
    // }))

    e.Pre(middleware.RemoveTrailingSlash())

    e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {}))

    return e 
}

var e = ConfigGlobalRoutes()
