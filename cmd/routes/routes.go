// Package routes consolidates all Goman API routes.
package routes

import (
	"github.com/CommonBerry/goman/internal/core"
	"github.com/CommonBerry/goman/internal/infra"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App, db *infra.PostgresDataBase) {
	var idb core.IDataBase = db

	// Template Routes
	templateGroup := app.Group("/templates")

	templateGroup.Get("/", func(c fiber.Ctx) error {
		templates, err := idb.ListTemplates(c)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(templates)
	})

	templateGroup.Get("/:name", func(c fiber.Ctx) error {
		name := c.Params("name")
		template, err := idb.GetTemplateByName(c, name)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Template Not Found"})
		}

		return c.JSON(template)
	})

	templateGroup.Post("/", func(c fiber.Ctx) error {
		template := new(core.Template)

		if err := c.Bind().Body(template); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
		}

		if err := idb.CreateTemplate(c, template); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(201).JSON(template)
	})

	templateGroup.Put("/:uuid", func(c fiber.Ctx) error {
		template := new(core.Template)
		uuid := c.Params("uuid")

		if err := c.Bind().Body(template); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
		}

		if err := idb.UpdateTemplate(c, uuid, template); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(201).JSON(template)
	})

	templateGroup.Delete("/:uuid", func(c fiber.Ctx) error {
		uuid := c.Params("uuid")

		if err := idb.DeleteTemplate(c, uuid); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		return c.SendStatus(201)
	})

	// Alias Routes
	aliasGroup := app.Group("/aliases")

	aliasGroup.Get("/", func(c fiber.Ctx) error {
		aliases, err := idb.ListAliases(c)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(aliases)
	})

	aliasGroup.Get("/:name", func(c fiber.Ctx) error {
		name := c.Params("name")

		alias, err := idb.GetAliasByName(c, name)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Alias Not Found"})
		}

		return c.JSON(alias)
	})

	aliasGroup.Post("/", func(c fiber.Ctx) error {
		alias := new(core.Alias)

		if err := c.Bind().Body(alias); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
		}

		if err := idb.CreateAlias(c, alias); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(201).JSON(alias)
	})

	aliasGroup.Put("/:uuid", func(c fiber.Ctx) error {
		alias := new(core.Alias)
		uuid := c.Params("uuid")

		if err := c.Bind().Body(alias); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
		}

		if err := idb.UpdateAlias(c, uuid, alias); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(201).JSON(alias)
	})

	aliasGroup.Delete("/:uuid", func(c fiber.Ctx) error {
		uuid := c.Params("uuid")

		if err := idb.DeleteAlias(c, uuid); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		return c.SendStatus(201)
	})
}
