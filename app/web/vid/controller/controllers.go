package controllers

import "github.com/gofiber/fiber/v2"

type WebMapsView interface {
	VideoReport(ctx *fiber.Ctx) error
}
