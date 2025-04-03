package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type ContactHandler struct {
}

func NewContactHandler() CRUDHandler {
	return &ContactHandler{}
}

func (c *ContactHandler) Read() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.Println("read contact")
		return ctx.SendString("user")
	}
}

func (c *ContactHandler) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		return ctx.SendString("create user")
	}
}

func (c *ContactHandler) Update() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		return ctx.SendString("update user")
	}
}

func (c *ContactHandler) Delete() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		return ctx.SendString("delete user")
	}
}
