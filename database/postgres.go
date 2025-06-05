package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectPostgres() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	if host == "" || user == "" || password == "" || dbName == "" || port == "" {
		log.Fatal("Variáveis de ambiente do banco de dados PostgreSQL não estão definidas corretamente.")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbName, port,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados PostgreSQL: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Erro ao obter instância de conexão com o banco de dados: %v", err)
	}

	// Configurações de pool de conexões
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Conexão com PostgreSQL estabelecida com sucesso.")
}
