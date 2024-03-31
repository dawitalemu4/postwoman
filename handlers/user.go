package handlers

import (
    "context"
    "encoding/json"
    "time"

    "github.com/labstack/echo/v4"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"

    "postwoman/models"
)

func authUser(data models.User) map[string]interface{} {

    var password string

    err := db.QueryRow(context.Background(), `SELECT password FROM "user" WHERE email = $1`, data.Email).Scan(&password)

    if err != nil {
        return map[string]interface{}{"status": 500, "res": errorJSON("Server Error", err.Error())}
    }

    if bcrypt.CompareHashAndPassword([]byte(password), []byte(data.Password)) != nil {
        return map[string]interface{}{"status": 401, "res": errorJSON("User Error", "No users found from this email and password")}
    }

    return map[string]interface{}{"status": 200, "res": true}
}

func createJWT(data models.User) map[string]interface{} {

    dataWithExpiration := &models.User{
        data.Username, data.Email, data.Password, data.History, data.Favorites, data.Date, data.Deleted,
        jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 504)),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, dataWithExpiration)

    res, err := token.SignedString([]byte("secret"))

    if err != nil {
        return map[string]interface{}{"status": 500, "res": errorJSON("Server Error", err.Error())}
    }

    return map[string]interface{}{"status": 200, "res": res}
}

func GetUser(c echo.Context) error {

    var data models.User

    json.NewDecoder(c.Request().Body).Decode(&data)

    if data.Validated(data) != true {
        return c.JSONPretty(404, errorJSON("User Error", "Invalid data"), " ")
    }

    authenticated := authUser(data)

    if authenticated["res"] != true {
        return c.JSONPretty(authenticated["status"].(int), authenticated["res"], " ")
    }

    return c.JSONPretty(200, createJWT(data), " ") 
}

func CreateUser(c echo.Context) error {

    var data models.User

    json.NewDecoder(c.Request().Body).Decode(&data)

    if data.Validated(data) {

        hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
        
        err := db.QueryRow(context.Background(), `INSERT INTO "user" (username, email, password, history, favorites, date, deleted) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
            data.Username, data.Email, hashedPassword, data.History, data.Favorites, data.Date, data.Deleted).Scan()

        if err != nil && err.Error() != "no rows in result set" {
            return c.JSONPretty(500, errorJSON("Server Error", err.Error()), " ")
        }
    } else {
        return c.JSONPretty(404, errorJSON("User Error", "Invalid data"), " ")
    }

    return c.JSONPretty(200, createJWT(data), " ") 
}

func UpdateUser(c echo.Context) error {

    var data models.User

    json.NewDecoder(c.Request().Body).Decode(&data)

    if data.Validated(data) {

        authenticated := authUser(data)

        if authenticated["res"] != true {
            return c.JSONPretty(authenticated["status"].(int), authenticated["res"], " ")
        }

        hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
        
        err := db.QueryRow(context.Background(), `UPDATE "user" SET username = $1, email = $2, password = $3, history = $4, favorites = $5, date = $6, deleted = $7 WHERE email = $8`,
            data.Username, data.Email, hashedPassword, data.History, data.Favorites, data.Date, data.Deleted, data.Email).Scan()

        if err != nil && err.Error() != "no rows in result set" {
            return c.JSONPretty(500, errorJSON("Server Error", err.Error()), " ")
        }
    } else {
        return c.JSONPretty(404, errorJSON("User Error", "Invalid data"), " ")
    }

    return c.JSONPretty(200, createJWT(data), " ") 
}

func DeleteUser(c echo.Context) error {

    var res bool
    var data models.User

    json.NewDecoder(c.Request().Body).Decode(&data)

    if data.Validated(data) {

        authenticated := authUser(data)

        if authenticated["res"] != true {
            return c.JSONPretty(authenticated["status"].(int), authenticated["res"], " ")
        }

        err := db.QueryRow(context.Background(), `UPDATE "user" SET deleted = $1 WHERE email = $3 RETURNING deleted`, data.Deleted, data.Email).Scan(&res)

        if err != nil && err.Error() != "no rows in result set" {
            return c.JSONPretty(500, errorJSON("Server Error", err.Error()), " ")
        }
    } else {
        return c.JSONPretty(404, errorJSON("User Error", "Invalid data"), " ")
    }

    return c.JSONPretty(200, res, " ")
}

func UpdateHistory(c echo.Context) error {

    var data models.User
    remove := c.QueryParam("remove")

    json.NewDecoder(c.Request().Body).Decode(&data)

    if data.Validated(data) {

        if remove == "true" {
            
            hidden := HideRequest(c, data.Email)

            if hidden["res"] != "successful" {
                return c.JSONPretty(hidden["status"].(int), hidden["res"], " ")
            }
        }

        authenticated := authUser(data)

        if authenticated["res"] != true {
            return c.JSONPretty(authenticated["status"].(int), authenticated["res"], " ")
        }

        err := db.QueryRow(context.Background(), `UPDATE "user" SET history = $1 WHERE email = $2`, data.History, data.Email).Scan()

        if err != nil && err.Error() != "no rows in result set" {
            return c.JSONPretty(500, errorJSON("Server Error", err.Error()), " ")
        }
    } else {
        return c.JSONPretty(404, errorJSON("User Error", "Invalid data"), " ")
    }

    return c.JSONPretty(200, createJWT(data), " ") 
}

func UpdateFavorites(c echo.Context) error {

    var data models.User

    json.NewDecoder(c.Request().Body).Decode(&data)

    if data.Validated(data) {

        authenticated := authUser(data)

        if authenticated["res"] != true {
            return c.JSONPretty(authenticated["status"].(int), authenticated["res"], " ")
        }

        err := db.QueryRow(context.Background(), `UPDATE "user" SET favorites = $1 WHERE email = $2`, data.Favorites, data.Email).Scan()

        if err != nil && err.Error() != "no rows in result set" {
            return c.JSONPretty(500, errorJSON("Server Error", err.Error()), " ")
        }
    } else {
        return c.JSONPretty(404, errorJSON("User Error", "Invalid data"), " ")
    }

    return c.JSONPretty(200, createJWT(data), " ") 
}
