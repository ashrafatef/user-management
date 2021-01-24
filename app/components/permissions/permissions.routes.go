package permissions

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SetUpPermissionsRoutes(api fiber.Router) {

	api.Get("/permissions", Get)

	api.Post("/permissions/:id", func(c *fiber.Ctx) error {
		fmt.Println(c.Params("id"))
		return c.Status(http.StatusCreated).Send([]byte("Permission Created Successfully"))

	})
}
