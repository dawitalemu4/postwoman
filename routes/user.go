package routes

import (
    "github.com/labstack/echo/v4"
    
    "postwoman/handlers"
)

func UserHandler() *echo.Echo {

    e.POST("/api/auth/user", handlers.GetUser)
    e.POST("/api/user/new", handlers.CreateUser)
    e.PUT("/api/user/update/:id", handlers.UpdateUser)
    e.DELETE("/api/user/delete/:id", handlers.DeleteUser)

    e.PATCH("/api/user/history/:userID/:reqID", handlers.UpdateHistory)
    e.PATCH("/api/user/favorites/:userID/:reqID", handlers.UpdateFavorites)

    return e
}
