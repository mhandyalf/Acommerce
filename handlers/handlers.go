package handlers

import (
	"acommerce/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Auth struct {
	db *gorm.DB
}

func NewAuth(db *gorm.DB) *Auth {
	return &Auth{db: db}
}

func (a *Auth) Register(e echo.Context) error {
	user := new(models.User)
	if err := e.Bind(user); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	user.Password = string(hashedPassword)
	if err := a.db.Create(user).Error; err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusCreated, user)

}

func (a *Auth) Login(e echo.Context) error {
	user := new(models.User)
	if err := e.Bind(user); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	var foundUser models.User
	if err := a.db.Where("username = ?", user.Username).First(&foundUser).Error; err != nil {
		return e.JSON(http.StatusUnauthorized, map[string]string{
			"message": err.Error(),
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		return e.JSON(http.StatusUnauthorized, map[string]string{
			"message": err.Error(),
		})
	}

	// Buat token JWT
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = foundUser.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token berlaku selama 1 hari

	// Sign token dengan secret key
	tokenString, err := token.SignedString([]byte("secret-key")) // Gantilah "secret-key" dengan kunci rahasia yang lebih kuat

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal membuat token JWT",
		})
	}

	// Kirim token sebagai respons
	return e.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}

func (a *Auth) GetProducts(e echo.Context) error {
	var products []models.Product
	if err := a.db.Find(&products).Error; err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, products)
}

func (a *Auth) GetTransactions(e echo.Context) error {
	var transactions []models.Transaction
	if err := a.db.Find(&transactions).Error; err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, transactions)
}

func (a *Auth) GetStores(e echo.Context) error {
	var stores []models.Store
	if err := a.db.Table("public.store").Find(&stores).Error; err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, stores)
}

func (a *Auth) GetWeatherByCityName(e echo.Context) error {
	city := e.QueryParam("city") // Mengambil nama kota dari query parameter

	weatherDataStr, err := GetWeatherRapidAPI(city)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch weather data"})
	}

	var weatherData models.WeatherData
	if err := json.Unmarshal([]byte(weatherDataStr), &weatherData); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to unmarshal weather data"})
	}

	return e.JSON(http.StatusOK, weatherData)
}

func (a *Auth) GetStoreByID(e echo.Context) error {
	storeID := e.Param("id") // Ambil ID toko dari URL parameter

	var store models.Store
	if err := a.db.Table("public.store").First(&store, storeID).Error; err != nil {
		return e.JSON(http.StatusNotFound, map[string]string{
			"message": "Store not found",
		})
	}

	// Mengambil data cuaca berdasarkan kota toko
	weatherDataStr, err := GetWeatherRapidAPI(store.Alamat)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch weather data for store",
		})
	}

	// Unmarshal JSON response menjadi objek WeatherData
	var weatherData models.WeatherData
	err = json.Unmarshal([]byte(weatherDataStr), &weatherData)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to parse weather data",
		})
	}

	// Menyimpan data cuaca ke dalam toko
	store.WeatherData = weatherData

	// Kirim response

	return e.JSON(http.StatusOK, store)
}
