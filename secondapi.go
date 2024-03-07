package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Usuario struct {
	Name     string `json:"name"`
	Humor    string `json:"humor"`
	Faltas   int    `json:"faltas"`
	Filepath string `json:"filepath"`
}

var usuarios []Usuario

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// Novo middleware para tratar CORS preflight
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func updateStatesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var novosUsuarios []Usuario
		err := json.NewDecoder(r.Body).Decode(&novosUsuarios)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		usuarios = append(usuarios, novosUsuarios...)
		fmt.Printf("Usuários recebidos: %+v\n", novosUsuarios)
		fmt.Fprintf(w, "Dados recebidos com sucesso\n")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func sendUserDataHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "URL inválida. Esperado /sendUserData/{índice}.", http.StatusBadRequest)
		return
	}
	indexStr := parts[len(parts)-1]
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		http.Error(w, "Índice inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if index < 0 || index >= len(usuarios) {
		http.Error(w, "Índice fora do intervalo.", http.StatusNotFound)
		return
	}

	usuario := usuarios[index]
	jsonData, err := json.Marshal(usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/updateStates", corsMiddleware(updateStatesHandler))
	http.HandleFunc("/sendUserData/", corsMiddleware(sendUserDataHandler)) // Nota: Adicionado "/" no fim para capturar tudo após /sendUserData

	fmt.Println("Servidor iniciado na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
