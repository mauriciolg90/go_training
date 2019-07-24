package service

import (
    "errors"

    "../entity"
    "../repository"
    "../value"
)

// Represents the service to find a match
type MatchingService interface {
    FindMatch(person *entity.Person, location value.Location) (*entity.Person, error)
}

var matchingService MatchingService

// It's the implementation of MatchingService
type MatchingServiceImpl struct{}

// Returns the current service
func GetMatchingService() MatchingService {
    return matchingService
}

// Sets a new service with its implementation
func SetMatchingService(newService MatchingService) {
    matchingService = newService
}

// Finds a match for a person based on its location
func (s *MatchingServiceImpl) FindMatch(person *entity.Person, location value.Location) (*entity.Person, error) {
    if person == nil {
        return nil, errors.New("Nil person")
    }

    persons, err := repository.GetPersonRepository().GetAll()

    if err != nil {
        return nil, err
    }

    randomNumber := location.X + location.Y - location.X - location.Y
    match := persons[randomNumber]
    if match.ID == person.ID {
        match = persons[(randomNumber+1)%len(persons)]
    }

    return match, nil
}