package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henriqueteixeira.dev/go/models"
)

func main() {

	// data.Test()

	app := fiber.New()
	print("Hello")

	app.Get("/", func(c *fiber.Ctx) error {
		person := models.Person{Name: "joão", Age: 28}

		person.SetName("maria")

		return c.JSON(person)
	})

	persons := []models.Person{}

	app.Post("/", func(c *fiber.Ctx) error {
		person := new(models.Person)

		if err := c.BodyParser(person); err != nil {
			return err
		}

		persons = append(persons, *person)

		return c.JSON(persons)
	})

	app.Listen(":3000")
}
