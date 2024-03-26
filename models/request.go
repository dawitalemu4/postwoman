package models

type Request struct {
    ID         int       `json:"id"`
    User_id    int       `json:"user_id"`
    Url        string    `json:"url"`
    Method     string    `json:"method"`
    Headers    string    `json:"headers"`
    Body       string    `json:"body"`
    Status     string    `json:"status"`
    Date       string    `json:"date"`
    Deleted    bool      `json:"deleted"`
}

func (request Request) Validated(data Request) bool {

    method_options := map[string]bool{
        "GET": true, "POST": true, "UPDATE": true, "PUT": true, "DELETE": true,
    }

    if data.User_id != 0 &&
        data.Url != "" &&
        method_options[data.Method] &&
        data.Headers != "" || data.Headers == "" &&
        data.Body != "" || data.Body == "" &&
        data.Status != "" &&
        data.Date != "" &&
        data.Deleted == false {

        return true 
    }

    return false 
}
