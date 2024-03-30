package handlers

import (
    "context"
    "encoding/json"
    "time"

    "github.com/labstack/echo/v4"
    "github.com/golang-jwt/jwt/v5"

    "postwoman/models"
)

func CreateJWT(c echo.Context) error {

    var data models.User

    json.NewDecoder(c.Request().Body).Decode(&data)

    dataWithExpiration := &models.User{
        data.ID, data.Username, data.Email, data.Password, data.History, data.Favorites, data.Date, data.Token, data.Deleted,
        jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 336)),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, dataWithExpiration)

    res, err := token.SignedString([]byte("secret"))

    if err != nil {
        return c.JSONPretty(500, errorJSON("Server Error", err.Error()), " ")
    }

    return c.JSONPretty(200, res, " ")
}

func GetUser(c echo.Context) error {

    var res int
    var data models.User

    json.NewDecoder(c.Request().Body).Decode(&data)

    if data.Validated(data) {

        err := db.QueryRow(context.Background(), `SELECT id FROM "user" WHERE id = $1 AND email = $2 AND password = $3`, data.ID, data.Email, data.Password).Scan(&res)

        if res != data.ID {
            return c.JSONPretty(401, errorJSON("User Error", "No users found from this id and cred combo"), " ")
        }

        if err != nil {
            return c.JSONPretty(500, errorJSON("Server Error", err.Error()), " ")
        }
    } else {
        return c.JSONPretty(404, errorJSON("User Error", "Invalid data"), " ")
    }

    return CreateJWT(c)
}

func CreateUser(c echo.Context) error {

    var res string
    var data models.User

    json.NewDecoder(c.Request().Body).Decode(&data)

    if data.Validated(data) {

        err := db.QueryRow(context.Background(), `INSERT INTO "user" (username, email, password, history, favorites, date, token, deleted) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`,
            data.Username, data.Email, data.Password, data.History, data.Favorites, data.Date, data.Token, data.Deleted).Scan(&res)

        if err != nil && err.Error() != "no rows in result set" {
            return c.JSONPretty(500, errorJSON("Server Error", err.Error()), " ")
        }
    } else {
        return c.JSONPretty(404, errorJSON("User Error", "Invalid data"), " ")
    }

    return c.JSONPretty(200, res, " ")
}

func UpdateUser(c echo.Context) error {

    var res string
    var data models.User

    json.NewDecoder(c.Request().Body).Decode(&data)

    if data.Validated(data) {

        err := db.QueryRow(context.Background(), `UPDATE "user" SET username = $1, email = $2, password = $3, history = $4, favorites = $5, date = $6, token = $7, deleted = $8 WHERE id = $9 RETURNING id`,
            data.Username, data.Email, data.Password, data.History, data.Favorites, data.Date, data.Token, data.Deleted, data.ID).Scan(&res)

        if err != nil && err.Error() != "no rows in result set" {
            return c.JSONPretty(500, errorJSON("Server Error", err.Error()), " ")
        }
    } else {
        return c.JSONPretty(404, errorJSON("User Error", "Invalid data"), " ")
    }

    return c.JSONPretty(200, res, " ")
}

func DeleteUser(c echo.Context) error {

    var res bool
    var data models.User

    json.NewDecoder(c.Request().Body).Decode(&data)

    if data.Validated(data) {

        err := db.QueryRow(context.Background(), `UPDATE "user" SET deleted = $1 WHERE id = $2 AND email = $3 AND password = $4 RETURNING deleted`, data.Deleted, data.ID, data.Email, data.Password).Scan(&res)

        if err != nil && err.Error() != "no rows in result set" {
            return c.JSONPretty(500, errorJSON("Server Error", err.Error()), " ")
        }
    } else {
        return c.JSONPretty(404, errorJSON("User Error", "Invalid data"), " ")
    }

    return c.JSONPretty(200, res, " ")
}

func UpdateHistory(c echo.Context) error {

    var res []int
    var data models.User
    remove := c.QueryParam("remove")

    json.NewDecoder(c.Request().Body).Decode(&data)

    if data.Validated(data) {

        if remove == "true" {
            HideRequest(c)
        }

        err := db.QueryRow(context.Background(), `UPDATE "user" SET history = $1 WHERE id = $2 AND email = $3 AND password = $4 RETURNING history`, data.History, data.ID, data.Email, data.Password).Scan(&res)

        if err != nil && err.Error() != "no rows in result set" {
            return c.JSONPretty(500, errorJSON("Server Error", err.Error()), " ")
        }
    } else {
        return c.JSONPretty(404, errorJSON("User Error", "Invalid data"), " ")
    }

    return c.JSONPretty(200, res, " ")
}

func UpdateFavorites(c echo.Context) error {

    var res []int
    var data models.User

    json.NewDecoder(c.Request().Body).Decode(&data)

    if data.Validated(data) {

        err := db.QueryRow(context.Background(), `UPDATE "user" SET favorites = $1 WHERE id = $2 AND email = $3 AND password = $4 RETURNING favorites`, data.Favorites, data.ID, data.Email, data.Password).Scan(&res)

        if err != nil && err.Error() != "no rows in result set" {
            return c.JSONPretty(500, errorJSON("Server Error", err.Error()), " ")
        }
    } else {
        return c.JSONPretty(404, errorJSON("User Error", "Invalid data"), " ")
    }

    return c.JSONPretty(200, res, " ")
}
