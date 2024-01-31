package router

import (
	rec_ctr "recserver/app/api/video_record/controller"
	rec_router "recserver/app/api/video_record/router"
	rec_svc "recserver/app/api/video_record/service"

	rec_view_ctr "recserver/app/web/vid/controller"
	rec_view_router "recserver/app/web/vid/router"

	"github.com/gofiber/fiber/v2"
)

type PublicRoutersImpl struct {
	App *fiber.App
}

func NewRouters(app *fiber.App) PublicRouters {
	return &PublicRoutersImpl{
		App: app,
	}
}

func (r PublicRoutersImpl) Setup() {
	r.App.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/video/report", fiber.StatusFound)
	})

	recSvc := rec_svc.NewVideoRecordService()
	recCtr := rec_ctr.NewVideoRecController(recSvc)
	rec_router.SetupVideoRecordRoutes(r.App, recCtr)

	recView := rec_view_ctr.NewWebFeatureView()
	rec_view_router.SetupVideoRouters(r.App, recView)

}
