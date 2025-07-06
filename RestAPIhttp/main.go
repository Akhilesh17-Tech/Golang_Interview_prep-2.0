package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getMessages(w, r)
		case http.MethodPost:
			createMessages(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

var messages = []Message{
	{ID: "1", Content: "Hello, world!"},
	{ID: "2", Content: "Welcome to the Go REST API."},
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	//logic to get messages
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func createMessages(w http.ResponseWriter, r *http.Request) {
	var newMessage Message
	err := json.NewDecoder(r.Body).Decode(&newMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	messages = append(messages, newMessage)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMessage)
}
