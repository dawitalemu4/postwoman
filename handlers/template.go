package handlers

import (
    "errors"
    "time"
    "strconv"

    "github.com/golang-jwt/jwt/v5"
    "github.com/labstack/echo/v4"
    "github.com/dustin/go-humanize"

    "postwoman/models"
)

var emptyError error = blankTokenError()

func blankTokenError() error {
    return errors.New("token is blank")
}

func parseToken(tokenString string) (*models.User, error) {

    var claims *models.User

    if tokenString == "null" {
        return nil, blankTokenError()
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
        return c.HTML(500, "<p>$  Server Error: " + err.Error() + "</p>")
    }

    return c.HTML(200, `
        <a id="` + pages["profile"] + `" href="/profile">` + token.Username + ` /</a>
        <a href="/" onclick="localStorage.removeItem("auth");">/ logout</a>
    `)
}

func RenderUsername(c echo.Context) error {

    token, err := parseToken(c.Param("token"))

    if err != nil && err.Error() == emptyError.Error() {
        return c.HTML(200, "<p>$  hello anon! Signup or login to save your request history</p>")
    }

    if err != nil {
        return c.HTML(500, "<p>$  Server Error: " + err.Error() + "</p>")
    }

    return c.HTML(200, "<p>$  hello " + token.Username + "!</p>")
}

func RenderLogin(c echo.Context) error {

    token, err := parseToken(c.Param("token"))

    if err != nil && err.Error() == emptyError.Error() {
        return c.HTML(200, "<p>$  incorrect credentials</p>")
    }

    if err != nil {
        return c.HTML(500, "<p>$  Server Error: " + err.Error() + "</p>")
    }

    return c.HTML(200, "<p>$  welcome back " + token.Username + "!</p>")
}

func RenderSignup(c echo.Context) error {

    token, err := parseToken(c.Param("token"))

    if err != nil && err.Error() == emptyError.Error() {
        return c.HTML(200, "<p>$  invalid input</p>")
    }

    if err != nil {
        return c.HTML(500, "<p>$  Server Error: " + err.Error() + "</p>")
    }

    return c.HTML(200, "<p>$  account created! username: " + token.Username + ", email: " + token.Email + "</p>")
}

func RenderProfileInfo(c echo.Context) error {

    token, err := parseToken(c.Param("token"))
    date, _ := strconv.ParseInt(token.Date, 10, 64)
    userSince := humanize.Time(time.UnixMilli(date));

    if err != nil && err.Error() == emptyError.Error() {
        return c.HTML(200, "<p>$  invalid token</p>")
    }

    if err != nil {
        return c.HTML(500, "<p>$  Server Error: " + err.Error() + "</p>")
    }

    return c.HTML(200, "<p>$  username: " + token.Username + ", email: " + token.Email + ", user since " + userSince + "</p>")
}

func RenderProfileUpdate(c echo.Context) error {

    token, err := parseToken(c.Param("token"))

    if err != nil && err.Error() == emptyError.Error() {
        return c.HTML(200, "<p>$  invalid input</p>")
    }

    if err != nil {
        return c.HTML(500, "<p>$  Server Error: " + err.Error() + "</p>")
    }

    return c.HTML(200, "<p>$  account updated! username: " + token.Username + ", email: " + token.Email + ", password: " + token.Password + "</p>")
}

func RenderProfileDelete(c echo.Context) error {

    deleted := c.Param("deleted")

    if deleted != "true" {
        return c.HTML(200, "<p>$  invalid token, try to log back in</p>")
    }

    return c.HTML(200, "<p>$  deleting account</p>")
}

func RenderHomeShortcuts(c echo.Context) error {

    token, err := parseToken(c.Param("token"))

    if token == nil && err.Error() == emptyError.Error() {
        return c.HTML(200, `
            <div><kbd>ctrl</kbd> + <kbd>alt</kbd> + <kbd>l</kbd> - login page</div>
            <div><kbd>ctrl</kbd> + <kbd>alt</kbd> + <kbd>s</kbd> - signup page</div>
        `)
    }

    return c.HTML(200, `
        <div><kbd>ctrl</kbd> + <kbd>alt</kbd> + <kbd>p</kbd> - profile page</div>
        <div><kbd>ctrl</kbd> + <kbd>alt</kbd> + <kbd>l</kbd> - logout</div>
    `)
}

func RenderNewRequest(c echo.Context) error {

    email := c.Param("email")

    return c.HTML(200, `
        <form class="new-request"
            hx-post="/curl/request"
            hx-target=".request-response:last-child"
            hx-ext="json-enc"
        >
            $  curl -X <select name="method" autofocus required>
                <option value="GET">GET</option>
                <option value="POST">POST</option>
                <option value="PUT">PUT</option>
                <option value="PATCH">PATCH</option>
                <option value="DELETE">DELETE</option>
            </select> \ <br />
            -H '<input name="headers" type="text" placeholder="headers" />' \ <br />
            -H '<input name="origin" type="text" placeholder="origin" />' \ <br />
            -d '<textarea name="body" type="text" placeholder="body"></textarea>' \ <br />
            <input name="url" type="text" placeholder="url" required />
            <input name="user_email" value="` + email + `" hidden />
            <input type="submit" hidden />
        </form>
        <div class="request-response"></div>
    `)
}
