package main

import (
    "log"
    "strconv"
    "net/http"
    "math/rand"
    "encoding/json"
    "github.com/gorilla/mux"
)

// Item struct
type Item struct {
    Id      string  `json: "id"`
    Name    string  `json: "name"`
    Price   string  `json: "price"`
    Vendor  *Vendor `json: "vendor"`
}

// Vendor struct
type Vendor struct {
    FirstName   string  `json:"firstname"`
    LastName    string  `json:"lastname"`
}

// Define items variable as a slice
var items []Item

// GetItemsHandler - GET - /api/items
func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

// GetItemHandler - GET - /api/items/{id}
func GetItemHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r) // Get params from the request
    // Iter over items and find id
    for _, item := range items {
        if item.Id == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    // Encode an empty item if the id does't exist
    json.NewEncoder(w).Encode(&Item{})
}

// CreateItemHandler - POST - /api/items
func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
    var newItem Item
    w.Header().Set("Content-Type", "application/json")
    err := json.NewDecoder(r.Body).Decode(&newItem)
    // Get panic if error on json decoding
    if err != nil {
        panic(err)
    }
    newItem.Id = strconv.Itoa(rand.Intn(1000000))
    items = append(items, newItem)
    json.NewEncoder(w).Encode(items)
}

// UpdateItemHandler - PUT - /api/items/{id}
func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
    var newItem Item
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r) // Get params from the request
    // Iter over items and find id
    for index, item := range items {
        if item.Id == params["id"] {
            items = append(items[:index], items[index+1:]...) // Remove the item
            _ = json.NewDecoder(r.Body).Decode(&newItem)
            // Insert the updated item
            newItem.Id = params["id"]
            items = append(items, newItem)
            json.NewEncoder(w).Encode(newItem)
            return
        }
    }
    // Encode old items if the id does't exist
    json.NewEncoder(w).Encode(items)
}

// DeleteItemHandler - DELETE - /api/items/{id}
func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r) // Get params from the request
    // Iter over items and find id
    for index, item := range items {
        if item.Id == params["id"] {
            items = append(items[:index], items[index+1:]...) // Remove the item
            break
        }
    }
    // Encode the items
    json.NewEncoder(w).Encode(items)
}

func main() {
    // Initialize router
    router := mux.NewRouter()

    // [TODO] Implement DB (populate slice for test purpose only)
    items = append(
        items,
        Item{Id: "1", Name: "Go Programming Language", Price: "800", Vendor: &Vendor{FirstName: "Leonardo", LastName: "Gonzalez"}})
    items = append(
        items,
        Item{Id: "2", Name: "Go In Practice", Price: "650", Vendor: &Vendor{FirstName: "Mauricio", LastName: "Gonzalez"}})

    // Configure endpoints/handlers
    router.HandleFunc("/api/items", GetItemsHandler).Methods("GET")
    router.HandleFunc("/api/items", CreateItemHandler).Methods("POST")
    router.HandleFunc("/api/items/{id}", GetItemHandler).Methods("GET")
    router.HandleFunc("/api/items/{id}", UpdateItemHandler).Methods("PUT")
    router.HandleFunc("/api/items/{id}", DeleteItemHandler).Methods("DELETE")

    // Let's start the server!
    log.Fatal(http.ListenAndServe(":8000", router))
}