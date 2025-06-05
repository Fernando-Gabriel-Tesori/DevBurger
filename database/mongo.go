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
var MongoDatabase *mongo.Database

func ConnectMongo() {
	mongoURI := os.Getenv("MONGO_URI")
	mongoDBName := os.Getenv("MONGO_DB_NAME")

	if mongoURI == "" || mongoDBName == "" {
		log.Fatal("Variáveis de ambiente MONGO_URI ou MONGO_DB_NAME não estão definidas.")
	}

	// Define um contexto com timeout para a conexão
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define as opções de conexão
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}

	// Testa a conexão
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Erro ao realizar ping no MongoDB: %v", err)
	}

	MongoClient = client
	MongoDatabase = client.Database(mongoDBName)

	log.Println("Conexão com o MongoDB estabelecida com sucesso.")
}
