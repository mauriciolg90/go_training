package web

import (
    "github.com/gin-gonic/gin"
)

// Returns a http router with endpoints set
func Router() *gin.Engine {
    router := gin.Default()

    router.GET("/persons/:id", getPerson)
    router.GET("/persons", getPersons)
    router.POST("/persons", addPerson)
    router.GET("/persons-match/:id", getPersonMatch)

    return router
}
