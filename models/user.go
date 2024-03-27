package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
    ID           int       `json:"id"`
    Username     string    `json:"username"`
    Email        string    `json:"email"`
    Password     string    `json:"password"`
    History      []int     `json:"history"`
    Favorites    []int     `json:"favorites"`
    Date         string    `json:"date"`
    Deleted      bool      `json:"deleted"`
	jwt.RegisteredClaims
}

func (user User) Validated(data User) bool {

    if data.Username != "" &&
        data.Email != "" &&
        data.Password != "" &&
        len(data.History) != 0 || len(data.History) == 0 &&
        len(data.Favorites) != 0 || len(data.Favorites) == 0 &&
        data.Date != "" &&
        data.Deleted == false {

        return true
    }

    return false
}
