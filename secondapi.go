package main
func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

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


/*
func main() {
	http.HandleFunc("/updateStates", corsMiddleware(updateStatesHandler))
	http.HandleFunc("/sendUserData/", corsMiddleware(sendUserDataHandler))

	fmt.Println("Servidor iniciado na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
/*