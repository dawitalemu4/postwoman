package routes

import (
    "github.com/labstack/echo/v4"

    "postwoman/handlers"
)

func ConfigGlobalRoutes() *echo.Echo {

    e := echo.New()
    
    e.Renderer = handlers.RenderTemplate()

    // e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    //     AllowOrigins: []string{env["CLIENT_URL"], env["LOCAL_URL"]},
    // }))

    return e 
}

var e = ConfigGlobalRoutes()
