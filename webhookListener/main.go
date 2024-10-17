package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// WebhookPayload - Structure of the Payload from GitHub Repositories
type WebhookPayload struct {
	Ref    string `json:"ref"`    // Ref - Branch or tag that was pushed
	Before string `json:"before"` // Before - The commit SHA before the Push
	After  string `json:"after"`  // After - The commit SHA after the Push
	// Repository - Information about the repository where the event occurred
	Repository struct {
		ID       int    `json:"id"`        // ID - The unique identifier of the repository
		Name     string `json:"name"`      // Name - The name of the repository
		FullName string `json:"full_name"` // FullName - The name of the repository (e.g., volvo-cars/action-cog)
		Private  bool   `json:"private"`   // Private - Whether the repository is Private.
		// Owner - Information about the owner of the repository
		Owner struct {
			Name  string `json:"name"`  // Name - Owner's Name
			Email string `json:"email"` // Email - Owner's Email
		} `json:"owner"`
	} `json:"repository"`
	// Pusher - Information about the user who pushed the commit
	Pusher struct {
		Name  string `json:"name"`  // Name - pusher's name
		Email string `json:"email"` // Email - pusher's email
	} `json:"pusher"`
	// Sender - Information about the user who triggered the event
	Sender struct {
		Login string `json:"login"` // Login - sender's Github username
		ID    int    `json:"id"`    // ID - sender's Github User ID
	} `json:"sender"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	//var payload WebhookPayload
	//err := json.NewDecoder(r.Body).Decode(&payload)
	//if err != nil {
	//	http.Error(w, "Bad Request Found", http.StatusBadRequest)
	//	return
	//}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request: Unable to Read Request Body", http.StatusBadRequest)
		log.Printf("Error reading request body:%v", err)
		return
	}
	if len(body) == 0 {
		_, _ = fmt.Fprintf(w, "Welcome to Webhook Endpoint !")
		return
	}
	var payload WebhookPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, "Bad Request: Invalid Json", http.StatusBadRequest)
		log.Printf("Error unmarshalling JSON: %v", err)
		return
	}

	// Process the payload
	fmt.Printf("Received webhook for repository: %s,ref: %s\n", payload.Repository.FullName, payload.Ref)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
