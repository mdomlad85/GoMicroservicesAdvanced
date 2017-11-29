package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/mdomlad85/GoMicroservices/api"
)

func main() {
	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// the hello message endpoint with JSON response from map
	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin Framework."})
	})

	// get all books
	engine.GET("/api/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, api.AllBooks())
	})

	// create new book
	engine.POST("/api/books", func(c *gin.Context) {
		var book api.Book
		if c.BindJSON(&book) == nil {
			isbn, created := api.CreateBook(book)
			if created {
				c.Header("Location", "/api/books/"+isbn)
				c.Status(http.StatusCreated)
			} else {
				c.Status(http.StatusConflict)
			}
		}
	})

	// get book by isbn
	engine.GET("/api/books/:isbn", func(c *gin.Context) {
		isbn := c.Params.ByName("isbn")
		book, found := api.GetBook(isbn)
		if found {
			c.JSON(http.StatusOK, book)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})

	// update book
	engine.PUT("/api/books/:isbn", func(c *gin.Context) {
		isbn := c.Params.ByName("isbn")

		var book api.Book
		if c.BindJSON(&book) == nil {
			exists := api.UpdateBook(isbn, book)
			if exists {
				c.Status(http.StatusOK)
			} else {
				c.Status(http.StatusConflict)
			}
		}
	})

	engine.DELETE("/api/books/:isbn", func(c *gin.Context) {
		isbn := c.Params.ByName("isbn")
		api.DeleteBook(isbn)
		c.Status(http.StatusOK)
	})

	// configuration for static files and templates
	engine.LoadHTMLGlob("./templates/*.html")
	engine.StaticFile("/favico.ico", "./favico.ico")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Advanced Cloud Native Go with Gin",
		})
	})

	// run server on PORT
	engine.Run(port())
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
