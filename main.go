package main

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode("👋 Chao Xìn")
}

func main()  {
	router := mux.NewRouter()
	
	router.HandleFunc("/", homeHandler).Methods("GET")

	http.ListenAndServe(":8000", router)
}
