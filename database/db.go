package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func InitDB() *gorm.DB {
	dsn := "user=postgres password=$zLP8ZZJzzByL.$ dbname=postgres host=db.yamzgmeiezqqkgteurma.supabase.co port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
