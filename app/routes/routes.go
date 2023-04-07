package routes

import (
	"go-training/app/controllers/authors"
	"go-training/app/controllers/books"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	{
		app.Post("/books", books.Create)
		app.Get("/books", books.Index)
		app.Get("/books/:id", books.Show)
		app.Patch("/books/:id", books.Update)
		app.Delete("/books/:id", books.Delete)
	}

	{
		app.Post("/authors", authors.Create)
		app.Get("/authors", authors.Index)
		app.Get("/authors/:id", authors.Show)
		app.Patch("/authors/:id", authors.Update)
		app.Delete("/authors/:id", authors.Delete)
	}
}
