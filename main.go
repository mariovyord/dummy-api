package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/users/id/", handleSingleUser)
	http.HandleFunc("/cars", handleCars)
	http.HandleFunc("/apartments", handleApartments)

	fmt.Println("API is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleRoot(w http.ResponseWriter, _ *http.Request) {
	availableRoutes := []string{
		"/users",
		"/cars",
		"/apartments",
	}

	str := "Hello, this a dummy API. Available routes: " + strings.Join(availableRoutes, ", ")

	io.WriteString(w, str)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./data/users.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()

	var users []User
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func handleSingleUser(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./data/users.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()

	var users []User
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id, err := strconv.Atoi(r.URL.Path[len("/users/id/"):])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	user := users[id]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func handleCars(w http.ResponseWriter, _ *http.Request) {
	file, err := os.Open("./data/cars.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()

	var cars []Car
	err = json.NewDecoder(file).Decode(&cars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func handleApartments(w http.ResponseWriter, _ *http.Request) {
	file, err := os.Open("./data/apartments.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()

	var apartments []Apartment
	err = json.NewDecoder(file).Decode(&apartments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apartments)
}

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	Description string `json:"description"`
}

type Car struct {
	ID      int    `json:"id"`
	Maker   string `json:"maker"`
	Model   string `json:"model"`
	Year    int    `json:"year"`
	Color   string `json:"color"`
	OwnerID int    `json:"ownerId"`
}

type Apartment struct {
	ID        int    `json:"id"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zipcode   string `json:"zipcode"`
	Bedrooms  int    `json:"bedrooms"`
	Bathrooms int    `json:"bathrooms"`
	Rent      int    `json:"rent"`
	OwnerID   int    `json:"ownerId"`
}
