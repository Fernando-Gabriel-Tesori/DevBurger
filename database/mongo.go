package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("Erro ao conectar no MongoDB: %v", err)
	}

	// Verifica se a conexão está ativa
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Erro ao realizar ping no MongoDB: %v", err)
	}

	MongoClient = client
	log.Println("✅ Conectado ao MongoDB com sucesso.")
}

// Helper para acessar o banco (padrão)
func GetMongoDatabase() *mongo.Database {
	return MongoClient.Database(os.Getenv("MONGO_DB"))
}
