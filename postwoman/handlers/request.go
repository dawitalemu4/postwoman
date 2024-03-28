package handlers

import (
    "github.com/labstack/echo/v4"
    
    "postwoman/controllers"
)

func RequestHandler() *echo.Echo {

    e.GET("/api/request/all/:userID", controllers.GetAllRequests)
    e.POST("/api/request/new/:userID", controllers.CreateRequest)
    e.DELETE("/api/request/delete/:reqID/:userID", controllers.DeleteRequest)

    return e
}
