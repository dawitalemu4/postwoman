package routes

import (
    "github.com/labstack/echo/v4"

    "postwoman/handlers"
)

func RequestRoutes() *echo.Echo {

    e.GET("/api/request/all/:email", handlers.GetAllRequests)
    e.POST("/api/request/new/:email", handlers.CreateRequest)
    e.DELETE("/api/request/delete/:email/:reqID", handlers.DeleteRequest)

    return e
}
