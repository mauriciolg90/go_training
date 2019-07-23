package application

import (
    "../domain/entity"
    "../domain/repository"
    "../domain/service"
    "../domain/value"
)

// Represents the application to be called by interface layer
type Tinder interface {
    GetPerson(id int64) (*entity.Person, error)
    GetPersons() ([]*entity.Person, error)
    AddPerson(name string) error
    GetPersonMatch(id int64, x, y int) (*entity.Person, error)
}

var tinder Tinder

// It's the implementation of Tinder
type TinderImpl struct{}

// Returns the current application
func GetTinder() Tinder {
    return tinder
}

// Sets a new application with its implementation
func SetTinder(newTinder Tinder) {
    tinder = newTinder
}

// Returns a person with the given id
func (t *TinderImpl) GetPerson(id int64) (*entity.Person, error) {
    return repository.GetPersonRepository().Get(id)
}

// Returns all persons stored in repository
func (t *TinderImpl) GetPersons() ([]*entity.Person, error) {
    return repository.GetPersonRepository().GetAll()
}

// Sets a new person with the given name to repository
func (t *TinderImpl) AddPerson(name string) error {
    return repository.GetPersonRepository().Save(&entity.Person{
        Name: name,
    })
}

// Gets a matching person from MatchingService given a person's id and x, y location
func (t *TinderImpl) GetPersonMatch(id int64, x, y int) (*entity.Person, error) {
    person, err := repository.GetPersonRepository().Get(id)

    if err != nil {
        return nil, err
    }

    return service.GetMatchingService().FindMatch(person, value.Location{
        X: x,
        Y: y,
    })
}
