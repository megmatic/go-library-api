package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID			string 	`json:"id"`
	Title		string 	`json:"title"`
	Author		string 	`json:"author"`
	Quantity 	int	   	`json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Project Hail Mary", Author: "Andy Weir", Quantity: 15},
	{ID: "2", Title: "The Fifth Season", Author: "N. K. Jemisin", Quantity: 5},
	{ID: "3", Title: "The Seven Husbands of Evelyn Hugo", Author: "Taylor Jenkins Reid", Quantity: 5},
}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)

}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}