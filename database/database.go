package database

import (
	"DevBurger/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=devburger port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao conectar no banco: ", err)
	}

	db.AutoMigrate(&models.Order{}, &models.Product{}) // AutoMigrando modelos

	DB = db
	fmt.Println("Banco conectado e modelos migrados com sucesso!")
}
