package utils

import (
    "github.com/labstack/echo/v4"
    "github.com/joho/godotenv"
)

func GetEnv() map[string]string {

    e := echo.New()
    var env map[string]string

    godotenv.Load(".env")
    env, err := godotenv.Read()

    if err != nil {
        e.Logger.Fatal("Error loading .env file")
    }

    return env
}

var env = GetEnv()
