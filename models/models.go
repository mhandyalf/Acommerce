package models

type User struct {
	ID            uint   `gorm:"primaryKey"`
	Username      string `gorm:"unique"`
	Password      string
	DepositAmount float64
}

type Product struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Stock   int
	Price   float64
	StoreID uint `gorm:"foreignKey:StoreID"`
}

type Transaction struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	ProductID   uint
	Quantity    int
	TotalAmount float64
	StoreID     uint `gorm:"foreignKey:StoreID"`
}

type Store struct {
	StoreID     int64 `gorm:"primaryKey"`
	NamaStore   string
	Alamat      string
	Longitude   float64
	Latitude    float64
	Rating      float64
	WeatherData WeatherData `gorm:"embedded"` // Menyimpan WeatherData sebagai bagian dari Store
}

type WeatherData struct {
	CloudPct    int     `json:"cloud_pct"`
	Temp        int     `json:"temp"`
	FeelsLike   int     `json:"feels_like"`
	Humidity    int     `json:"humidity"`
	MinTemp     int     `json:"min_temp"`
	MaxTemp     int     `json:"max_temp"`
	WindSpeed   float64 `json:"wind_speed"`
	WindDegrees int     `json:"wind_degrees"`
	Sunrise     int     `json:"sunrise"`
	Sunset      int     `json:"sunset"`
}
