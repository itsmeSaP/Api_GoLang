package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itsmeSaP/starting-go/pkg/database"
	"github.com/itsmeSaP/starting-go/pkg/models"
)

func healthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Status": "Ok",
	})

}

func addPost(c *fiber.Ctx) error {
	u := new(models.Posts)
	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error":   true,
			"Message": "Invalid Input",
		})
	}
	db := database.GetDB()
	err := db.Create(&u).Error
	// post, err := controllers.AddPosts(u.Title, u.Description)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error":   true,
			"Message": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Error":   false,
		"Message": "Post Created",
		"Post":    u,
	})
}

//create 1 route to check health of our goland application
func MountRoutes(app *fiber.App) {
	app.Get("/", healthCheck)
	app.Post("/api/v1/posts/create", addPost)
}
