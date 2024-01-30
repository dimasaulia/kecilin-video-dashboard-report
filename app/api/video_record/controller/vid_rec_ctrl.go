package controller

import "github.com/gofiber/fiber/v2"

type VideoRecController interface {
	Report(ctx *fiber.Ctx) error
}
