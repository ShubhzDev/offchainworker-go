package main

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"
	"strconv"
)

type Validator struct {
	ID       int `json:"validator_id"`
	WorkDone int `json:"work_done"`
}

var validators map[int]int

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		handleGetRequest(w, r)
	} else if r.Method == http.MethodPost {
		handlePostRequest(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	var validator Validator

	err := json.NewDecoder(r.Body).Decode(&validator)
	if err != nil {
		log.Println("Failed to parse reward request:", err)
		http.Error(w, "Failed to parse reward request", http.StatusBadRequest)
		return
	}

	// Perform the necessary logic to handle the reward request
	// You can access the validator ID and work done from the `validator` struct
	// Replace this with your own logic to update the validator's balance or process the reward

	// Store the work done for the validator ID
	validators[validator.ID] = validator.WorkDone

	response := "Reward request processed successfully"

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	validatorIDStr := r.URL.Query().Get("validator_id")
	if validatorIDStr == "" {
		log.Println("Validator ID not found in the query parameters")
		http.Error(w, "Validator ID not found", http.StatusBadRequest)
		return
	}

	validatorID, err := strconv.Atoi(validatorIDStr)
	if err != nil {
		log.Println("Failed to convert validator ID to integer:", err)
		http.Error(w, "Failed to convert validator ID to integer", http.StatusBadRequest)
		return
	}

	// Lookup the work done for the provided validator ID
	workDone, ok := validators[validatorID]
	if !ok {
		log.Println("Work done value not found for validator ID:", validatorID)
		http.Error(w, "Work done value not found", http.StatusNotFound)
		return
	}

	// response := fmt.Sprintf("GET request processed successfully. Validator ID: %d, Work Done: %d", validatorID, workDone)
    response := strconv.Itoa(workDone)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func main() {
	validators = make(map[int]int)

	http.HandleFunc("/reward", handleRequest)

	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
