package handlers

import (
    "errors"
    "time"
    "strconv"
    "os/exec"
    "encoding/json"
    "bytes"

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
        return c.HTML(200, "<p>$  hello anon! Signup or login to save your favorite requests and organize your request history in your own profiles</p>")
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
        <form id="new-request"
            hx-post="/curl/request"
            hx-target="#request-response"
            hx-swap="innerHTML"
            hx-ext="json-enc"
            hx-on::before-request="dots()"
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
        <div id="request-response"></div>
    `)
}

func RenderHistoryList(c echo.Context) error {

    var historyList []models.Request
    var htmlHistoryList string
    statusColors := map[string]string{"1": "green", "2": "green", "3": "yellow", "4": "red", "5": "orange"}

    email := c.Param("email")

    historyListResponse, _ := exec.Command("curl", "http://localhost:13234/api/request/all/" + email).Output()
    json.NewDecoder(bytes.NewReader(historyListResponse)).Decode(&historyList)

    if len(historyList) == 0 {
        return c.HTML(200, `<br /><p style="margin-left:15px;">$  no history</p>`)
    }

    for i, request := range historyList {

        int64Date, _ := strconv.ParseInt(request.Date, 10, 64)

        if request.Hidden == false {

            request.Date = humanize.Time(time.UnixMilli(int64Date))

            htmlHistoryList += `
                <div class="history-item" tabindex="` + strconv.Itoa(i+1) + `" id="` + strconv.Itoa(request.ID) + `">
                    <div class="history-item-left-container">
                        <p style="color: ` + statusColors[request.Status[0:1]] + `;font-size:18px;">` + request.Status + `</p>
                        <p>` + request.Method + `</p>
                    </div>
                    <div class="history-item-right-container">
                        <p>` + request.Url + `</p>
                        <p>` + request.Date + `</p>
                    </div>
                    <div class="added-favorite">added to favorites</div>
                    <div class="removed-favorite">removed from favorites</div>
                </div>
            `
        }
    }

    return c.HTML(200, htmlHistoryList)
}

func RenderFavoritesList(c echo.Context) error {

    var favoritesList []models.Request
    var htmlFavoritesList string
    statusColors := map[string]string{"1": "green", "2": "green", "3": "yellow", "4": "red", "5": "orange"}

    email := c.Param("email")

    favoritesListResponse, _ := exec.Command("curl", "http://localhost:13234/api/request/favorites/" + email).Output()
    json.NewDecoder(bytes.NewReader(favoritesListResponse)).Decode(&favoritesList)

    if len(favoritesList) == 0 {
        return c.HTML(200, `<br /><p style="margin-left:15px;">$  no favorites</p>`)
    }

    for i, request := range favoritesList {

        int64Date, _ := strconv.ParseInt(request.Date, 10, 64)

        if request.Hidden == false {

            request.Date = humanize.Time(time.UnixMilli(int64Date))

            htmlFavoritesList += `
                <div class="favorites-item" tabindex="` + strconv.Itoa(i+1) + `" id="` + strconv.Itoa(request.ID) + `">
                    <div class="favorites-item-left-container">
                        <p style="color: ` + statusColors[request.Status[0:1]] + `;font-size:18px;">` + request.Status + `</p>
                        <p>` + request.Method + `</p>
                    </div>
                    <div class="favorites-item-right-container">
                        <p>` + request.Url + `</p>
                        <p>` + request.Date + `</p>
                    </div>
                    <div class="added-favorite">added to favorites</div>
                    <div class="removed-favorite">removed from favorites</div>
                </div>
            `
        }
    }
    return c.HTML(200, htmlFavoritesList)
}
