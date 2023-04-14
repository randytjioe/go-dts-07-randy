package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	router := gin.Default()

	// Add sample books
	books = append(books, Book{ID: 1, Title: "Dune", Author: "Frank Herbert"})
	books = append(books, Book{ID: 2, Title: "Foundation", Author: "Isaac Asimov"})

	// Get all books
	router.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, books)
	})

	// Get book by ID
	router.GET("/books/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		for _, b := range books {
			if b.ID == id {
				c.JSON(http.StatusOK, b)
				return
			}
		}

		c.AbortWithStatus(http.StatusNotFound)
	})

	// Add book
	router.POST("/books", func(c *gin.Context) {
		var b Book
		if err := c.ShouldBindJSON(&b); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		b.ID = len(books) + 1
		books = append(books, b)

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	// Update book
	router.PUT("/books/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		var b Book
		if err := c.ShouldBindJSON(&b); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		for i, oldB := range books {
			if oldB.ID == id {
				b.ID = id
				books[i] = b
				c.JSON(http.StatusOK, gin.H{"status": "success"})
				return
			}
		}

		c.AbortWithStatus(http.StatusNotFound)
	})

	// Delete book
	router.DELETE("/books/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		for i, b := range books {
			if b.ID == id {
				books = append(books[:i], books[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"status": "success"})
				return
			}
		}

		c.AbortWithStatus(http.StatusNotFound)
	})

	router.Run(":8080")
}
