package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getMessages(w, r)
		case http.MethodPost:
			createMessage(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))

}

type Message struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Messages = []Message{
	{"John", 35},
	{"akhilesh", 25},
}

func getMessages(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "text/json")
	json.NewEncoder(w).Encode(Messages)
}

func createMessage(w http.ResponseWriter, r *http.Request) {
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	Messages = append(Messages, message)
	w.Header().Set("content-type", "text/json")
	json.NewEncoder(w).Encode(message)

}
