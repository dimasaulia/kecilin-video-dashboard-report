package service

import "recserver/app/api/video_record/payload"

type VideoRecordService interface {
	Report(req payload.ReportRequest) (payload.ReportResponse, error)
}
