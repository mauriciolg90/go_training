package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/mauriciolg90/go_training/api-using-gin/controllers"
)

func CreateItemRoutes(router *gin.Engine, itemCtrl *controllers.ItemController) {
    router.GET("/items", itemCtrl.GetItemsHandler)
    router.GET("/item/:id", itemCtrl.GetItemHandler)
    router.PUT("/item/:id", itemCtrl.UpdateItemHandler)
    router.POST("/item", itemCtrl.CreateItemHandler)
    router.DELETE("/item/:id", itemCtrl.DeleteItemHandler)
}
