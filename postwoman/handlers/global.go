package handlers

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "postwoman/utils"
    "postwoman/views"
)

func ConfigGlobalHandler() *echo.Echo {

    var env = utils.GetEnv()
    e := echo.New()
    
    e.Renderer = views.RenderTemplate()

    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{env["CLIENT_URL"], env["LOCAL_URL"], env["AUTH0_DOMAIN"]},
    }))

    e.Pre(middleware.RemoveTrailingSlash())

    e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {}))

    return e 
}

var e = ConfigGlobalHandler()
