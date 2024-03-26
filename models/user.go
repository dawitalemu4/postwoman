package models

type User struct {
    ID           int
    Username     string    `json:"username"`
    Email        string    `json:"email"`
    Password     string    `json:"password"`
    History      []int     `json:"history"`
    Favorites    []int     `json:"favorites"`
    Deleted      bool      `json:"deleted"`
}

func (user User) validate(data User) bool {

    if data.Username != "" &&
        data.Email != "" &&
        data.Password != "" &&
        len(data.History) != 0 || len(data.History) == 0 &&
        len(data.Favorites) != 0 || len(data.Favorites) == 0 &&
        data.Deleted == false {

        return true
    }

    return false
}
