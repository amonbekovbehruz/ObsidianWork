package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserService handles user operations
type UserService struct {
	users  map[int]User
	nextID int
	mu     sync.RWMutex
}

// NewUserService creates a new UserService
func NewUserService() *UserService {
	return &UserService{
		users:  make(map[int]User),
		nextID: 1,
	}
}

// Create creates a new user
func (s *UserService) Create(user User) User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user.ID = s.nextID
	s.nextID++
	s.users[user.ID] = user
	return user
}

// Get retrieves a user by ID
func (s *UserService) Get(id int) (User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[id]
	return user, exists
}

// Update updates an existing user
func (s *UserService) Update(id int, updatedUser User) (User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		return User{}, false
	}

	updatedUser.ID = id
	s.users[id] = updatedUser
	return updatedUser, true
}

// Delete removes a user by ID
func (s *UserService) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		return false
	}

	delete(s.users, id)
	return true
}

// GetAll returns all users
func (s *UserService) GetAll() []User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

var userService = NewUserService()

// createUserHandler handles POST /users
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdUser := userService.Create(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

// getUserHandler handles GET /users/{id}
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Simple validation
	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Convert string to int (simple implementation)
	var userID int
	fmt.Sscanf(id, "%d", &userID)

	user, exists := userService.Get(userID)
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// updateUserHandler handles PUT /users/{id}
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	var userID int
	fmt.Sscanf(id, "%d", &userID)

	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, exists := userService.Update(userID, updatedUser)
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// deleteUserHandler handles DELETE /users/{id}
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	var userID int
	fmt.Sscanf(id, "%d", &userID)

	if !userService.Delete(userID) {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// getAllUsersHandler handles GET /users
func getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := userService.GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	r := mux.NewRouter()

	// User endpoints
	r.HandleFunc("/users", getAllUsersHandler).Methods("GET")
	r.HandleFunc("/users", createUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", updateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUserHandler).Methods("DELETE")

	fmt.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
