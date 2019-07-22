package controllers

import (
    "log"
    "strconv"
    "net/http"
    "database/sql"

    "../repositories"
    "github.com/gin-gonic/gin"
)

type ItemInterface interface {
    GetItemHandler(context *gin.Context)
    GetItemsHandler(context *gin.Context)
    CreateItemHandler(context *gin.Context)
    UpdateItemHandler(context *gin.Context)
    DeleteItemHandler(context *gin.Context)
}

type ItemController struct {
    dbPtr *sql.DB
}

func NewItemController(db *sql.DB) *ItemController {
    return &ItemController{dbPtr: db}
}

func (itemCtrl *ItemController) GetItemHandler(context *gin.Context) {
    // Get parameters from the request
    id := context.Params.ByName("id")
    idInteger, err := strconv.Atoi(id)
    if err != nil {
        context.JSON(http.StatusBadRequest, map[string]string{"error": "Bad item id",})
        return
    }
    // Call to the API to build a query
    item, err := repositories.GetItem(itemCtrl.dbPtr, idInteger)
    if err != nil {
        if err == sql.ErrNoRows {
            context.String(http.StatusNotFound, "Item not found!")
            return
        }
        context.String(http.StatusInternalServerError, "API error!")
        log.Println("err", err)
        return
    }
    // Query was successful!
    context.JSON(http.StatusOK, item)
}

func (itemCtrl *ItemController) GetItemsHandler(context *gin.Context) {
    // Call to the API to build a query
    items, err := repositories.GetItems(itemCtrl.dbPtr)
    if err != nil {
        context.String(http.StatusInternalServerError, "API error!")
        log.Println("err", err)
        return
    }
    // Query was successful!
    context.JSON(http.StatusOK, items)
}

func (itemCtrl *ItemController) CreateItemHandler(context *gin.Context) {
    // Get parameters from the request
    title := context.PostForm("title")
    description := context.PostForm("description")
    // Call to the API to build a query
    err := repositories.CreateItem(itemCtrl.dbPtr, title, description)
    if err != nil {
        context.String(http.StatusInternalServerError, "API error!")
        log.Println("err", err)
        return
    }
    // Query was successful!
    context.String(http.StatusCreated, "Item was created!")
}

func (itemCtrl *ItemController) UpdateItemHandler(context *gin.Context) {
    // Get parameters from the request
    id := context.Params.ByName("id")
    idInteger, err := strconv.Atoi(id)
    title := context.PostForm("title")
    description := context.PostForm("description")
    if err != nil {
        context.JSON(http.StatusBadRequest, map[string]string{"error": "Bad item id",})
        log.Println("err", err)
        return
    }
    // Call to the API to build a query
    err = repositories.UpdateItem(itemCtrl.dbPtr, idInteger, title, description)
    if err != nil {
        context.String(http.StatusInternalServerError, "API error!")
        log.Println("err", err)
        return
    }
    // Query was successful!
    context.String(http.StatusOK, "")
}

func (itemCtrl *ItemController) DeleteItemHandler(context *gin.Context) {
    // Get parameters from the request
    id := context.Param("id")
    idInteger, err := strconv.Atoi(id)
    if err != nil {
        context.JSON(http.StatusBadRequest, map[string]string{"error": "Bad item id",})
        return
    }
    // Call to the API to build a query
    err = repositories.DeleteItem(itemCtrl.dbPtr, idInteger)
    if err != nil {
        context.String(http.StatusInternalServerError, "API error!")
        log.Println("err", err)
        return
    }
    // Query was successful!
    context.String(http.StatusOK, "")
}
