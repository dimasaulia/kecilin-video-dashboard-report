package payload

import "time"

type TimeError struct {
	EndTime   time.Time `json:"end_time"`
	StartTime time.Time `json:"start_time"`
}

type ErrorRecording struct {
	Duration  float64 `json:"duration"`
	FileName  string  `json:"filename"`
	TimeError TimeError
}

type RecordingFile struct {
	Duration float64 `json:"duration"`
	FileName string  `json:"filename"`
}

type Result struct {
	ErrorPercentage  float64          `json:"error_percentage"`
	ErrorTime        float64          `json:"error_time"`
	RecordPercentage float64          `json:"record_percentage"`
	RecordTime       float64          `json:"record_time"`
	TotalError       int              `json:"total_error"`
	TotalRecording   int              `json:"total_recording"`
	TotalTime        float64          `json:"total_time"`
	Error            []ErrorRecording `json:"error"`
	RecordingFile    []RecordingFile  `json:"recording_file"`
}

type ReportResponse struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Result    Result    `json:"result"`
}
