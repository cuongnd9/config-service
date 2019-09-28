package main

import (
	"io/ioutil"
	"net/http"
	"log"
	"encoding/json"

	"github.com/gorilla/mux"
)

type Service struct {
	Name string
	Url string	
}

type Schema struct {
	Services []Service
}

func getData(path string) Schema {
	data := Schema{}
	raw, _ := ioutil.ReadFile(path)
	json.Unmarshal(raw, &data)
	return data
}

func homeHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode("👋 Xin Chào Việt Nam 🇻🇳")
}

func devHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data := getData("./dev.json")
  json.NewEncoder(w).Encode(data)
}

func prodHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data := getData("./prod.json")
  json.NewEncoder(w).Encode(data)
}

func main()  {
	router := mux.NewRouter().StrictSlash(true)
	
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/dev", devHandler).Methods("GET")
	router.HandleFunc("/prod", prodHandler).Methods("GET")

	router.Use(mux.CORSMethodMiddleware(router))

	log.Fatal(http.ListenAndServe(":8888", router))
}
