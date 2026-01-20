package main

import (
	mwhttp "github.com/middleware-labs/golang-apm-http/http"
	track "github.com/middleware-labs/golang-apm/tracker"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Simulated database
var users = map[int]User{
	1: {ID: 1, Name: "Alice", Email: "alice@example.com"},
	2: {ID: 2, Name: "Bob", Email: "bob@example.com"},
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get user ID from query parameter
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "missing user id", http.StatusBadRequest)
		return
	}

	var id int
	fmt.Sscanf(idStr, "%d", &id)

	user, exists := users[id]
	if !exists {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user.ID = len(users) + 1
	users[user.ID] = user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func main() {
	config, err := track.Track(
		track.WithConfigTag(track.Service, "middleware-go-demo"),
		track.WithConfigTag("accessToken", "<your-access-token>"),
		track.WithConfigTag("target", "<your-target-url>"),
		// track.WithConfigTag("debug", true),
		// track.WithConfigTag("pauseMetrics", true),
	)

	if err != nil {
		log.Fatalf("failed to initialize Middleware APM: %v", err)
	}

	_ = config // optional if you don’t use it yet

	http.Handle("/user", mwhttp.MiddlewareHandler(http.HandlerFunc(getUserHandler), "getUserHandler"))
	http.Handle("/user/create", mwhttp.MiddlewareHandler(http.HandlerFunc(createUserHandler), "createUserHandler"))

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
