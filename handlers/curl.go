package handlers

import (
    "encoding/json"
    "os/exec"
    "strconv"
    "time"
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

    if err != nil && err.Error() == "exit status 6" {
        return c.HTML(200, "<p>$  error: " + err.Error() + ", probably an invalid url given</p>")
    }

    headersCmd := append(command[:1], append([]string{"-LI"}, command[1:]...)...)
    headers, err := exec.Command(headersCmd[0], headersCmd[1:]...).Output()

    if err != nil {
        return c.HTML(200, "<p>$  error: " + err.Error() + ", probably an invalid header or body given</p>")
    }

    statusRegex := regexp.MustCompile(`HTTP\/\d\.\d\s(\d{3})`)
    statusMatch := statusRegex.FindStringSubmatch(string(headers))
    splicedStatus := statusMatch[0][len(statusMatch[0])-3:]

    request.Status = splicedStatus 
    request.Date = strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
    request.Hidden = false

    stringRequest, _ := json.Marshal(request)

    exec.Command("curl", "-X", "POST", "-d", string(stringRequest), "http://localhost:" + env["GO_PORT"] + "/api/request/new/" + request.User_email).Output()

    errorResponseRegex := regexp.MustCompile(`<title>(?s).*Error.*<\/title>`)
    if errorMatch := errorResponseRegex.FindStringSubmatch(string(response)); errorMatch != nil {

        preTagRegex := regexp.MustCompile(`<pre>(?s).*?<\/pre>`)
        preTagMatch := preTagRegex.FindStringSubmatch(string(response))

        if preTagMatch == nil {
            return c.HTML(200, "$  status: " + splicedStatus + "response: " + string(response))
        } else {
            return c.HTML(200, "$  error: " + preTagMatch[0] + "<br /><br />status: " + splicedStatus)
        }
    }

    return c.HTML(200, `
        $  status: ` + splicedStatus + `
        <br /><br />
        <textarea id="response-textarea" readonly>` + string(response) + `&#013;</textarea>
    `)
}
