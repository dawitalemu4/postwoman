package handlers

import (
    "github.com/labstack/echo/v4"
    
    "postwoman/controllers"
)

func UserHandler() *echo.Echo {

    e.POST("/api/auth/user", controllers.GetUser)
    e.POST("/api/user/new", controllers.CreateUser)
    e.PUT("/api/user/update/:id", controllers.UpdateUser)
    e.DELETE("/api/user/delete/:id", controllers.DeleteUser)

    e.PATCH("/api/user/history/:userID/:reqID", controllers.UpdateHistory)
    e.PATCH("/api/user/favorites/:userID/:reqID", controllers.UpdateFavorites)

    return e
}
