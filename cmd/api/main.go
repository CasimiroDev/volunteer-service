package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CasimiroDev/volunteer-service/internal/handler"
	"github.com/CasimiroDev/volunteer-service/internal/repository"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Arquivo .env não encontrado.")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("Erro crítico: DATABASE_URL não está configurada.")
	}

	fmt.Println("Conectando ao banco de dados...")
	db, err := repository.NewDatabaseConnection(dbURL)
	if err != nil {
		log.Fatalf("Erro crítico: falha ao iniciar o banco de dados: %v", err)
	}
	defer db.Close()
	fmt.Println("Conexão com o PostgreSQL estabelecida com sucesso!")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.HealthCheck)

	fmt.Printf("Servidor Gestão de Voluntários iniciado na porta %s...\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Erro crítico no servidor: %v", err)
	}
}
