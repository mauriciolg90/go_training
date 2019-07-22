package routes

import (
    "../controllers"
    "github.com/gin-gonic/gin"
)

func CreateItemRoutes(router *gin.Engine, itemCtrl *controllers.ItemController) {
    router.GET("/items", itemCtrl.GetItemsHandler)
    router.GET("/item/:id", itemCtrl.GetItemHandler)
    router.PUT("/item/:id", itemCtrl.UpdateItemHandler)
    router.POST("/item", itemCtrl.CreateItemHandler)
    router.DELETE("/item/:id", itemCtrl.DeleteItemHandler)
}
