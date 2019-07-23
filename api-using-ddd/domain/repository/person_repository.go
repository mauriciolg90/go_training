package repository

import (
    "../entity"
)

// Represents a storage of all existing persons
type PersonRepository interface {
    Get(ID int64) (*entity.Person, error)
    GetAll() ([]*entity.Person, error)
    Save(person *entity.Person) error
}

var personRepository PersonRepository

// Sets a new repository with its implementation
func SetPersonRepository(newRepository PersonRepository) {
    personRepository = newRepository
}

// Returns the current repository
func GetPersonRepository() PersonRepository {
    return personRepository
}
