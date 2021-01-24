package permissions

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx) error {
	fmt.Println("from get")
	return c.SendString("Hello, from get permissions!")
}
