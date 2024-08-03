package routes

import (
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	EMAIL    string `json:"email"`
}

func SetupRoutes(app *fiber.App) {
	renderer := Renderer()

	app.Get("/", handlerInicio)
	app.Get("/about", handlerAbout)
	app.Get("/contact", handlerContact)
	app.Get("/saludo/:nombre", handlerSaludo)

	app.Get("/:page", dynamicPageHandler(renderer))

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

func Renderer() *template.Template {
	return template.Must(template.ParseGlob("views/*.html"))

}

func dynamicPageHandler(_ *template.Template) fiber.Handler {
	return func(c *fiber.Ctx) error {
		page := c.Params("page")

		if strings.HasSuffix(page, ".html") {
			page = strings.TrimSuffix(page, ".html")
		}
		if _, err := os.Stat("views/" + page + ".html"); err == nil {
			return c.Render(page, nil)
		}
		return c.Status(http.StatusNotFound).SendString("Página no encontrada")
	}
}

func handlerInicio(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":   "Mi aplicacion",
		"Heading": "¡Hola, mundo!",
		"Message": "Bienvenido a mi aplicación web con Fiber y plantillas HTML.",
	})
	//return c.SendString("¡Hola, mundo!")
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
