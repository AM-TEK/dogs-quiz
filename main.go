package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "errors"
)

type dog struct {
	ID		string	`json:"id"`
	Breed	string	`json:"breed"`	
	Image	string	`json:"image"`
}

var dogs = []dog {
	{ID: "1", Breed: "Beagle", Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQu-0gG2jbmtLHCOxJHvQr9c7uxbpDTMTx6wg&s"},
	{ID: "2", Breed: "Brittany Spaniel", Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQU2hATbC9ewdkgKzxs34lClrS8wI2EVZKgqQ&s"},
	{ID: "3", Breed: "German Shepherd", Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTsmqipwdA3Okk0gVFHbau3zk3O7_8eGrMTgQ&s"},
}

func getDogs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dogs)
}

func main() {
	router := gin.Default()
	router.GET("/dogs", getDogs)
	router.Run("localhost:8080")
}