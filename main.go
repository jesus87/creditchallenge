package main

import (
	"creditchallenge/assigncredits"
	"creditchallenge/domain/entity"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/credit-assignment", creditAssignment).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", router))
}

func creditAssignment(w http.ResponseWriter, r *http.Request) {
	var investment entity.Investment
	err := json.NewDecoder(r.Body).Decode(&investment)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c := assigncredits.New()
	a300, a500, a700, err := c.Asssign(investment.Investment)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(entity.CreditResult{
		CreditType1: a300,
		CreditType2: a500,
		CreditType3: a700,
	})
}
