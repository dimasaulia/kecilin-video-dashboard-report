package router

import (
	"recserver/app/api/video_record/controller"

	"github.com/gofiber/fiber/v2"
)

type PublicRoutersImpl struct {
	App *fiber.App
}

func SetupVideoRecordRoutes(app *fiber.App, ctr controller.VideoRecController) {
	api := app.Group("/api")
	v1 := api.Group("/v1/video/")

	v1.Get("/report", ctr.Report)
}
