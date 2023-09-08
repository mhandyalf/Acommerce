package main

import (
	"acommerce/database"
	"acommerce/handlers"
	"acommerce/middleware"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	authhandler := handlers.NewAuth(database.InitDB())
	e.POST("/register", authhandler.Register)
	e.POST("/login", authhandler.Login)
	e.GET("/products", authhandler.GetProducts, middleware.JWTAuth)
	e.GET("/transactions", authhandler.GetTransactions, middleware.JWTAuth)
	e.GET("/stores", authhandler.GetStores)
	e.GET("/weather", authhandler.GetWeatherByCityName)
	e.GET("/store/:id", authhandler.GetStoreByID)
	e.POST("/buy", authhandler.BuyProduct, middleware.JWTAuth)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))

}
