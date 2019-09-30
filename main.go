package main

import (
	"io/ioutil"
	"net/http"
	"log"
	"encoding/json"

	"github.com/gorilla/mux"
)

type Service struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type Schema struct {
	Services []Service `json:"services"`
}

func GetData(path string) Schema {
	data := Schema{}
	raw, _ := ioutil.ReadFile(path)
	json.Unmarshal(raw, &data)
	return data
}

func HomeHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode("👋 Xin Chào Việt Nam 🇻🇳")
}

func DevHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data := GetData("./dev.json")
  json.NewEncoder(w).Encode(data)
}

func ProdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data := GetData("./prod.json")
  json.NewEncoder(w).Encode(data)
}

func main()  {
	router := mux.NewRouter().StrictSlash(true)
	
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/dev", DevHandler).Methods("GET")
	router.HandleFunc("/prod", ProdHandler).Methods("GET")

	router.Use(mux.CORSMethodMiddleware(router))

	log.Fatal(http.ListenAndServe(":8888", router))
}
