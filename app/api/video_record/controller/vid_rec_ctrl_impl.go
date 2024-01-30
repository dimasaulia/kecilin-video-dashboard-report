package controller

import (
	"recserver/app/api/video_record/payload"
	"recserver/app/api/video_record/service"

	"github.com/gofiber/fiber/v2"
)

type VideoRecControllerImpl struct {
	Svc service.VideoRecordService
}

func NewVideoRecController(svc service.VideoRecordService) VideoRecController {
	return &VideoRecControllerImpl{
		Svc: svc,
	}
}

func (ci VideoRecControllerImpl) Report(ctx *fiber.Ctx) error {
	req := new(payload.ReportRequest)
	query := ctx.Queries()
	req.StartDate = query["start_date"]
	req.EndDate = query["end_date"]

	resp, err := ci.Svc.Report(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.Status(fiber.StatusCreated).JSON(resp)
}
