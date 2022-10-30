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
	{ID: "1", Title: "Project Hail Mary", Author: "Andy Weir", Quantity: 2},
	{ID: "2", Title: "The Fifth Season", Author: "N. K. Jemisin", Quantity: 2},
	{ID: "3", Title: "The Seven Husbands of Evelyn Hugo", Author: "Taylor Jenkins Reid", Quantity: 2},
}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func checkoutBook(context *gin.Context) {
	id, ok := context.GetQuery("id")

	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := findBookById((id))

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	if book.Quantity <= 0 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available."})
		return
	}

	book.Quantity -= 1
	context.IndentedJSON(http.StatusOK, book)
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
	router.GET("/books/:id", getBookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.Run("localhost:8080")
}