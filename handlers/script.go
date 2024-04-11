package handlers

import (
    "os/exec"
    "strconv"
    "time"
    "encoding/json"
    "regexp"

    "github.com/labstack/echo/v4"

    "postwoman/models"
)

func buildCommand(request models.Request) []string {

    command := []string{"curl", "-v"}

    if request.Method != "" {
        command = append(command, "-X", request.Method)
    }

    if request.Headers != "" {
        command = append(command, "-H", `'` + request.Headers + `'`)
    }

    if request.Origin != "" {
        command = append(command, "-H", `'Origin: ` + request.Origin + `'`)
    }

    if request.Body != "" {
        command = append(command, "-d", `'` + request.Body + `'`)
    }

    command = append(command, request.Url)
    return command
}

func ExecuteCurlRequest(c echo.Context) error {

    var request models.Request

    json.NewDecoder(c.Request().Body).Decode(&request)

    command := buildCommand(request)
    response, err := exec.Command(command[0], command[1:]...).Output()

    request.Status = "200"
    request.Date = strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
    request.Hidden = false

    stringRequest, _ := json.Marshal(request)

    exec.Command("curl", "-X", "POST", "-d", string(stringRequest), "http://localhost:13234/api/request/new/" + request.User_email).Output()

    errorResponseRegex := regexp.MustCompile(`<title>(?s).*Error.*<\/title>`)
    if errorMatch := errorResponseRegex.FindStringSubmatch(string(response)); errorMatch != nil || err != nil {

        preTagRegex := regexp.MustCompile(`<pre>(?s).*?<\/pre>`)
        match := preTagRegex.FindStringSubmatch(string(response))

        if match == nil {

            if err.Error() == "exit status 6" {
                return c.HTML(200, "<p>$  error: " + err.Error() + ", probably an invalid url</p>")
            }

            return c.HTML(200, "<p>$  error: " + err.Error() + "<br />response: " + string(response) + "</p>")

        } else {
            return c.HTML(200, "$  error: " + match[0])
        }
    }

    return c.HTML(200, `<textarea class="response-textarea" readonly>` + string(response) + `</textarea>`)
}
