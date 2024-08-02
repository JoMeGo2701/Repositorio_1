package main

import (
	"Fiber/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

/*
	type User struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		EMAIL    string `json:"email"`
	}
*/
func main() {
	app := fiber.New()

	app.Use(logginMiddleware)

	routes.SetupRoutes(app)

	/*
		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("¡Hola, mundo!")
		})

		app.Get("/about", func(c *fiber.Ctx) error {
			return c.SendString("Página de información sobre nuestra aplicación")
		})

		app.Get("/contact", func(c *fiber.Ctx) error {
			return c.SendString("Página de contacto sobre nuestra aplicación")
		})

		//ruta dinamica
		app.Get("/saludo/:nombre", func(c *fiber.Ctx) error {
			nombre := c.Params("nombre")
			return c.SendString("¡Hola," + nombre + "!")
		})
	*/
	log.Fatal(app.Listen(":3000"))

}
func logginMiddleware(c *fiber.Ctx) error {
	log.Printf("Solicitud recibida: %s %s", c.Method(), c.Path())
	return c.Next()
}
