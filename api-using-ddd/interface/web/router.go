package web

import (
    "github.com/gin-gonic/gin"
)

// Returns a http router with endpoint sets
func Router() *gin.Engine {
    router := gin.Default()
    router.GET("/persons", getPersons)
    router.GET("/persons/:id", getPerson)
    router.GET("/persons-match/:id", getPersonMatch)
    router.POST("/persons", addPerson)
    return router
}
