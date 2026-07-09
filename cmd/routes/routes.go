// Package routes consolidates all Goman API routes.
package routes

import (
	"github.com/CommonBerry/goman/internal/core"
	"github.com/CommonBerry/goman/internal/infra"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App, db *infra.DataBase) {
	var idb core.IDataBase = db

	// Template Routes
	templateGroup := app.Group("/templates")

	templateGroup.Get("/", func(c fiber.Ctx) error {
		templates := idb.ListTemplates()
		return c.JSON(templates)
	})

	templateGroup.Get("/:name", func(c fiber.Ctx) error {
		name := c.Params("name")
		template, err := idb.GetTemplateByName(name)
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

		if err := idb.CreateTemplate(template); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(201).JSON(template)
	})

	templateGroup.Put("/:name", func(c fiber.Ctx) error {
		template := new(core.Template)
		oldName := c.Params("name")

		if err := c.Bind().Body(template); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
		}

		if err := idb.UpdateTemplate(oldName, template); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(201).JSON(template)
	})

	templateGroup.Delete("/:name", func(c fiber.Ctx) error {
		name := c.Params("name")

		if err := idb.DeleteTemplate(name); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		return c.SendStatus(201)
	})

	// Alias Routes
	aliasGroup := app.Group("/aliases")

	aliasGroup.Get("/", func(c fiber.Ctx) error {
		aliases := idb.ListAliases()
		return c.JSON(aliases)
	})

	aliasGroup.Get("/:name", func(c fiber.Ctx) error {
		name := c.Params("name")
		alias, err := idb.GetAliasByName(name)
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

		if err := idb.CreateAlias(alias); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(201).JSON(alias)
	})

	aliasGroup.Put("/:name", func(c fiber.Ctx) error {
		alias := new(core.Alias)
		oldName := c.Params("name")

		if err := c.Bind().Body(alias); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
		}

		if err := idb.UpdateAlias(oldName, alias); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(201).JSON(alias)
	})

	aliasGroup.Delete("/:name", func(c fiber.Ctx) error {
		name := c.Params("name")

		if err := idb.DeleteAlias(name); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		return c.SendStatus(201)
	})
}
