package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/mriskyn/books-api.git/book"
	"github.com/mriskyn/books-api.git/database"
	// _ "github.com/mriskyn/books-api.git/database/dialects/sqlite"
	// "github.com/mriskyn/books-api.git/book"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello world")
}

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/book", book.GetBooks)
	app.Post("api/v1/book/", book.NewBook)
	app.Get("api/v1/book/:id", book.GetBook)
	app.Delete("api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error

	database.DBConn, err = gorm.Open("sqlite3", "books.db")

	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Database succesfully opened")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")

	if port == "" {
		port = "3333"
	}

	initDatabase()
	defer database.DBConn.Close()

	app.Get("/", helloWorld)
	setupRoutes(app)

	app.Listen(port)
}
