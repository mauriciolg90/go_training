package application

import (
    "github.com/mauriciolg90/go_training/api-using-ddd/domain/entity"
    "github.com/mauriciolg90/go_training/api-using-ddd/domain/repository"
    "github.com/mauriciolg90/go_training/api-using-ddd/domain/service"
    "github.com/mauriciolg90/go_training/api-using-ddd/domain/value"
)

// Singleton object (initialized to nil)
var finder Finder

// Represents the application to be called by interface layer
type Finder interface {
    GetPersons() ([]*entity.Person, error)
    GetPerson(id int64) (*entity.Person, error)
    AddPerson(name string) error
    GetPersonMatch(x, y int) (*entity.Person, error)
}

// It's the implementation of Finder
type FinderImpl struct{}

// Sets a new application with its implementation
func SetFinder(newFinder Finder) {
    finder = newFinder
}

// Returns the current application
func GetFinder() Finder {
    return finder
}

// Returns all persons stored in repository
func (t *FinderImpl) GetPersons() ([]*entity.Person, error) {
    return repository.GetPersonRepository().GetAll()
}

// Returns a person with the given id
func (t *FinderImpl) GetPerson(id int64) (*entity.Person, error) {
    return repository.GetPersonRepository().Get(id)
}

// Adds a new person with the given name
func (t *FinderImpl) AddPerson(name string) error {
    return repository.GetPersonRepository().Add(&entity.Person{Name: name,})
}

// Gets a matching person given a person's id and location (x, y)
func (t *FinderImpl) GetPersonMatch(x, y int) (*entity.Person, error) {
    return service.GetMatchingService().FindMatch(value.Location{X: x, Y: y,})
}
