package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "../models"
    "github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, endpoint string) *httptest.ResponseRecorder {
    req, _ := http.NewRequest(method, endpoint, nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    return w
}

func TestGetItems(t *testing.T) {
    // Setup database and router
    db := SetupDatabase()
    router := SetupRouter(db)

    // Perform a GET request with that router
    w := performRequest(router, "GET", "/items")

    // Assert we encoded correctly the request gives a 200
    assert.Equal(t, http.StatusOK, w.Code)

    // Convert the JSON response to a slice
    var response []models.Item
    err := json.Unmarshal([]byte(w.Body.String()), &response)

    // Assert we decoded correctly the response
    assert.Nil(t, err)

    // Assert all the fields
    for _, item := range response {
        assert.Greater(t, item.Id, 0)
        assert.NotEmpty(t, item.Title)
        assert.NotEmpty(t, item.Description)
    }
}
