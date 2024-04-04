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

    if tokenString == "empty" {
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

    token, err := parseToken(c.Param("token"))

    if err.Error() == emptyError.Error() {
        return c.Render(200, "navbar_profile", `
            <a id="{{ $login_active }}" href="/login">login /</a>
            <a id="{{ $signup_active }}" href="/signup">/ signup</a>
        `)
    }

    if err != nil {
        return c.String(500, "Server Error: " + err.Error())
    }

    return c.Render(200, "navbar_profile", `
        <a id="{{ $profile_active }}" href="/profile">` + token.Username + ` /</a>
        <a href="/logout">/ logout</a>
    `)
}

func RenderUsername(c echo.Context) error {

    token, err := parseToken(c.Param("token"))

    if err.Error() == emptyError.Error() {
        return c.Render(200, "username", "<p>$  Hello anon! Signup or login to save your request history</p>")
    }

    if err != nil {
        return c.String(500, "Server Error: " + err.Error())
    }

    return c.Render(200, "username", "<p>$  Hello " + token.Username + "!</p>")
}

func RenderLogin(c echo.Context) error {

    token, err := parseToken(c.Param("token"))

    if err != nil {
        return c.String(500, "Server Error: " + err.Error())
    }

    return c.Render(200, "login_response", "<p>$  Welcome back " + token.Username + "!</p>")
}

func RenderSignup(c echo.Context) error {

    token, err := parseToken(c.Param("token"))

    if err != nil {
        return c.String(500, "Server Error: " + err.Error())
    }

    return c.Render(200, "signup_response", map[string]interface{}{"username": "<p>$  " + token.Username + "</p>", "email": "<p>$  email: " + token.Email + "</p>"})
}
