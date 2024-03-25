package models

type Request struct {
    id         int
    user_id    int
    url        string
    method     string
    headers    string
    body       string
    status     string
    date       string
    deleted    bool
}

func (request Request) validate(data Request) bool {
    
    method_options := map[string]bool{
        "GET": true, "POST": true, "UPDATE": true, "PUT": true, "DELETE": true,
    }

    if data.id != 0 ||
        data.user_id != 0 ||
        data.url != "" ||
        method_options[data.method] ||
        data.method != "" ||
        data.status != "" ||
        data.date != "" ||
        data.deleted == false {
        
        return true 
    } else {
        return false 
    }
}
