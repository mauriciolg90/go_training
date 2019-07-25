package repository

import (
    "../entity"
)

// Singleton object (initialized to nil)
var personRepository PersonRepository

// Represents a storage of all existing persons
type PersonRepository interface {
    GetAll() ([]*entity.Person, error)
    Get(ID int64) (*entity.Person, error)
    Add(person *entity.Person) error
}

// Sets a new repository with its implementation
func SetPersonRepository(newRepository PersonRepository) {
    personRepository = newRepository
}

// Returns the current repository
func GetPersonRepository() PersonRepository {
    return personRepository
}
