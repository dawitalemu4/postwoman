package handlers

import (
    "os/exec"
    "encoding/json"
    "time"
    "strconv"

    "github.com/labstack/echo/v4"

    "postwoman/models"
)

func checkArgs(request models.Request) (string, error) {
    
    if request.Method != "" && request.Headers != "" && request.Origin != "" && request.Body != "" {
        response, err := exec.Command("curl", "-v", "-X", request.Method, "-H", request.Headers, "-H", `' Origin: ` + request.Origin + `'`, "-d", `'` + request.Body + `'`, request.Url).Output()
        return string(response), err
    } else if request.Method != "" && request.Headers != "" && request.Origin != "" {
        response, err := exec.Command("curl", "-v", "-X", request.Method, "-H", request.Headers, "-H", `' Origin: ` + request.Origin + `'`, request.Url).Output()
        return string(response), err
    } else if request.Method != "" && request.Headers != "" {
        response, err := exec.Command("curl", "-v", "-X", request.Method, "-H", `'` + request.Headers + `'`, request.Url).Output()
        return string(response), err
    } else if request.Method != "" {
        response, err := exec.Command("curl", "-v", "-X", request.Method, request.Url).Output()
        return string(response), err
    }
    
    response, err := exec.Command("curl", "-v", request.Url).Output()
    return string(response), err
}

func ExecuteCurlRequest(c echo.Context) error {

    var request models.Request

    json.NewDecoder(c.Request().Body).Decode(&request)

    response, err := checkArgs(request)

    if err != nil {
        return c.HTML(200, "<p>$  Server Error: " + err.Error() + "</p>")
    }

    request.Status = "200"
    request.Date = strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
    request.Hidden = false

    stringRequest, err := json.Marshal(request)

    exec.Command("curl", "-X", "POST", "-d", string(stringRequest), "http://localhost:1323/api/request/new/" + request.User_email).Output()
    exec.Command("curl", "http://localhost:1323/handle/request/new/" + request.User_email).Output()

    return c.HTML(200, "<p>$  " + response + "</p>")
}
