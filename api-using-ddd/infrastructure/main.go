package main

import (
    "../application"
    "../domain/repository"
    "../interface/web"
    "../infrastructure/persistence"
)

func main() {
    // We init a new tinder application
    tinder := &application.TinderImpl{}
    application.SetTinder(tinder)

    // We init a new database connection and a person repository
    sqlDB := persistence.SetupDB("mysql", "root", "root", "tcp(localhost:3306)", "golang")
    personRepository := &persistence.PersonRepositoryImpl{DB: sqlDB}
    repository.SetPersonRepository(personRepository)

    router := web.Router()
    router.Run(":8080")
}
