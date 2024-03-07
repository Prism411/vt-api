package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Usuario struct {
	Nome     string `json:"name"`
	Humor    string `json:"humor"`
	Faltas   int    `json:"faltas"`
	Filepath string `json:"filepath"`
}

var blocos [][]Usuario

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// Novo middleware para aplicar CORS
func corsMiddleware(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		handler.ServeHTTP(w, r)
	}
}

func updateStatesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	var usuariosRecebidos []Usuario
	err := json.NewDecoder(r.Body).Decode(&usuariosRecebidos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, usuario := range usuariosRecebidos {
		if len(blocos) == 0 || len(blocos[len(blocos)-1]) >= 10 {
			blocos = append(blocos, []Usuario{})
		}
		blocos[len(blocos)-1] = append(blocos[len(blocos)-1], usuario)
	}

	fmt.Fprintf(w, "Dados recebidos com sucesso!")
}

func listarUsuariosHandler(w http.ResponseWriter, r *http.Request) {
	blocoQuery := r.URL.Query().Get("bloco")
	bloco, err := strconv.Atoi(blocoQuery)
	if err != nil || bloco < 1 || bloco > len(blocos) {
		http.Error(w, "Bloco inválido", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(blocos[bloco-1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func disableCacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Adicionando cabeçalhos para desabilitar cache
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
		w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
		w.Header().Set("Expires", "0")                                         // Proxies.
		next.ServeHTTP(w, r)
	})
}

func main() {
	fs := http.FileServer(http.Dir("./packages/graph"))
	http.Handle("/packages/graph/", corsMiddleware(disableCacheMiddleware(http.StripPrefix("/packages/graph/", fs))))

	http.HandleFunc("/updateStates", corsMiddleware(http.HandlerFunc(updateStatesHandler)))
	http.HandleFunc("/usuarios", corsMiddleware(http.HandlerFunc(listarUsuariosHandler)))

	fmt.Println("Servidor iniciado na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
