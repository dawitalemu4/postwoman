package controllers

import (
    "net/http"
    "context"

    "github.com/labstack/echo/v4"

    // "postwoman/models"
)

func GetAllRequests(c echo.Context) error {
    
    var res string
    userID := c.Param("id")
    
    err := db.QueryRow(context.Background(), "SELECT * FROM request WHERE user_id=$1", userID).Scan(&res)
    
    if err != nil {
        echo.NewHTTPError(http.StatusNotFound, "")
    }
    
    return c.JSONPretty(http.StatusOK, res, " ")
}

func CreateRequest(c echo.Context) error {
    
    var res string
    userID := c.Param("id")
    
    // data := models.Request{}

    err := db.QueryRow(context.Background(), "INSERT INTO request (user_id) VALUES ($1)", userID).Scan(&res)
    
    if err != nil {
        echo.NewHTTPError(http.StatusBadRequest, "")
    }

    return c.JSONPretty(http.StatusOK, res, " ")
}

func DeleteRequest(c echo.Context) error {
    
    var res string
    userID := c.Param("id")
    requestID := c.Param("reqID")
    
    err := db.QueryRow(context.Background(), "DELETE * FROM request WHERE id = $1 AND user_id = $2", requestID, userID).Scan(&res)
    
    if err != nil {
        echo.NewHTTPError(http.StatusInternalServerError, "Error: " + err.Error())
        c.String(http.StatusInternalServerError, "Error: " + err.Error())
    }

    return c.NoContent(http.StatusOK)
}
