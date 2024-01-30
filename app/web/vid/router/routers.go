package routers

import (
	controllers "recserver/app/web/vid/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupVideoRouters(app *fiber.App, cont controllers.WebMapsView) {
	videoGroup := app.Group("/video")
	videoGroup.Get("/report", cont.VideoReport)
}
