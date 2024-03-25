package controllers

import (
    "net/http"
    "context"

    "github.com/labstack/echo/v4"

    // "postwoman/models"
)

func ValidateCreds(c echo.Context) error {
    
    var res string
    userID := c.Param("id")
    
    err := db.QueryRow(context.Background(), "SELECT * FROM request WHERE user_id=$1", userID).Scan(&res)
    
    if err != nil {
        echo.NewHTTPError(http.StatusNotFound, "")
    }
    
    return c.JSONPretty(http.StatusOK, res, " ")
}

func GetUser(c echo.Context) error {
    
    var res string
    userID := c.Param("id")
    
    err := db.QueryRow(context.Background(), "SELECT * FROM request WHERE user_id=$1", userID).Scan(&res)
    
    if err != nil {
        echo.NewHTTPError(http.StatusNotFound, "")
    }
    
    return c.JSONPretty(http.StatusOK, res, " ")
}

func CreateUser(c echo.Context) error {
    
    var res string
    userID := c.Param("id")
    
    err := db.QueryRow(context.Background(), "SELECT * FROM request WHERE user_id=$1", userID).Scan(&res)
    
    if err != nil {
        echo.NewHTTPError(http.StatusNotFound, "")
    }
    
    return c.JSONPretty(http.StatusOK, res, " ")
}

func UpdateUser(c echo.Context) error {
    
    var res string
    userID := c.Param("id")
    
    err := db.QueryRow(context.Background(), "SELECT * FROM request WHERE user_id=$1", userID).Scan(&res)
    
    if err != nil {
        echo.NewHTTPError(http.StatusNotFound, "")
    }
    
    return c.JSONPretty(http.StatusOK, res, " ")
}

func DeleteUser(c echo.Context) error {
    
    var res string
    userID := c.Param("id")
    
    err := db.QueryRow(context.Background(), "SELECT * FROM request WHERE user_id=$1", userID).Scan(&res)
    
    if err != nil {
        echo.NewHTTPError(http.StatusNotFound, "")
    }
    
    return c.JSONPretty(http.StatusOK, res, " ")
}

func UpdateHistory(c echo.Context) error {
    
    var res string
    userID := c.Param("id")
    
    err := db.QueryRow(context.Background(), "SELECT * FROM request WHERE user_id=$1", userID).Scan(&res)
    
    if err != nil {
        echo.NewHTTPError(http.StatusNotFound, "")
    }
    
    return c.JSONPretty(http.StatusOK, res, " ")
}

func UpdateFavorites(c echo.Context) error {    
    var res string
    userID := c.Param("id")
    
    err := db.QueryRow(context.Background(), "SELECT * FROM request WHERE user_id=$1", userID).Scan(&res)
    
    if err != nil {
        echo.NewHTTPError(http.StatusNotFound, "")
    }
    
    return c.JSONPretty(http.StatusOK, res, " ")
}
