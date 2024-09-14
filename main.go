package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Roll is Model for Sushi API
type Roll struct {
	ID          string `json: "id"`
	Name        string `json: "name"`
	Description string `json: "description"`
	Ingerdient  string `json: "ingredient"`
}

var rolls []Roll

func getRolls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rolls)
}

func getRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range rolls {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newRoll Roll
	json.NewDecoder(r.Body).Decode(&newRoll)
	newRoll.ID = strconv.Itoa(len(rolls) + 1)
	rolls = append(rolls, newRoll)
	json.NewEncoder(w).Encode(newRoll)
}

func updateRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:index], rolls[+1:]...)
			var newRoll Roll
			json.NewDecoder(r.Body).Decode(&newRoll)
			newRoll.ID = params["id"]
			rolls = append(rolls, newRoll)
			json.NewEncoder(w).Encode(newRoll)
			return
		}
	}
}

func deleteRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:index], rolls[+1:]...)
			break
		}
	}
}

func main() {
	// Generate Mock Data
	rolls = append(rolls,
		Roll{
			ID:          "1",
			Name:        "California Roll",
			Description: "Sushi from California",
			Ingerdient:  "Cucumber, Avocado, Crab",
		},
		Roll{
			ID:          "2",
			Name:        "Salmon Roll",
			Description: "Sushi from Salmon",
			Ingerdient:  "Salmon, Cucumber",
		})

	router := mux.NewRouter()

	//handle end point or routing
	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}
