package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Food struct {
	ID       int
	Name     string
	Category string
}

var (
	Foods []Food
)

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := "Hello REST API"
	response, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Error saat translate data", http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(Foods) < 1 {
		http.Error(w, "Data makanan tidak ditemukan", http.StatusNotFound)
	}

	response, err := json.Marshal(Foods)
	if err != nil {
		http.Error(w, "Error saat translate data", http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func CreateFoodHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var temporer Food
	if err := json.NewDecoder(r.Body).Decode(&temporer); err != nil {
		http.Error(w, "Error terdapat kesalahan pada input", http.StatusBadRequest)
		return
	}

	temporer.ID = len(Foods) + 1

	Foods = append(Foods, temporer)

	response, err := json.Marshal(temporer)
	if err != nil {
		http.Error(w, "Error saat translate data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func main() {
	fmt.Println("Hello World!")

	route := mux.NewRouter()
	route.HandleFunc("/", CheckHandler).Methods("GET")
	route.HandleFunc("/foods", GetAllHandler).Methods("GET")
	route.HandleFunc("/foods", CreateFoodHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", route))
}
