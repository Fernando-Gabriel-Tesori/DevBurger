package initializers

import "github.com/DevBurger/models"

func SyncDatabase() {
	DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{}, // <- importante
	)
}
