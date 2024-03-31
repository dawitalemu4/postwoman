package models

type Request struct {
    ID            int       `json:"id"`
    User_email    string    `json:"user_email"`
    Url           string    `json:"url"`
    Method        string    `json:"method"`
    Origin        string    `json:"origin"`
    Headers       string    `json:"headers"`
    Body          string    `json:"body"`
    Status        string    `json:"status"`
    Date          string    `json:"date"`
    Hidden        bool      `json:"hidden"`
}

func (request Request) Validated(data Request) bool {

    method_options := map[string]bool{
        "GET": true, "POST": true, "UPDATE": true, "PUT": true, "DELETE": true,
    }

    if data.Url != "" &&
        method_options[data.Method] &&
        data.Origin != "" &&
        data.Headers != "" || data.Headers == "" &&
        data.Body != "" || data.Body == "" &&
        data.Status != "" &&
        data.Date != "" &&
        data.Hidden == false {

        return true 
    }

    return false 
}
