package main

import (
	"acommerce/database"
	"acommerce/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	authhandler := handlers.NewAuth(database.InitDB())
	e.POST("/register", authhandler.Register)

	e.Logger.Fatal(e.Start(":8080"))

}
