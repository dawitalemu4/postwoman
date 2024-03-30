package utils 

import (
    "context"

    "github.com/labstack/echo/v4"
    "github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {
    
    e := echo.New()

    db, err := pgx.Connect(context.Background(), env["DB_URL"])

    if err != nil {
        e.Logger.Fatal("Unable to connect to database: %v\n", err)
    }

    return db
}
