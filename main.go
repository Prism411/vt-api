package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Usuario define a estrutura dos dados recebidos.
type Usuario struct {
	Nome     string `json:"name"`     // Nome do usuário
	Humor    string `json:"humor"`    // Estado de humor do usuário
	Faltas   int    `json:"faltas"`   // Número de faltas
	Filepath string `json:"filepath"` // Caminho do arquivo
}

// Lista para armazenar os dados recebidos em blocos de até 10 usuários.
var blocos [][]Usuario

func updateStatesHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica se o método é POST.
	if r.Method != "POST" {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// Decodifica o corpo da requisição para um slice de Usuario.
	var usuariosRecebidos []Usuario
	err := json.NewDecoder(r.Body).Decode(&usuariosRecebidos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Adiciona os usuários recebidos em blocos de 10.
	for _, usuario := range usuariosRecebidos {
		if len(blocos) == 0 || len(blocos[len(blocos)-1]) >= 10 {
			// Cria um novo bloco se necessário.
			blocos = append(blocos, []Usuario{})
		}
		blocos[len(blocos)-1] = append(blocos[len(blocos)-1], usuario)
	}

	// Resposta de sucesso.
	fmt.Fprintf(w, "Dados recebidos com sucesso!")
}

func listarUsuariosHandler(w http.ResponseWriter, r *http.Request) {
	// Obtém o número do bloco da query string, começando de 1.
	blocoQuery := r.URL.Query().Get("bloco")
	bloco, err := strconv.Atoi(blocoQuery)
	if err != nil || bloco < 1 || bloco > len(blocos) {
		http.Error(w, "Bloco inválido", http.StatusBadRequest)
		return
	}

	// Configura o Content-Type como application/json.
	w.Header().Set("Content-Type", "application/json")

	// Codifica o bloco especificado de usuarios em JSON e envia na resposta.
	err = json.NewEncoder(w).Encode(blocos[bloco-1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Define o handler para a rota "/updateStates".
	http.HandleFunc("/updateStates", updateStatesHandler)

	// Nova rota para listar os usuários por blocos.
	http.HandleFunc("/usuarios", listarUsuariosHandler)

	// Inicia o servidor.
	fmt.Println("Servidor iniciado na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
