package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
    Username     string    `json:"username"`
    Email        string    `json:"email"`
    Password     string    `json:"password"`
    Favorites    []int     `json:"favorites"`
    Date         string    `json:"date"`
    Deleted      bool      `json:"deleted"`
    OldPw        string    `json:"oldPassword"`
    jwt.RegisteredClaims
}

func (user User) Validated(data User) bool {

    if data.Username != "" &&
        data.Email != "" &&
        data.Password != "" &&
        data.Deleted == false {

        return true
    }

    return false
}
