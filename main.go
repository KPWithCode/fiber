package main
import (
	"github.com/gofiber/fiber/v2"
	"log"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)
type Book struct {
	Author		string		`json:"author"`
	Title 		string		`json:"title"`	
	Publisher	string		`json:"publisher"`

}
type Repository struct{
	DB *gorm.DB
}

// routes
func(r *Repository) SetupRoutes(app *fiber.App) {

}



func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal("could not load the database")
	}


	app := fiber.New()
	app.Listen(":8080")
}