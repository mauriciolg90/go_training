package main

import (
    "github.com/mauriciolg90/go_training/api-using-ddd/application"
    "github.com/mauriciolg90/go_training/api-using-ddd/domain/repository"
    "github.com/mauriciolg90/go_training/api-using-ddd/domain/service"
    "github.com/mauriciolg90/go_training/api-using-ddd/infrastructure/persistence"
    "github.com/mauriciolg90/go_training/api-using-ddd/interface/web"
)

func main() {
    // We init a new database connection and a person repository
    sqlDB := persistence.GetDB("mysql", "root", "root", "tcp(localhost:3306)", "golang")
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
