package service

import (
	"fmt"
	"os"
	"recserver/app/api/video_record/payload"
	"recserver/util/video"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type RecordInfo struct {
	FileName     string
	FileSize     int64
	FileDuration int64
	FileDate     time.Time
}

type VideoRecordServiceImpl struct{}

func NewVideoRecordService() VideoRecordService {
	return &VideoRecordServiceImpl{}
}

const timeLayout = "2006-01-02T15-04-05"

func (si *VideoRecordServiceImpl) Report(req payload.ReportRequest) (payload.ReportResponse, error) {
	resp := payload.ReportResponse{}
	records := []RecordInfo{}
	recordError := []payload.ErrorRecording{}
	recordSuccess := []payload.RecordingFile{}
	startTime, err := time.Parse(timeLayout, req.StartDate)
	if err != nil {
		log.Error(err)
		return resp, err
	}

	endTime, err := time.Parse(timeLayout, req.EndDate)
	if err != nil {
		log.Error(err)
		return resp, err
	}

	dir, err := os.Open("./public/video")
	if err != nil {
		log.Error(err)
		return resp, err
	}
	defer dir.Close()

	recordFiles, err := dir.Readdir(-1)
	if err != nil {
		log.Error(err)
		return resp, err
	}

	// Scan All File First
	for _, record := range recordFiles {
		if record.Mode().IsRegular() {

			fileName := record.Name()
			fileSize := record.Size()
			// File Name memiliki format 2024-01-09T00-03-00.mp4, untuk mendapatkan tanggal maka mp4 harus dihapus terlebih dahulu
			clearFileName := strings.Replace(fileName, ".mp4", "", 1)
			// Convert string to time
			recordDate, err := time.Parse(timeLayout, clearFileName)
			if err != nil {
				log.Error(err)
				return resp, err
			}

			// Filter File Berdasarkan Waktu
			if recordDate.After(startTime) && recordDate.Before(endTime) {
				videoDuration, err := video.GetVideoDuration(fmt.Sprintf("./public/video/%v", record.Name()))
				if err != nil {
					fmt.Println(err)
				}

				records = append(records, RecordInfo{
					FileName:     fileName,
					FileSize:     fileSize,
					FileDate:     recordDate,
					FileDuration: int64(videoDuration.Seconds()),
				})
			}
		}
	}

	// Algoritma untuk memastikan durasi antar file tidak lebih dr 5 menit dan ukuran file di atas threshold
	for i := 0; i < len(records)-1; i++ {
		recordDuration := records[i+1].FileDate.Sub(records[i].FileDate)
		if recordDuration.Minutes() > 5 || records[i].FileDuration < 300 {
			fileBeforBroken := records[i]
			fileAfterBroken := records[i+1]
			duration := recordDuration.Seconds()
			if duration < 300 {
				duration = float64(records[i].FileDuration)
			}

			recordError = append(recordError, payload.ErrorRecording{
				FileName: fileBeforBroken.FileName,
				Duration: float64(duration),
				TimeError: payload.TimeError{
					EndTime:   fileAfterBroken.FileDate,
					StartTime: fileBeforBroken.FileDate,
				},
			})
		} else {

			recordSuccess = append(recordSuccess, payload.RecordingFile{
				FileName: records[i].FileName,
				Duration: float64(records[i].FileDuration),
			})
		}

		// Untuk record terakhir karena tidak memiliki pembanding untuk pengecekan interval maka cek durasinya saja
		if i+1 == len(records)-1 {
			if records[i+1].FileDuration == 300 {
				recordSuccess = append(recordSuccess, payload.RecordingFile{
					FileName: records[i].FileName,
					Duration: float64(records[i].FileDuration),
				})
			} else {
				fileBeforBroken := records[i]
				fileAfterBroken := records[i+1]
				duration := records[i].FileDuration

				recordError = append(recordError, payload.ErrorRecording{
					FileName: fileBeforBroken.FileName,
					Duration: float64(duration),
					TimeError: payload.TimeError{
						StartTime: fileAfterBroken.FileDate,
						EndTime:   fileBeforBroken.FileDate,
					},
				})
			}
		}
	}

	var errDuration float64
	var successDuration float64
	for _, v := range recordError {
		errDuration += v.Duration
	}

	for _, v := range recordSuccess {
		successDuration += v.Duration
	}

	// Resp
	resp.StartTime = startTime
	resp.EndTime = endTime
	resp.Result.ErrorTime = errDuration
	resp.Result.RecordTime = successDuration
	resp.Result.TotalTime = successDuration + errDuration
	resp.Result.ErrorPercentage = (errDuration / (errDuration + successDuration)) * 100
	resp.Result.RecordPercentage = (successDuration / (errDuration + successDuration)) * 100
	resp.Result.TotalError = len(recordError)
	resp.Result.TotalRecording = len(recordSuccess)
	resp.Result.Error = recordError
	resp.Result.RecordingFile = recordSuccess
	return resp, nil
}
