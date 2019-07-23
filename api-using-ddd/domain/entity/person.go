package entity

// Represents a entity stored in repository
type Person struct {
    ID   int64  `json:"id"`
    Name string `json:"name"`
}
