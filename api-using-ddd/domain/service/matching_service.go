package service

import (
    "math/rand"

    "github.com/mauriciolg90/go_training/api-using-ddd/domain/entity"
    "github.com/mauriciolg90/go_training/api-using-ddd/domain/repository"
    "github.com/mauriciolg90/go_training/api-using-ddd/domain/value"
)

// Singleton object (initialized to nil)
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

    // Find match according to the location (RANDOM implementation)
    rand.Seed(int64(location.X + location.Y))
    min, max := 0, len(persons)
    number := rand.Intn(max - min) + min
    match := persons[number]

    return match, nil
}