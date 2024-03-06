package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type UserData struct {
	Name       string `json:"name"`
	Faltas     int    `json:"faltas"`
	Humor      string `json:"humor"`
	GraphImage string `json:"graphImage"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*") // ATENÇÃO: usar '*' em produção pode ser inseguro
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func fetchUserDataHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Recebido uma requisição em /fetchUserData")

	// Habilita CORS para esta resposta específica
	enableCors(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	usersData := []UserData{
		{"Usuario 1", 2, "Bom", "assets/graphs/regressao_10.png"},
		{"Usuario 2", 0, "Ótimo", "assets/graphs/regressao_20.png"},
		// Adicione mais usuários conforme necessário
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(usersData); err != nil {
		// Log the error and return a server error response
		log.Printf("Erro ao codificar a resposta JSON: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/fetchUserData", fetchUserDataHandler)
	port := ":8080"
	log.Printf("Servidor iniciando na porta %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
