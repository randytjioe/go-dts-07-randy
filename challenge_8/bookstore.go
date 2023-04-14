package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// Book adalah model data buku
type Book struct {
	ID     uint   `json:"id" gorm:"primarykey"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	// konfigurasi koneksi database
	dsn := "root:password@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// migrasi schema
	err = db.AutoMigrate(&Book{})
	if err != nil {
		panic("failed to migrate schema")
	}

	// inisialisasi router
	r := gin.Default()

	// handler untuk GET all books
	r.GET("/books", func(c *gin.Context) {
		var books []Book
		result := db.Find(&books)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		c.JSON(200, books)
	})

	// handler untuk GET book by ID
	r.GET("/books/:id", func(c *gin.Context) {
		var book Book
		id := c.Param("id")
		result := db.First(&book, id)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		c.JSON(200, book)
	})

	// handler untuk POST create book
	r.POST("/books", func(c *gin.Context) {
		var book Book
		err := c.ShouldBindJSON(&book)
		if err != nil {
			log.Fatal(err)
		}
		result := db.Create(&book)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		c.JSON(200, book)
	})

	// handler untuk PUT update book
	r.PUT("/books/:id", func(c *gin.Context) {
		var book Book
		id := c.Param("id")
		result := db.First(&book, id)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		err := c.ShouldBindJSON(&book)
		if err != nil {
			log.Fatal(err)
		}
		result = db.Save(&book)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		c.JSON(200, book)
	})

	// handler untuk DELETE book
	r.DELETE("/books/:id", func(c *gin.Context) {
		var book Book
		id := c.Param("id")
		result := db.Delete(&book, id)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Book with ID %s deleted", id),
		})
	})

	// jalankan aplikasi
	r.Run(":8001")
}
