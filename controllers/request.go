package controllers

import (
    "context"
    "encoding/json"
    "strconv"

    "github.com/labstack/echo/v4"
    "github.com/stephenafamo/scan"
    "github.com/stephenafamo/scan/pgxscan"

    "postwoman/models"
)

func GetAllRequests(c echo.Context) error {

    userID := c.Param("userID")

    res, err := pgxscan.All(context.Background(), db, scan.StructMapper[models.Request](), "SELECT * FROM request WHERE user_id = $1", userID)

    if len(res) == 0 {
        return c.JSONPretty(404, errorJSON("User Error", "No requests found from this user ID"), " ")
    }

    if err != nil {
        return c.JSONPretty(500, errorJSON("Server Error", err.Error()), " ")
    }

    return c.JSONPretty(200, res, " ")
}

func CreateRequest(c echo.Context) error {

    var data models.Request
    userID := c.Param("userID")
    intUserID, _ := strconv.Atoi(userID)

    json.NewDecoder(c.Request().Body).Decode(&data)

    if data.Validated(data) && intUserID == data.User_id {

        err := db.QueryRow(context.Background(), "INSERT INTO request (user_id, url, method, headers, body, status, date, deleted) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
            userID, data.Url, data.Method, data.Headers, data.Body, data.Status, data.Date, data.Deleted).Scan()

        if err != nil && err.Error() != "no rows in result set" {
            return c.JSONPretty(500, errorJSON("Server Error", err.Error()), " ")
        }
    } else {
        return c.JSONPretty(404, errorJSON("User Error", "Invalid data"), " ")
    }

    return c.NoContent(200)
}

func DeleteRequest(c echo.Context) error {

    var res string
    requestID := c.Param("reqID")
    userID := c.Param("userID")

    err := db.QueryRow(context.Background(), "DELETE FROM request WHERE id = $1 AND user_id = $2 RETURNING $2", requestID, userID).Scan(&res)

    if res != userID {
        return c.JSONPretty(404, errorJSON("User Error", "No requests found made with the IDs provided"), " ")
    }

    if err != nil && err.Error() != "no rows in result set" {
        return c.JSONPretty(500, errorJSON("Server Error", err.Error()), " ")
    }

    return c.NoContent(200)
}
