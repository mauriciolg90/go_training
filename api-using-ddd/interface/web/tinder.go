package web

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "../../application"
    "net/http"
    "strconv"
)

func getPerson(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)

    if err != nil {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    person, err := application.GetTinder().GetPerson(id)

    if err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }

    c.JSON(http.StatusOK, person)
}

func getPersons(c *gin.Context) {
    persons, err := application.GetTinder().GetPersons()

    if err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }

    c.JSON(http.StatusOK, persons)
}

func addPerson(c *gin.Context) {
    rawData, _ := c.GetRawData()
    data := struct {
        Name string `json:"name"`
    }{}
    err := json.Unmarshal(rawData, &data)

    if err != nil {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    err = application.GetTinder().AddPerson(data.Name)

    if err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }

    c.JSON(http.StatusOK, nil)
}

func getPersonMatch(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)

    if err != nil {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    rawData, _ := c.GetRawData()
    data := struct {
        X int `json:"x"`
        Y int `json:"y"`
    }{}
    err = json.Unmarshal(rawData, &data)

    if err != nil {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    match, err := application.GetTinder().GetPersonMatch(id, data.X, data.Y)

    if err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }

    c.JSON(http.StatusOK, match)
}
