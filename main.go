
package main

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "net/http"
)

var API_KEY = "live_VJ8ucMcYkmK7tvdgMMUUR81PILYS3q0r4mqNsrLXx2ZPImhQzXAE3mrW0DzbJIqs"
var baseURL = "http://api.thedogapi.com/v1/images/search?has_breeds=true&order=RANDOM&page=0&limit=5&api_key=" + API_KEY

type dog struct {
    ID     string `json:"id"`
    Breeds []struct {
        Name string `json:"name"`
    } `json:"breeds"`
    URL string `json:"url"`
}

func fetchDogs() ([]dog, error) {
    resp, err := http.Get(baseURL)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var dogs []dog
    if err := json.NewDecoder(resp.Body).Decode(&dogs); err != nil {
        return nil, err
    }

    return dogs, nil
}

func getDogs(c *gin.Context) {
    dogs, err := fetchDogs()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusOK, dogs)
}

func getDogById(c *gin.Context) {
    id := c.Param("id")
    dogs, err := fetchDogs()
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    for _, d := range dogs {
        if d.ID == id {
            c.IndentedJSON(http.StatusOK, d)
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Dog not found"})
}

func main() {
    router := gin.Default()
    router.GET("/dogs", getDogs)
    router.GET("/dogs/:id", getDogById)
    router.Run("localhost:8080")
}