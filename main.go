package main

/*
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type StateUpdate struct {
	Name  string `json:"name"`
	Humor string `json:"humor"`
}

type HumorUpdate struct {
	Name  string `json:"Name"`
	Humor string `json:"Humor"`
}

var humorUpdates []HumorUpdate

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func applyCorsAndOptions(w http.ResponseWriter, r *http.Request) bool {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return false
	}
	return true
}

func receiveHumorUpdates(w http.ResponseWriter, r *http.Request) {
	if !applyCorsAndOptions(w, r) {
		return
	}
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "POST" {
		var updates []HumorUpdate
		err := json.NewDecoder(r.Body).Decode(&updates)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		humorUpdates = append(humorUpdates, updates...)
		fmt.Fprintf(w, "Atualizações de humor recebidas e armazenadas com sucesso.")
	} else {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}

func fetchUserDataHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var usersData []UserData = []UserData{
		{"Usuario 1", 2, "Bom", "packages/graph/regressao_20.png"},
		{"Usuario 2", 0, "Ótimo", "packages/graph/regressao_20.png"},
		// Adicione mais usuários conforme necessário
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(usersData); err != nil {
		log.Printf("Erro ao codificar a resposta JSON: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
	}
}

func updateStatesHandler(w http.ResponseWriter, r *http.Request) {
	if !applyCorsAndOptions(w, r) {
		return
	}
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	var stateUpdates []StateUpdate
	if err := json.NewDecoder(r.Body).Decode(&stateUpdates); err != nil {
		log.Printf("Erro ao decodificar a requisição JSON: %v", err)
		http.Error(w, "Erro no formato de entrada", http.StatusBadRequest)
		return
	}

	log.Printf("Atualizações de Humor Recebidas: %+v", stateUpdates)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Atualizações de humor recebidas com sucesso"})
}

type UserData struct {
	Name       string `json:"name"`
	Faltas     int    `json:"faltas"`
	Humor      string `json:"humor"`
	GraphImage string `json:"graphImage"`
}

func getHumorUpdatesHandler(w http.ResponseWriter, r *http.Request) {
	if !applyCorsAndOptions(w, r) {
		return
	}
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(humorUpdates); err != nil {
		log.Printf("Erro ao codificar a resposta JSON: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/fetchUserData", fetchUserDataHandler)
	http.HandleFunc("/updateStates", updateStatesHandler)
	http.HandleFunc("/updateHumor", receiveHumorUpdates)
	http.HandleFunc("/getHumorUpdates", getHumorUpdatesHandler)

	fs := http.FileServer(http.Dir("packages/graph"))
	http.Handle("/packages/graph/", http.StripPrefix("/packages/graph/", fs))

	port := ":8080"
	log.Printf("Servidor iniciando na porta %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
*/
