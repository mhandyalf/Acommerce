package models

type User struct {
	ID            uint   `gorm:"primaryKey"`
	Username      string `gorm:"unique"`
	Password      string
	DepositAmount float64
}

type Product struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Stock int
	Price float64
}

type Transaction struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	ProductID   uint
	Quantity    int
	TotalAmount float64
}
