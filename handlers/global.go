package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

    "postwoman/utils"
	"postwoman/views"
)

var env = utils.GetEnv()

func ConfigGlobalHandler() *echo.Echo {

    e := echo.New()
    
    e.Renderer = views.RenderTemplate()

    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{env["CLIENT_URL"], env["LOCAL_URL"]},
    }))

    e.Pre(middleware.RemoveTrailingSlash())

    e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {}))

    return e 
}

var e = ConfigGlobalHandler()
