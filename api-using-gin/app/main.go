package main

import (
    "log"
    "database/sql"

    "../routes"
    "../database"
    "../controllers"
    "github.com/gin-gonic/gin"
)

func main() {
    sqlConn := SetupDatabase()
    ginRouter := SetupRouter(sqlConn)
    ginRouter.Run(":8080")
}

func SetupDatabase() *sql.DB {
    user := "root"
    pass := "root"
    driver := "mysql"
    schema := "golang"
    host := "tcp(localhost:3306)"
    conn, _ := database.ConnectSql(driver, user, pass, host, schema)
    // Test connection
    err := conn.Ping()
    if err != nil {
        log.Fatal("Error connecting database!\n", err)
    }
    return conn
}

func SetupRouter(db *sql.DB) *gin.Engine {
    router := gin.Default()
    routes.CreateItemRoutes(router, controllers.NewItemController(db))
    return router
}
