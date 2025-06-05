package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	MongoURI   string
}

var (
	config *Config
	once   sync.Once
)

// LoadConfig carrega as variáveis de ambiente do arquivo .env e do ambiente de sistema
func LoadConfig() error {
	var err error
	once.Do(func() {
		// Tenta carregar .env (ignora erro se não existir)
		err = godotenv.Load()
		if err != nil {
			log.Println("Aviso: arquivo .env não encontrado, carregando variáveis do ambiente")
		}

		config = &Config{
			Port:       getEnv("PORT", "8080"),
			DBHost:     getEnv("DB_HOST", "localhost"),
			DBPort:     getEnv("DB_PORT", "5432"),
			DBUser:     getEnv("DB_USER", "postgres"),
			DBPassword: getEnv("DB_PASSWORD", ""),
			DBName:     getEnv("DB_NAME", "devburger"),
			JWTSecret:  getEnv("JWT_SECRET", "your-secret-key"),
			MongoURI:   getEnv("MONGO_URI", "mongodb://localhost:27017"),
		}
	})
	return err
}

// Get retorna a instância da configuração carregada
func Get() *Config {
	if config == nil {
		log.Fatal("Configuração não carregada. Chame LoadConfig() antes.")
	}
	return config
}

// getEnv busca uma variável de ambiente com fallback padrão
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
