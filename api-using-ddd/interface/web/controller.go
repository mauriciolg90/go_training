package web

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "../../application"
)

func getPersons(context *gin.Context) {
    persons, err := application.GetTinder().GetPersons()
    // Check error if any
    if err != nil {
        context.JSON(http.StatusInternalServerError, err)
        log.Println("Error:", err)
        return
    }
    // Operation was successful!
    context.JSON(http.StatusOK, persons)
}

func getPerson(context *gin.Context) {
    id, err := strconv.ParseInt(context.Param("id"), 10, 64)

    if err != nil {
        context.JSON(http.StatusBadRequest, err)
        return
    }

    person, err := application.GetTinder().GetPerson(id)

    if err != nil {
        context.JSON(http.StatusInternalServerError, err)
        return
    }

    context.JSON(http.StatusOK, person)
}

func addPerson(context *gin.Context) {
    rawData, _ := context.GetRawData()
    data := struct {
        Name string `json:"name"`
    }{}
    err := json.Unmarshal(rawData, &data)

    if err != nil {
        context.JSON(http.StatusBadRequest, err)
        return
    }

    err = application.GetTinder().AddPerson(data.Name)

    if err != nil {
        context.JSON(http.StatusInternalServerError, err)
        return
    }

    context.JSON(http.StatusOK, nil)
}

func getPersonMatch(context *gin.Context) {
    id, err := strconv.ParseInt(context.Param("id"), 10, 64)

    if err != nil {
        context.JSON(http.StatusBadRequest, err)
        return
    }

    rawData, _ := context.GetRawData()
    data := struct {
        X int `json:"x"`
        Y int `json:"y"`
    }{}
    err = json.Unmarshal(rawData, &data)

    if err != nil {
        context.JSON(http.StatusBadRequest, err)
        return
    }

    match, err := application.GetTinder().GetPersonMatch(id, data.X, data.Y)

    if err != nil {
        context.JSON(http.StatusInternalServerError, err)
        return
    }

    context.JSON(http.StatusOK, match)
}
