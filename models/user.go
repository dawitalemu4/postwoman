package models

type User struct {
    id           int
    username     string
    email        string
    password     string
    history      []int
    favorites    []int
    deleted      bool
}

func (user User) validate(data User) bool {
    
    if data.id != 0 ||
        data.username != "" ||
        data.email != "" ||
        data.password != "" ||
        data.deleted == false {

        return true
    } else {
        return false
    }
}
