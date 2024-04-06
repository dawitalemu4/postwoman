package handlers

import (
    "errors"

    "github.com/golang-jwt/jwt/v5"
    "github.com/labstack/echo/v4"

    "postwoman/models"
)

var emptyError error = blankTokenError()

func blankTokenError() error {
    return errors.New("token is blank")
}

func parseToken(tokenString string) (*models.User, error) {

    var claims *models.User

    if tokenString == "null" {
        return claims, blankTokenError()
    }

    token, err := jwt.ParseWithClaims(tokenString, &models.User{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(env["JWT_SIGNATURE"]), nil
    })

    if claims, ok := token.Claims.(*models.User); ok && token.Valid {
        return claims, err 
    }

    return claims, err
}

func RenderNavbar(c echo.Context) error {

    pages := map[string]string{"login": "", "signup": "", "profile": ""}

    token, err := parseToken(c.Param("token"))
    page := c.Param("page")

    pages[page] = "navbar-active"

    if err != nil && err.Error() == emptyError.Error() {
        return c.HTML(200, `
            <a id="` + pages["login"] + `" href="/login">login /</a>
            <a id="` + pages["signup"] + `" href="/signup">/ signup</a>
        `)
    }

    if err != nil {
        return c.HTML(500, "$  Server Error: " + err.Error())
    }

    return c.HTML(200, `
        <a id="` + pages["profile"] + `" href="/profile">` + token.Username + ` /</a>
        <a href="/" onclick="localStorage.clear();">/ logout</a>
    `)
}

func RenderUsername(c echo.Context) error {

    token, err := parseToken(c.Param("token"))

    if err != nil && err.Error() == emptyError.Error() {
        return c.HTML(200, "<p>$  Hello anon! Signup or login to save your request history</p>")
    }

    if err != nil {
        return c.HTML(500, "$  Server Error: " + err.Error())
    }

    return c.HTML(200, "<p>$  hello " + token.Username + "!</p>")
}

func RenderLogin(c echo.Context) error {

    token, err := parseToken(c.Param("token"))

    if err != nil && err.Error() == emptyError.Error() {
        return c.HTML(200, "<p>$  Incorrect Credentials</p>")
    }

    if err != nil {
        return c.HTML(500, "$  Server Error: " + err.Error())
    }

    return c.HTML(200, "<p>$  welcome back " + token.Username + "!</p>")
}

func RenderSignup(c echo.Context) error {

    token, err := parseToken(c.Param("token"))

    if err != nil && err.Error() == emptyError.Error() {
        return c.HTML(200, "<p>$  Invalid Input</p>")
    }

    if err != nil {
        return c.HTML(500, "$  Server Error: " + err.Error())
    }

    return c.HTML(200, "<p>$  username: " + token.Username + ", email: " + token.Email + "</p>")
}
