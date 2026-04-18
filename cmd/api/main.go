package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CasimiroDev/volunteer-service/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.HealthCheck)

	port := os.Getenv("PORT")
	fmt.Printf("Servidor Gestão de Voluntários iniciado na porta %s...\n", port)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("Erro crítico no servidor: %v", err)
	}
}
