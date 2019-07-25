package web

import (
    "github.com/gin-gonic/gin"
)

// Returns a http router with endpoint sets
func Router() *gin.Engine {
    router := gin.Default()
    router.GET("/persons", getPersons)
    router.GET("/persons/:id", getPerson)
    router.POST("/persons", addPerson)
    router.GET("/persons-match", getPersonMatch)
    return router
}
