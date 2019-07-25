package service

import (
    "../entity"
    "../repository"
    "../value"
)

// Internal var (initialized to nil)
var matchingService MatchingService

// Represents the service to find a match
type MatchingService interface {
    FindMatch(location value.Location) (*entity.Person, error)
}

// It's the implementation of MatchingService
type MatchingServiceImpl struct{}

// Sets a new service with its implementation
func SetMatchingService(newService MatchingService) {
    matchingService = newService
}

// Returns the current service
func GetMatchingService() MatchingService {
    return matchingService
}

// Finds a match for a person based on its location
func (s *MatchingServiceImpl) FindMatch(location value.Location) (*entity.Person, error) {
    // Get all persons stored in repository
    persons, err := repository.GetPersonRepository().GetAll()

    if err != nil {
        return nil, err
    }

    // FIXME, aca buscar un match en base a la ubicacion
    match := persons[0]

    return match, nil
}