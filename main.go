package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
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

func findBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func getBookById(context *gin.Context) {
	id := context.Param("id")
	book, err := findBookById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, book)
}

func createBook(context *gin.Context) {
	var newBook book

	if err := context.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	context.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", getBookById)
	router.Run("localhost:8080")
}