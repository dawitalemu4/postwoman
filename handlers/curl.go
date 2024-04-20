package handlers

import (
    "strings"
    "bytes"
    "encoding/json"
    "os/exec"
    "strconv"
    "time"
    "regexp"

    "github.com/labstack/echo/v4"

    "postwoman/models"
)

func buildCommand(request models.Request) []string {

    shell := []string{"bash", "-c"}
    command := []string{"curl", "-L", "-v", "--http1.1"}

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
    shell = append(shell, strings.Join(command, " "))
    return shell 
}

func ExecuteCurlRequest(c echo.Context) error {

    var request models.Request
    var response, headers bytes.Buffer

    json.NewDecoder(c.Request().Body).Decode(&request)

    command := buildCommand(request)
    curlRequest := exec.Command(command[0], command[1:]...)

    curlRequest.Stdout = &response
    curlRequest.Stderr = &headers

    err := curlRequest.Run()

    if err != nil && err.Error() == "exit status 6" {
        return c.HTML(200, "<p>$  error: " + err.Error() + ", probably an invalid url given</p>")
    } else if err != nil {
        return c.HTML(200, "<p>$  error: " + err.Error() + ", probably an invalid header or body given</p>")
    }

    statusRegex := regexp.MustCompile(`HTTP\/\d\.\d\s(\d{3})`)
    statusMatch := statusRegex.FindAllStringSubmatch(headers.String(), -1)
    splicedStatus := statusMatch[len(statusMatch) - 1][1]

    request.Status = splicedStatus 
    request.Date = strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
    request.Hidden = false

    stringRequest, _ := json.Marshal(request)

    exec.Command("curl", "-X", "POST", "-d", string(stringRequest), "http://localhost:" + env["GO_PORT"] + "/api/request/new/" + request.User_email).Output()

    errorResponseRegex := regexp.MustCompile(`<title>(?s).*Error.*<\/title>`)

    if errorMatch := errorResponseRegex.FindStringSubmatch(response.String()); errorMatch != nil {

        preTagRegex := regexp.MustCompile(`<pre>(?s).*?<\/pre>`)
        preTagMatch := preTagRegex.FindStringSubmatch(response.String())

        if preTagMatch == nil {
            return c.HTML(200, "$  status: " + splicedStatus + "response: " + response.String())
        } else {
            return c.HTML(200, "$  error: " + preTagMatch[0] + "<br /><br />status: " + splicedStatus)
        }
    }

    return c.HTML(200, `
        $  status: ` + splicedStatus + `
        <br /><br />
        <textarea id="response-textarea" readonly>` + response.String() + `&#013;</textarea>
    `)
}
