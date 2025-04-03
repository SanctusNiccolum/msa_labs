package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type GroupHandler struct {
}

func NewGroupHandler() CRUDHandler {
	return &ContactHandler{}
}

func (c *GroupHandler) Read() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.Println("read groups")
		return ctx.SendString("group of contacts")
	}
}

func (c *GroupHandler) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		return ctx.SendString("create ugroup of contactsser")
	}
}

func (c *GroupHandler) Update() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		return ctx.SendString("update group of contacts")
	}
}

func (c *GroupHandler) Delete() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		return ctx.SendString("delete group of contacts")
	}
}
