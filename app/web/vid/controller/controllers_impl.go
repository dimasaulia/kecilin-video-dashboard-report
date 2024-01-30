package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type WebWebMapsViewImpl struct {
}

func NewWebFeatureView() WebMapsView {
	return &WebWebMapsViewImpl{}
}

func (web *WebWebMapsViewImpl) VideoReport(ctx *fiber.Ctx) error {
	return ctx.Render("video", fiber.Map{
		"title":   "Video",
		"styles":  []string{"/public/style/kecilin.0.0.1.css"},
		"scripts": []string{"/public/js/util/flowbite.js", "/public/js/util/datepicker.js", "/public/js/util/apexcharts.js", "https://code.jquery.com/jquery-latest.min.js", "https://d3js.org/d3.v4.min.js", "/public/js/util/d3-timelines.js", "/public/js/util/http_fetcher.js", "/public/js/util/data_loader.js", "/public/js/video_report.js"},
	})
}
