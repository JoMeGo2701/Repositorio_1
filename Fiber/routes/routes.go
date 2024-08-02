package routes

import "github.com/gofiber/fiber/v2"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	EMAIL    string `json:"email"`
}

func SetupRoutes(app *fiber.App) {

	app.Get("/", handlerInicio)
	app.Get("/about", handlerAbout)
	app.Get("/contact", handlerContact)
	app.Get("/saludo/:nombre", handlerSaludo)

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

	app.Post("/api/usuarios", func(c *fiber.Ctx) error {
		var usuario User
		if err := c.BodyParser(&usuario); err != nil {
			return err
		}
		return c.JSON(usuario)

	})
}

func handlerInicio(c *fiber.Ctx) error {
	return c.SendString("¡Hola, mundo!")
}

func handlerAbout(c *fiber.Ctx) error {
	return c.SendString("Página de información sobre nuestra aplicación")
}

func handlerContact(c *fiber.Ctx) error {
	return c.SendString("Página de contacto sobre nuestra aplicación")
}

func handlerSaludo(c *fiber.Ctx) error {
	nombre := c.Params("nombre")
	return c.SendString("¡Hola," + nombre + "!")
}
