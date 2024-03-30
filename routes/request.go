package routes

import (
    "github.com/labstack/echo/v4"

    "postwoman/handlers"
)

func RequestHandler() *echo.Echo {

    e.GET("/api/request/all/:userID", handlers.GetAllRequests)
    e.POST("/api/request/new/:userID", handlers.CreateRequest)
    e.DELETE("/api/request/delete/:reqID/:userID", handlers.DeleteRequest)

    return e
}
