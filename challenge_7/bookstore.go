package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	router := gin.Default()

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/bookstore")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	router.GET("/books", func(c *gin.Context) {
		var books []Book

		rows, err := db.Query("SELECT * FROM books")
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var book Book
			err := rows.Scan(&book.ID, &book.Title, &book.Author)
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			books = append(books, book)
		}

		c.JSON(http.StatusOK, books)
	})

	router.GET("/books/:id", func(c *gin.Context) {
		var book Book

		id := c.Param("id")

		err := db.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&book.ID, &book.Title, &book.Author)
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, book)
	})

	router.POST("/books", func(c *gin.Context) {
		var book Book

		if err := c.ShouldBindJSON(&book); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		result, err := db.Exec("INSERT INTO books(title, author) VALUES(?, ?)", book.Title, book.Author)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		book.ID = int(id)

		c.JSON(http.StatusOK, book)
	})

	router.PUT("/books/:id", func(c *gin.Context) {
		var book Book

		id := c.Param("id")

		err := db.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&book.ID, &book.Title, &book.Author)
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if err := c.ShouldBindJSON(&book); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		_, err = db.Exec("UPDATE books SET title = ?, author = ? WHERE id = ?", book.Title, book.Author, id)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, book)
	})

	router.DELETE("/books/:id", func(c *gin.Context) {
		id := c.Param("id")

		_, err := db.Exec("DELETE FROM books WHERE id = ?", id)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusNoContent)
	})

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
