package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type CRUDHandler interface {
	Create() fiber.Handler
	Update() fiber.Handler
	Read() fiber.Handler
	Delete() fiber.Handler
}

func main() {
	app := fiber.New()

	initCRUDHandler(app, "/api/v1/contacts", NewContactHandler())
	initCRUDHandler(app, "/api/v1/groups", NewGroupHandler())

	if err := app.Listen(":6080"); err != nil {
		log.Fatalln(err)
	}
}

func initCRUDHandler(app *fiber.App, path string, h CRUDHandler) {
	app.Get(path, h.Read())
	app.Post(path, h.Create())
	app.Put(path, h.Update())
	app.Delete(path, h.Delete())
}
