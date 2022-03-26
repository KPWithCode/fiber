package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}
type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}
	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
		return err
	}
	err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "book has been created"})
	return nil
}

// routes
func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}


	if err != nil {
		log.Fatal("could not load the database")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}