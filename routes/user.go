package routes

import (
    "github.com/labstack/echo/v4"
    
    "postwoman/handlers"
)

func UserHandler() *echo.Echo {

    e.POST("/api/user/auth", handlers.GetUser)
    e.POST("/api/user/new", handlers.CreateUser)
    e.PUT("/api/user/update", handlers.UpdateUser)
    e.DELETE("/api/user/delete", handlers.DeleteUser)

    e.PATCH("/api/user/history/:reqID", handlers.UpdateHistory)
    e.PATCH("/api/user/favorites", handlers.UpdateFavorites)

    return e
}
