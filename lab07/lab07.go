package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	// TODO: Finish struct
	ID    int    `json:"id"`
	NAME  string `json:"name"`
	PAGES int    `json:"pages"`
}

type B struct {
	NAME  string `json:"name"`
	PAGES int    `json:"pages"`
}

var top int = 2

var bookshelf = []Book{
	// TODO: Init bookshelf
	{
		ID:    1,
		NAME:  "Blue Bird",
		PAGES: 500,
	},
}

func getBooks(c *gin.Context) {
	if len(bookshelf) == 0 {
		c.JSON(200, []Book{})
	} else {
		c.JSON(200, bookshelf)
	}
	return
}
func getBook(c *gin.Context) {
	ID := c.Param("id")
	BookID, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println("no error")
	}
	fmt.Println(BookID)
	for _, book := range bookshelf {
		if book.ID == BookID {
			c.JSON(200, book)
			return
		}
	}
	c.JSON(404, gin.H{"message": "book not found"})
	return
}
func addBook(c *gin.Context) {
	var b B
	c.ShouldBindJSON(&b)
	fmt.Println(b)
	for _, book := range bookshelf {
		if book.NAME == b.NAME {
			c.JSON(409, gin.H{"message": "duplicate book name"})
			return
		}
	}
	var book Book
	book.ID = top
	book.NAME = b.NAME
	book.PAGES = b.PAGES
	bookshelf = append(bookshelf, book)
	fmt.Println(bookshelf)
	top++
	c.JSON(201, book)
	return
}
func deleteBook(c *gin.Context) {
	var Bookshelf []Book
	ID := c.Param("id")
	BookID, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println("no error")
	}
	for _, book := range bookshelf {
		if book.ID != BookID {
			Bookshelf = append(Bookshelf, book)
		}
	}
	bookshelf = Bookshelf
	fmt.Println(bookshelf)
	c.JSON(204, "")
	return
}
func updateBook(c *gin.Context) {
	ID := c.Param("id")
	BookID, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println("no error")
	}
	var b B
	c.ShouldBindJSON(&b)
	fmt.Println(b)
	for _, book := range bookshelf {
		if book.NAME == b.NAME {
			c.JSON(409, gin.H{"message": "duplicate book name"})
			return
		}
	}
	for index, book := range bookshelf {
		if book.ID == BookID {
			var update Book
			update.ID = book.ID
			update.NAME = b.NAME
			update.PAGES = b.PAGES
			bookshelf[index] = update
			c.JSON(200, update)
			return
		}
	}
	c.JSON(404, gin.H{"message": "book not found"})
	return
}

func main() {
	r := gin.Default()
	r.RedirectFixedPath = true

	// TODO: Add routes
	r.GET("/bookshelf", getBooks)
	r.GET("/bookshelf/:id", getBook)
	r.POST("/bookshelf", addBook)
	r.DELETE("/bookshelf/:id", deleteBook)
	r.PUT("/bookshelf/:id", updateBook)

	err := r.Run(":8087")
	if err != nil {
		return
	}
}
