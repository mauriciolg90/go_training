package main

import (
    "../application"
    "../domain/repository"
    "../domain/service"
    "../infrastructure/persistence"
    "../interface/web"
)

func main() {
    // We init a new database connection and a person repository
    sqlDB := persistence.SetupDB("mysql", "root", "root", "tcp(localhost:3306)", "golang")
    personRepository := &persistence.PersonRepositoryImpl{DB: sqlDB}
    repository.SetPersonRepository(personRepository)

    // We init a new matching service
    matchingService := &service.MatchingServiceImpl{}
    service.SetMatchingService(matchingService)

    // We init a new finder application
    finder := &application.FinderImpl{}
    application.SetFinder(finder)

    router := web.Router()
    router.Run(":8080")
}
